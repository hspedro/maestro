package workers

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/topfreegames/maestro/internal/config"
	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/services/operation_manager"
	"go.uber.org/zap"
)

// configurations paths for the worker
const (
	// Sync period: waiting time window respected by workers in
	// order to control executions
	SyncPeriodPath = "operation.worker.sync.period"
)

type OperationWorker struct {
	run              bool
	syncPeriod       int
	scheduler        *entities.Scheduler
	operationManager operation_manager.OperationManager
}

func NewOperationWorker(
	scheduler *entities.Scheduler,
	configs config.Config,
	operationManager operation_manager.OperationManager,
) *OperationWorker {
	return &OperationWorker{
		scheduler:        scheduler,
		operationManager: operationManager,
		syncPeriod:       configs.GetInt(SyncPeriodPath),
	}
}

func (w *OperationWorker) Start(ctx context.Context) error {

	w.run = true
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(w.syncPeriod) * time.Second)
	defer ticker.Stop()

	for w.run == true {
		select {
		case <-ticker.C:
			zap.L().Info("Running operation worker", zap.String("scheduler_name", w.scheduler.Name))

		case sig := <-sigchan:
			zap.L().Warn("caught signal: terminating\n", zap.String("signal", sig.String()))
			w.run = false
		}
	}

	return nil
}

func (w *OperationWorker) Stop(ctx context.Context) error {

	w.run = false

	return nil
}
