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

package healthcontroller

import (
	"context"
	"fmt"

	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/entities/game_room"
	"github.com/topfreegames/maestro/internal/core/logs"
	"github.com/topfreegames/maestro/internal/core/operations/add_rooms"
	"github.com/topfreegames/maestro/internal/core/operations/remove_rooms"

	"github.com/topfreegames/maestro/internal/core/ports"

	"github.com/topfreegames/maestro/internal/core/entities/operation"
	"github.com/topfreegames/maestro/internal/core/operations"
	"go.uber.org/zap"
)

type SchedulerHealthControllerExecutor struct {
	roomStorage      ports.RoomStorage
	instanceStorage  ports.GameRoomInstanceStorage
	schedulerStorage ports.SchedulerStorage
	operationManager ports.OperationManager
}

var _ operations.Executor = (*SchedulerHealthControllerExecutor)(nil)

func NewExecutor(roomStorage ports.RoomStorage, instanceStorage ports.GameRoomInstanceStorage, schedulerStorage ports.SchedulerStorage, operationManager ports.OperationManager) *SchedulerHealthControllerExecutor {
	return &SchedulerHealthControllerExecutor{
		roomStorage:      roomStorage,
		instanceStorage:  instanceStorage,
		schedulerStorage: schedulerStorage,
		operationManager: operationManager,
	}
}

func (ex *SchedulerHealthControllerExecutor) Execute(ctx context.Context, op *operation.Operation, definition operations.Definition) operations.ExecutionError {
	logger := zap.L().With(
		zap.String(logs.LogFieldSchedulerName, op.SchedulerName),
		zap.String(logs.LogFieldOperationDefinition, op.DefinitionName),
		zap.String("operation_phase", "Execute"),
		zap.String(logs.LogFieldOperationID, op.ID),
	)

	gameRoomIDs, instances, scheduler, err := ex.loadActualState(ctx, op, logger)
	if err != nil {
		return operations.NewErrUnexpected(err)
	}

	nonexistentGameRoomIDs := ex.checkNonexistentGameRoomsIDs(gameRoomIDs, instances)
	if len(nonexistentGameRoomIDs) > 0 {
		logger.Error("found registered rooms that no longer exists")
		ex.tryEnsureCorrectRoomsOnStorage(ctx, op, logger, nonexistentGameRoomIDs)

	}

	if len(instances) != scheduler.RoomsReplicas {
		err = ex.ensureDesiredAmountOfInstances(ctx, op, logger, len(instances), scheduler.RoomsReplicas)
		if err != nil {
			logger.Error("cannot ensure desired amount of instances", zap.Error(err))
			return operations.NewErrUnexpected(err)
		}
	}

	return nil
}

func (ex *SchedulerHealthControllerExecutor) Rollback(ctx context.Context, op *operation.Operation, definition operations.Definition, executeErr operations.ExecutionError) error {
	return nil
}

func (ex *SchedulerHealthControllerExecutor) Name() string {
	return OperationName
}

func (ex *SchedulerHealthControllerExecutor) loadActualState(ctx context.Context, op *operation.Operation, logger *zap.Logger) (gameRoomIDs []string, instances []*game_room.Instance, scheduler *entities.Scheduler, err error) {
	gameRoomIDs, err = ex.roomStorage.GetAllRoomIDs(ctx, op.SchedulerName)
	if err != nil {
		logger.Error("error fetching game rooms")
		return gameRoomIDs, instances, scheduler, err
	}
	instances, err = ex.instanceStorage.GetAllInstances(ctx, op.SchedulerName)
	if err != nil {
		logger.Error("error fetching instances")
		return gameRoomIDs, instances, scheduler, err
	}

	scheduler, err = ex.schedulerStorage.GetScheduler(ctx, op.SchedulerName)
	if err != nil {
		return gameRoomIDs, instances, scheduler, err
	}
	return gameRoomIDs, instances, scheduler, err
}

func (ex *SchedulerHealthControllerExecutor) checkNonexistentGameRoomsIDs(gameRoomIDs []string, gameRoomInstances []*game_room.Instance) []string {
	var nonexistentGameRoomsIDs []string
	for _, gameRoomID := range gameRoomIDs {
		found := false
		for _, instance := range gameRoomInstances {
			if instance.ID == gameRoomID {
				found = true
				break
			}
		}
		if !found {
			nonexistentGameRoomsIDs = append(nonexistentGameRoomsIDs, gameRoomID)
		}
	}
	return nonexistentGameRoomsIDs
}

func (ex *SchedulerHealthControllerExecutor) tryEnsureCorrectRoomsOnStorage(ctx context.Context, op *operation.Operation, logger *zap.Logger, nonexistentGameRoomIDs []string) {
	for _, gameRoomID := range nonexistentGameRoomIDs {
		err := ex.roomStorage.DeleteRoom(ctx, op.SchedulerName, gameRoomID)
		if err != nil {
			msg := fmt.Sprintf("could not delete nonexistent room %s from storage", gameRoomID)
			logger.Warn(msg, zap.Error(err))
			continue
		}
		logger.Sugar().Infof("remove nonexistent room on storage: %s", gameRoomID)
	}
}

func (ex *SchedulerHealthControllerExecutor) ensureDesiredAmountOfInstances(ctx context.Context, op *operation.Operation, logger *zap.Logger, actualAmount, desiredAmount int) error {
	var msgToAppend string

	if actualAmount > desiredAmount {
		removeAmount := actualAmount - desiredAmount
		removeOperation, err := ex.operationManager.CreatePriorityOperation(ctx, op.SchedulerName, &remove_rooms.RemoveRoomsDefinition{
			Amount: removeAmount,
		})
		if err != nil {
			return err
		}
		msgToAppend = fmt.Sprintf("created operation (id: %s) to remove %v rooms.", removeOperation.ID, removeAmount)
	} else {
		addAmount := desiredAmount - actualAmount
		addOperation, err := ex.operationManager.CreatePriorityOperation(ctx, op.SchedulerName, &add_rooms.AddRoomsDefinition{
			Amount: int32(addAmount),
		})
		if err != nil {
			return err
		}
		msgToAppend = fmt.Sprintf("created operation (id: %s) to add %v rooms.", addOperation.ID, addAmount)
	}

	logger.Info(msgToAppend)
	ex.operationManager.AppendOperationEventToExecutionHistory(ctx, op, msgToAppend)
	return nil
}
