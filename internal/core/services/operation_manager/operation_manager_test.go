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

//+build unit

package operation_manager

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	opflow "github.com/topfreegames/maestro/internal/adapters/operation_flow/mock"
	opstorage "github.com/topfreegames/maestro/internal/adapters/operation_storage/mock"
	"github.com/topfreegames/maestro/internal/core/entities/operation"
	"github.com/topfreegames/maestro/internal/core/operations"
	"github.com/topfreegames/maestro/internal/core/ports"
	porterrors "github.com/topfreegames/maestro/internal/core/ports/errors"
)

type testOperationDefinition struct {
	marshalResult   []byte
	unmarshalResult error
}

func (d *testOperationDefinition) Marshal() []byte            { return d.marshalResult }
func (d *testOperationDefinition) Unmarshal(raw []byte) error { return d.unmarshalResult }
func (d *testOperationDefinition) Name() string               { return "testOperationDefinition" }
func (d *testOperationDefinition) ShouldExecute(_ context.Context, _ []*operation.Operation) bool {
	return false
}

type opMatcher struct {
	status operation.Status
	def    operations.Definition
}

func (m *opMatcher) Matches(x interface{}) bool {
	op, _ := x.(*operation.Operation)
	_, err := uuid.Parse(op.ID)
	return err == nil && op.Status == m.status && m.def.Name() == op.DefinitionName
}

func (m *opMatcher) String() string {
	return fmt.Sprintf("a operation with definition \"%s\"", m.def.Name())
}

func TestCreateOperation(t *testing.T) {
	cases := map[string]struct {
		definition operations.Definition
		storageErr error
		flowErr    error
	}{
		"create without errors": {
			definition: &testOperationDefinition{marshalResult: []byte("test")},
		},
		"create with storage errors": {
			definition: &testOperationDefinition{},
			storageErr: porterrors.ErrUnexpected,
		},
		"create with flow errors": {
			definition: &testOperationDefinition{},
			flowErr:    porterrors.ErrUnexpected,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			schedulerName := "scheduler_name"
			operationFlow := opflow.NewMockOperationFlow(mockCtrl)
			operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
			definitionConstructors := operations.NewDefinitionConstructors()
			opManager := New(operationFlow, operationStorage, definitionConstructors)

			ctx := context.Background()
			testDefinition, _ := test.definition.(*testOperationDefinition)
			operationStorage.EXPECT().CreateOperation(ctx, &opMatcher{operation.StatusPending, test.definition}, testDefinition.marshalResult).Return(test.storageErr)

			if test.storageErr == nil {
				operationFlow.EXPECT().InsertOperationID(ctx, schedulerName, gomock.Any()).Return(test.flowErr)
			}

			op, err := opManager.CreateOperation(ctx, schedulerName, test.definition)

			if test.storageErr != nil {
				require.ErrorIs(t, err, test.storageErr)
				require.Nil(t, op)
				return
			}

			if test.flowErr != nil {
				require.ErrorIs(t, err, test.flowErr)
				require.Nil(t, op)
				return
			}

			require.NotNil(t, op)
			require.Equal(t, operation.StatusPending, op.Status)
		})
	}
}

func TestGetOperation(t *testing.T) {
	t.Run("find operation", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operations.Definition { return &testOperationDefinition{} }

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		definitionConstructors[defFunc().Name()] = defFunc
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			nil,
		)

		op, definition, err := opManager.GetOperation(ctx, schedulerName, operationID)
		require.NoError(t, err)
		require.NotNil(t, op)
		require.Equal(t, operationID, op.ID)
		require.Equal(t, schedulerName, op.SchedulerName)
		require.IsType(t, defFunc(), definition)
	})

	t.Run("defition not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operations.Definition { return &testOperationDefinition{} }

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			nil,
		)

		_, _, err := opManager.GetOperation(ctx, schedulerName, operationID)
		require.Error(t, err)
	})

	t.Run("operation not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operations.Definition { return &testOperationDefinition{} }

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			porterrors.ErrNotFound,
		)

		_, _, err := opManager.GetOperation(ctx, schedulerName, operationID)
		require.Error(t, err)
	})

	t.Run("unmarshal error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operations.Definition { return &testOperationDefinition{unmarshalResult: errors.New("invalid")} }

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			porterrors.ErrNotFound,
		)

		_, _, err := opManager.GetOperation(ctx, schedulerName, operationID)
		require.Error(t, err)
	})
}

func TestNextSchedulerOperation(t *testing.T) {
	t.Run("fetch operation", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operations.Definition { return &testOperationDefinition{} }
		definitionConstructors := operations.NewDefinitionConstructors()
		definitionConstructors[defFunc().Name()] = defFunc

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"

		operationFlow.EXPECT().NextOperationID(ctx, schedulerName).Return(operationID, nil)
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			nil,
		)

		op, definition, err := opManager.NextSchedulerOperation(ctx, schedulerName)
		require.NoError(t, err)
		require.NotNil(t, op)
		require.Equal(t, operationID, op.ID)
		require.Equal(t, schedulerName, op.SchedulerName)
		require.IsType(t, defFunc(), definition)
	})

	t.Run("no next operation", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operations.Definition { return &testOperationDefinition{} }
		definitionConstructors := operations.NewDefinitionConstructors()
		definitionConstructors[defFunc().Name()] = defFunc

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationFlow.EXPECT().NextOperationID(ctx, schedulerName).Return("", porterrors.ErrUnexpected)

		_, _, err := opManager.NextSchedulerOperation(ctx, schedulerName)
		require.Error(t, err)
	})

	t.Run("operation not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operations.Definition { return &testOperationDefinition{} }
		definitionConstructors := operations.NewDefinitionConstructors()
		definitionConstructors[defFunc().Name()] = defFunc

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"

		operationFlow.EXPECT().NextOperationID(ctx, schedulerName).Return(operationID, nil)
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			nil,
			[]byte{},
			porterrors.ErrNotFound,
		)

		_, _, err := opManager.NextSchedulerOperation(ctx, schedulerName)
		require.Error(t, err)
	})
}

func TestStartOperation(t *testing.T) {
	t.Run("starts operation with success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		op := &operation.Operation{ID: uuid.NewString(), DefinitionName: (&testOperationDefinition{}).Name()}

		operationStorage.EXPECT().UpdateOperationStatus(ctx, op.SchedulerName, op.ID, operation.StatusInProgress).Return(nil)
		err := opManager.StartOperation(ctx, op, func() {})
		require.NoError(t, err)
	})
}

func TestFinishOperation(t *testing.T) {
	t.Run("finishes operation with success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		op := &operation.Operation{
			SchedulerName:  uuid.NewString(),
			ID:             uuid.NewString(),
			DefinitionName: (&testOperationDefinition{}).Name(),
		}

		operationStorage.EXPECT().UpdateOperationStatus(ctx, op.SchedulerName, op.ID, operation.StatusInProgress).Return(nil)
		err := opManager.StartOperation(ctx, op, func() {})
		require.NoError(t, err)

		expectedStatus := operation.StatusError
		op.Status = expectedStatus
		operationStorage.EXPECT().UpdateOperationStatus(ctx, op.SchedulerName, op.ID, expectedStatus).Return(nil)
		err = opManager.FinishOperation(ctx, op)
		require.NoError(t, err)
	})
}

func TestListSchedulerActiveOperations(t *testing.T) {
	t.Run("it returns an operation list with pending status", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		operationsResult := []*operation.Operation{
			{ID: uuid.NewString()},
			{ID: uuid.NewString()},
			{ID: uuid.NewString()},
		}

		schedulerName := "test-scheduler"
		operationStorage.EXPECT().ListSchedulerActiveOperations(ctx, schedulerName).Return(operationsResult, nil)
		operations, err := opManager.ListSchedulerActiveOperations(ctx, schedulerName)
		require.NoError(t, err)
		require.ElementsMatch(t, operationsResult, operations)
	})
}

func TestListSchedulerFinishedOperations(t *testing.T) {
	t.Run("it returns an operation list with finished status", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		operationsResult := []*operation.Operation{
			{ID: uuid.NewString()},
			{ID: uuid.NewString()},
			{ID: uuid.NewString()},
		}

		schedulerName := "test-scheduler"
		operationStorage.EXPECT().ListSchedulerActiveOperations(ctx, schedulerName).Return(operationsResult, nil)
		operations, err := opManager.ListSchedulerActiveOperations(ctx, schedulerName)
		require.NoError(t, err)
		require.ElementsMatch(t, operationsResult, operations)
	})
}

func TestListSchedulerPendingOperations(t *testing.T) {
	t.Run("it returns an operation list with pending status", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		definitionConstructors := operations.NewDefinitionConstructors()
		opManager := New(operationFlow, operationStorage, definitionConstructors)

		ctx := context.Background()
		operationsResult := []*operation.Operation{
			{ID: uuid.NewString()},
			{ID: uuid.NewString()},
			{ID: uuid.NewString()},
		}

		schedulerName := "test-scheduler"
		operationFlow.EXPECT().ListSchedulerPendingOperationIDs(ctx, schedulerName).Return([]string{"1", "2", "3"}, nil)
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, "1").Return(operationsResult[0], []byte{}, nil)
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, "2").Return(operationsResult[1], []byte{}, nil)
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, "3").Return(operationsResult[2], []byte{}, nil)

		operations, err := opManager.ListSchedulerPendingOperations(ctx, schedulerName)
		require.NoError(t, err)
		require.ElementsMatch(t, operationsResult, operations)
	})
}

func TestWatchOperationCancellationRequests(t *testing.T) {
	schedulerName := uuid.New().String()
	operationID := uuid.New().String()

	t.Run("cancels a operation successfully", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		opManager := New(operationFlow, operationStorage, nil)

		cancelableContext, cancelFunction := context.WithCancel(context.Background())
		opManager.operationCancelFunctions.putFunction(schedulerName, operationID, cancelFunction)

		requestChannel := make(chan ports.OperationCancellationRequest, 1000)
		operationFlow.EXPECT().WatchOperationCancellationRequests(gomock.Any()).Return(requestChannel)

		ctx, ctxCancelFunction := context.WithCancel(context.Background())
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(&operation.Operation{
			SchedulerName: schedulerName,
			ID:            operationID,
			Status:        operation.StatusInProgress,
		}, nil, nil)

		go func() {
			err := opManager.WatchOperationCancellationRequests(ctx)
			require.NoError(t, err)
		}()

		requestChannel <- ports.OperationCancellationRequest{
			SchedulerName: schedulerName,
			OperationID:   operationID,
		}

		require.Eventually(t, func() bool {
			if cancelableContext.Err() != nil {
				require.ErrorIs(t, cancelableContext.Err(), context.Canceled)
				return true
			}

			return false
		}, time.Second, 100*time.Millisecond)

		ctxCancelFunction()
	})
}