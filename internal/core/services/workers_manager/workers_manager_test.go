package workers_manager

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	opflow "github.com/topfreegames/maestro/internal/adapters/operation_flow/mock"
	opstorage "github.com/topfreegames/maestro/internal/adapters/operation_storage/mock"
	schedulerStorageMock "github.com/topfreegames/maestro/internal/adapters/scheduler_storage/mock"
	configMock "github.com/topfreegames/maestro/internal/config/mock"
	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/ports/errors"
	"github.com/topfreegames/maestro/internal/core/services/operation_manager"
	"github.com/topfreegames/maestro/internal/core/workers"
)

func TestStart(t *testing.T) {

	t.Run("with success", func(t *testing.T) {

		core, recorded := observer.New(zap.InfoLevel)
		zl := zap.New(core)
		zap.ReplaceGlobals(zl)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		configs := configMock.NewMockConfig(mockCtrl)
		schedulerStorage := schedulerStorageMock.NewMockSchedulerStorage(mockCtrl)
		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		operationManager := operation_manager.New(operationFlow, operationStorage)

		configs.EXPECT().GetInt(syncOperationWorkersIntervalPath).Return(10)
		configs.EXPECT().GetInt(workers.OperationWorkerIntervalPath).Return(10)
		schedulerStorage.EXPECT().GetAllSchedulers(context.Background()).Return([]*entities.Scheduler{
			{
				Name:            "zooba-us",
				Game:            "zooba",
				State:           entities.StateCreating,
				RollbackVersion: "1.0.0",
				PortRange: &entities.PortRange{
					Start: 1,
					End:   10000,
				},
			},
		}, nil)

		workersManager := NewWorkersManager(configs, schedulerStorage, *operationManager)

		err := workersManager.Start(context.Background())
		require.NoError(t, err)

		require.Contains(t, workersManager.CurrentWorkers, "zooba-us")

		assertLogMessages(t, recorded, map[zapcore.Level][]string{
			zap.InfoLevel: {"starting to sync operation workers",
				"schedulers found, syncing operation workers",
				"new operation worker running",
				"all operation workers in sync"},
		})
	})

	t.Run("with soft error when schedulerStorage fails to list all schedulers", func(t *testing.T) {

		core, recorded := observer.New(zap.InfoLevel)
		zl := zap.New(core)
		zap.ReplaceGlobals(zl)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		configs := configMock.NewMockConfig(mockCtrl)
		schedulerStorage := schedulerStorageMock.NewMockSchedulerStorage(mockCtrl)
		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		operationManager := operation_manager.New(operationFlow, operationStorage)

		configs.EXPECT().GetInt(syncOperationWorkersIntervalPath).Return(10)
		schedulerStorage.EXPECT().GetAllSchedulers(context.Background()).Return(nil, errors.ErrUnexpected)

		workersManager := NewWorkersManager(configs, schedulerStorage, *operationManager)

		err := workersManager.Start(context.Background())
		require.NoError(t, err)

		require.Empty(t, workersManager.CurrentWorkers)

		assertLogMessages(t, recorded, map[zapcore.Level][]string{
			zap.ErrorLevel: {"initial sync operation workers failed"},
		})
	})

	t.Run("with success when scheduler added after bootstrap", func(t *testing.T) {

		core, recorded := observer.New(zap.InfoLevel)
		zl := zap.New(core)
		zap.ReplaceGlobals(zl)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		configs := configMock.NewMockConfig(mockCtrl)
		schedulerStorage := schedulerStorageMock.NewMockSchedulerStorage(mockCtrl)
		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		operationManager := operation_manager.New(operationFlow, operationStorage)

		configs.EXPECT().GetInt(syncOperationWorkersIntervalPath).Return(10)
		configs.EXPECT().GetInt(workers.OperationWorkerIntervalPath).Return(10)
		schedulerStorage.EXPECT().GetAllSchedulers(context.Background()).Return([]*entities.Scheduler{}, nil)

		workersManager := NewWorkersManager(configs, schedulerStorage, *operationManager)
		workersManager.Start(context.Background())
		require.Empty(t, workersManager.CurrentWorkers)

		schedulerStorage.EXPECT().GetAllSchedulers(context.Background()).Return([]*entities.Scheduler{
			{
				Name:            "zooba-us",
				Game:            "zooba",
				State:           entities.StateCreating,
				RollbackVersion: "1.0.0",
				PortRange: &entities.PortRange{
					Start: 1,
					End:   10000,
				},
			},
		}, nil)

		workersManager.SyncOperationWorkers(context.Background())

		require.Contains(t, workersManager.CurrentWorkers, "zooba-us")

		assertLogMessages(t, recorded, map[zapcore.Level][]string{
			zap.InfoLevel: {"starting to sync operation workers",
				"schedulers found, syncing operation workers",
				"new operation worker running",
				"all operation workers in sync"},
		})
	})

	t.Run("with success when scheduler removed after bootstrap", func(t *testing.T) {

		core, recorded := observer.New(zap.InfoLevel)
		zl := zap.New(core)
		zap.ReplaceGlobals(zl)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		configs := configMock.NewMockConfig(mockCtrl)
		schedulerStorage := schedulerStorageMock.NewMockSchedulerStorage(mockCtrl)
		operationFlow := opflow.NewMockOperationFlow(mockCtrl)
		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		operationManager := operation_manager.New(operationFlow, operationStorage)

		configs.EXPECT().GetInt(syncOperationWorkersIntervalPath).AnyTimes().Return(1)
		configs.EXPECT().GetInt(workers.OperationWorkerIntervalPath).AnyTimes().Return(1)

		schedulerStorage.EXPECT().GetAllSchedulers(context.Background()).Times(3).Return([]*entities.Scheduler{
			{
				Name:            "zooba-us",
				Game:            "zooba",
				State:           entities.StateCreating,
				RollbackVersion: "1.0.0",
				PortRange: &entities.PortRange{
					Start: 1,
					End:   10000,
				},
			},
		}, nil)

		workersManager := NewWorkersManager(configs, schedulerStorage, *operationManager)
		workersManager.Start(context.Background())

		// Has to sleep 1 second in order to start the goroutine
		time.Sleep(2 * time.Second)

		require.Contains(t, workersManager.CurrentWorkers, "zooba-us")
		operationWorker := workersManager.CurrentWorkers["zooba-us"]
		require.Equal(t, true, operationWorker.IsRunning(context.Background()))
		require.Greater(t, operationWorker.CountRuns(context.Background()), 0)

		assertLogMessages(t, recorded, map[zapcore.Level][]string{
			zap.InfoLevel: {"starting to sync operation workers",
				"schedulers found, syncing operation workers",
				"new operation worker running",
				"all operation workers in sync"},
		})

		core, recorded = observer.New(zap.InfoLevel)
		zl = zap.New(core)
		zap.ReplaceGlobals(zl)

		schedulerStorage.EXPECT().GetAllSchedulers(context.Background()).Return([]*entities.Scheduler{}, nil)

		workersManager.SyncOperationWorkers(context.Background())

		require.Empty(t, workersManager.CurrentWorkers)
		require.Equal(t, false, operationWorker.IsRunning(context.Background()))

		assertLogMessages(t, recorded, map[zapcore.Level][]string{
			zap.InfoLevel: {"starting to sync operation workers",
				"schedulers found, syncing operation workers",
				"all operation workers in sync"},
		})

	})
}

func assertLogMessages(t *testing.T, recorded *observer.ObservedLogs, messages map[zapcore.Level][]string) {
	for level, values := range messages {

		levelRecords := recorded.FilterLevelExact(level)
		for _, message := range values {
			require.NotEmpty(t, levelRecords.Filter(func(le observer.LoggedEntry) bool {
				return le.Message == message
			}))
		}
	}
}
