// MIT License
//
// Copyright (c) 2021 TFG Co
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package operation_manager

import (
	"context"
	goerrors "errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/topfreegames/maestro/internal/core/entities/operation"
	"github.com/topfreegames/maestro/internal/core/operations"
	"github.com/topfreegames/maestro/internal/core/ports"
	"github.com/topfreegames/maestro/internal/core/ports/errors"
	"go.uber.org/zap"
)

type OperationCancelFunctions struct {
	mutex     sync.RWMutex
	functions map[string]map[string]context.CancelFunc
}

func NewOperationCancelFunctions() *OperationCancelFunctions {
	return &OperationCancelFunctions{
		functions: map[string]map[string]context.CancelFunc{},
	}
}

func (of *OperationCancelFunctions) putFunction(schedulerName, operationID string, cancelFunc context.CancelFunc) {
	of.mutex.Lock()
	schedulerCancelFunctions := of.functions[schedulerName]
	if schedulerCancelFunctions == nil {
		of.functions[schedulerName] = map[string]context.CancelFunc{}
	}
	of.functions[schedulerName][operationID] = cancelFunc
	of.mutex.Unlock()
}

func (of *OperationCancelFunctions) removeFunction(schedulerName, operationID string) {
	of.mutex.Lock()

	schedulerOperationCancellationFunctions := of.functions[schedulerName]
	if schedulerOperationCancellationFunctions == nil {
		return
	}

	delete(schedulerOperationCancellationFunctions, operationID)

	of.mutex.Unlock()
}

func (of *OperationCancelFunctions) getFunction(schedulerName, operationID string) (context.CancelFunc, error) {
	of.mutex.RLock()
	schedulerOperationCancellationFunctions := of.functions[schedulerName]
	if schedulerOperationCancellationFunctions == nil {
		return nil, errors.NewErrNotFound("no cancel scheduler found for scheduler name: %s", schedulerName)
	}

	if schedulerOperationCancellationFunctions[operationID] == nil {
		return nil, errors.NewErrNotFound("no cancel function found for scheduler name: %s and operation id: %s", schedulerName, operationID)
	}

	function := schedulerOperationCancellationFunctions[operationID]
	of.mutex.RUnlock()

	return function, nil
}

type OperationManager struct {
	operationCancelFunctions        *OperationCancelFunctions
	flow                            ports.OperationFlow
	storage                         ports.OperationStorage
	operationDefinitionConstructors map[string]operations.DefinitionConstructor
}

func New(flow ports.OperationFlow, storage ports.OperationStorage, operationDefinitionConstructors map[string]operations.DefinitionConstructor) *OperationManager {
	return &OperationManager{
		flow:                            flow,
		storage:                         storage,
		operationDefinitionConstructors: operationDefinitionConstructors,
		operationCancelFunctions:        NewOperationCancelFunctions(),
	}
}

func (om *OperationManager) CreateOperation(ctx context.Context, schedulerName string, definition operations.Definition) (*operation.Operation, error) {
	op := &operation.Operation{
		ID:             uuid.NewString(),
		Status:         operation.StatusPending,
		DefinitionName: definition.Name(),
		SchedulerName:  schedulerName,
		CreatedAt:      time.Now(),
	}

	err := om.storage.CreateOperation(ctx, op, definition.Marshal())
	if err != nil {
		return nil, fmt.Errorf("failed to create operation: %w", err)
	}

	err = om.flow.InsertOperationID(ctx, op.SchedulerName, op.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert operation on flow: %w", err)
	}

	return op, nil
}

func (om *OperationManager) GetOperation(ctx context.Context, schedulerName, operationID string) (*operation.Operation, operations.Definition, error) {
	op, definitionContents, err := om.storage.GetOperation(ctx, schedulerName, operationID)
	if err != nil {
		return nil, nil, err
	}

	definitionConstructor := om.operationDefinitionConstructors[op.DefinitionName]
	if definitionConstructor == nil {
		return nil, nil, fmt.Errorf("no definition constructor implemented for %s: %s", op.DefinitionName, err)
	}

	definition := definitionConstructor()
	err = definition.Unmarshal(definitionContents)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal definition: %s", err)
	}

	return op, definition, nil
}

// NextSchedulerOperation returns the next scheduler operation to be processed.
func (om *OperationManager) NextSchedulerOperation(ctx context.Context, schedulerName string) (*operation.Operation, operations.Definition, error) {
	operationID, err := om.flow.NextOperationID(ctx, schedulerName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to retrieve the next operation: %w", err)
	}

	return om.GetOperation(ctx, schedulerName, operationID)
}

// StartOperation used when an operation will start executing.
func (om *OperationManager) StartOperation(ctx context.Context, op *operation.Operation, cancelFunction context.CancelFunc) error {
	err := om.storage.UpdateOperationStatus(ctx, op.SchedulerName, op.ID, operation.StatusInProgress)
	if err != nil {
		return fmt.Errorf("failed to start operation: %w", err)
	}

	op.Status = operation.StatusInProgress

	om.operationCancelFunctions.putFunction(op.SchedulerName, op.ID, cancelFunction)
	return nil
}

func (om *OperationManager) ListSchedulerPendingOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {

	pendingOperationIDs, err := om.flow.ListSchedulerPendingOperationIDs(ctx, schedulerName)
	if err != nil {
		return nil, fmt.Errorf("failed to list all pending operations: %w", err)
	}
	pendingOperations := make([]*operation.Operation, len(pendingOperationIDs))
	for i, operationID := range pendingOperationIDs {
		op, _, err := om.storage.GetOperation(ctx, schedulerName, operationID)
		if err != nil {
			return nil, err
		}
		pendingOperations[i] = op
	}

	return pendingOperations, nil
}

func (om *OperationManager) ListSchedulerActiveOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {

	return om.storage.ListSchedulerActiveOperations(ctx, schedulerName)
}

func (om *OperationManager) ListSchedulerFinishedOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {

	return om.storage.ListSchedulerFinishedOperations(ctx, schedulerName)
}

func (om *OperationManager) FinishOperation(ctx context.Context, op *operation.Operation) error {
	err := om.storage.UpdateOperationStatus(ctx, op.SchedulerName, op.ID, op.Status)
	if err != nil {
		return fmt.Errorf("failed to finish operation: %w", err)
	}

	om.operationCancelFunctions.removeFunction(op.SchedulerName, op.ID)

	return nil
}

func (om *OperationManager) EnqueueOperationCancellationRequest(ctx context.Context, schedulerName, operationID string) error {
	err := om.flow.EnqueueOperationCancellationRequest(ctx, ports.OperationCancellationRequest{
		SchedulerName: schedulerName,
		OperationID:   operationID,
	})
	if err != nil {
		return fmt.Errorf("failed to enqueue operation cancellation request: %w", err)
	}

	return nil
}

func (om *OperationManager) WatchOperationCancellationRequests(ctx context.Context) error {
	requestChannel := om.flow.WatchOperationCancellationRequests(ctx)

	for {
		select {
		case request, ok := <-requestChannel:
			if !ok {
				return errors.NewErrUnexpected("operation cancellation request channel closed")
			}

			err := om.cancelOperation(ctx, request.SchedulerName, request.OperationID)
			if err != nil {
				zap.L().
					With(zap.String("schedulerName", request.SchedulerName)).
					With(zap.String("operationID", request.OperationID)).
					With(zap.Error(err)).
					Error("failed to cancel operation")
			}

		case <-ctx.Done():
			if goerrors.Is(ctx.Err(), context.Canceled) {
				return nil
			}

			return fmt.Errorf("loop to consume cancel operation requests received an error context event: %w", ctx.Err())
		}
	}
}

func (om *OperationManager) cancelOperation(ctx context.Context, schedulerName, operationID string) error {

	op, _, err := om.storage.GetOperation(ctx, schedulerName, operationID)
	if err != nil {
		return fmt.Errorf("failed to fetch operation from storage: %w", err)
	}

	if op.Status == operation.StatusPending {
		err = om.storage.UpdateOperationStatus(ctx, schedulerName, operationID, operation.StatusCanceled)
		if err != nil {
			return fmt.Errorf("failed update operation as canceled: %w", err)
		}
	} else {
		cancelFn, err := om.operationCancelFunctions.getFunction(schedulerName, operationID)
		if err != nil {
			return fmt.Errorf("failed to fetch cancel function from operation: %w", err)
		}
		cancelFn()
	}

	return nil
}