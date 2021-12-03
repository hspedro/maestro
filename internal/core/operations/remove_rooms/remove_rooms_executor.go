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

package remove_rooms

import (
	"context"
	"fmt"
	"sync"

	"github.com/topfreegames/maestro/internal/core/entities/operation"
	"github.com/topfreegames/maestro/internal/core/operations"
	"github.com/topfreegames/maestro/internal/core/services/room_manager"
	"go.uber.org/zap"
)

type RemoveRoomsExecutor struct {
	roomManager *room_manager.RoomManager
}

func NewExecutor(roomManager *room_manager.RoomManager) *RemoveRoomsExecutor {
	return &RemoveRoomsExecutor{roomManager}
}

func (e *RemoveRoomsExecutor) Execute(ctx context.Context, op *operation.Operation, definition operations.Definition) error {
	removeDefinition := definition.(*RemoveRoomsDefinition)
	rooms, err := e.roomManager.ListRoomsWithDeletionPriority(ctx, op.SchedulerName, removeDefinition.Version, removeDefinition.Amount, &sync.Map{})
	if err != nil {
		return fmt.Errorf("failed to list rooms to delete: %w", err)
	}

	logger := zap.L().With(
		zap.String("scheduler_name", op.SchedulerName),
		zap.String("operation_definition", definition.Name()),
		zap.String("operation_id", op.ID),
	)

	logger.Debug("start deleting rooms", zap.Int("amount", len(rooms)))

	for _, room := range rooms {
		err = e.roomManager.DeleteRoom(ctx, room)
		if err != nil {
			reportDeletionFailedTotal(op.SchedulerName, op.ID)
			logger.Warn("failed to remove rooms", zap.Error(err))
		}
	}

	logger.Debug("finished deleting rooms")
	return nil
}

// OnError will do nothing.
func (e *RemoveRoomsExecutor) OnError(_ context.Context, _ *operation.Operation, _ operations.Definition, _ error) error {
	return nil
}

func (e *RemoveRoomsExecutor) Name() string {
	return OperationName
}