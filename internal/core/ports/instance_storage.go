package ports

import (
	"context"

	"github.com/topfreegames/maestro/internal/core/entities/game_room"
)

type GameRoomInstanceStorage interface {
	GetInstance(ctx context.Context, scheduler string, roomId string) (*game_room.Instance, error)
	UpsertInstance(ctx context.Context, instance *game_room.Instance) error
	DeleteInstance(ctx context.Context, scheduler string, roomId string) error
	GetAllInstances(ctx context.Context, scheduler string) ([]*game_room.Instance, error)
	GetInstanceCount(ctx context.Context, scheduler string) (int, error)
}