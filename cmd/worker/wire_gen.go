// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/topfreegames/maestro/internal/config"
	"github.com/topfreegames/maestro/internal/core/operations/providers"
	"github.com/topfreegames/maestro/internal/core/services/operation_manager"
	"github.com/topfreegames/maestro/internal/core/services/room_manager"
	"github.com/topfreegames/maestro/internal/core/services/scheduler_manager"
	"github.com/topfreegames/maestro/internal/core/services/workers_manager"
	"github.com/topfreegames/maestro/internal/core/workers"
	"github.com/topfreegames/maestro/internal/service"
)

// Injectors from wire.go:

func initializeWorker(c config.Config, builder workers.WorkerBuilder) (*workers_manager.WorkersManager, error) {
	schedulerStorage, err := service.NewSchedulerStoragePg(c)
	if err != nil {
		return nil, err
	}
	operationFlow, err := service.NewOperationFlowRedis(c)
	if err != nil {
		return nil, err
	}
	clock := service.NewClockTime()
	operationStorage, err := service.NewOperationStorageRedis(clock, c)
	if err != nil {
		return nil, err
	}
	v := providers.ProvideDefinitionConstructors()
	operationManager := operation_manager.New(operationFlow, operationStorage, v)
	runtime, err := service.NewRuntimeKubernetes(c)
	if err != nil {
		return nil, err
	}
	portAllocator, err := service.NewPortAllocatorRandom(c)
	if err != nil {
		return nil, err
	}
	roomStorage, err := service.NewRoomStorageRedis(c)
	if err != nil {
		return nil, err
	}
	gameRoomInstanceStorage, err := service.NewGameRoomInstanceStorageRedis(c)
	if err != nil {
		return nil, err
	}
	eventsForwarder, err := service.NewEventsForwarder(c)
	if err != nil {
		return nil, err
	}
	roomManagerConfig, err := service.NewRoomManagerConfig(c)
	if err != nil {
		return nil, err
	}
	roomManager := room_manager.NewRoomManager(clock, portAllocator, roomStorage, gameRoomInstanceStorage, runtime, eventsForwarder, roomManagerConfig)
	schedulerManager := scheduler_manager.NewSchedulerManager(schedulerStorage, operationManager)
	v2 := providers.ProvideExecutors(runtime, schedulerStorage, roomManager, schedulerManager)
	workerOptions := workers.ProvideWorkerOptions(operationManager, v2, roomManager, runtime)
	workersManager := workers_manager.NewWorkersManager(builder, c, schedulerStorage, workerOptions)
	return workersManager, nil
}