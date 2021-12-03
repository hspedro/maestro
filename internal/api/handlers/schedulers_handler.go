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

package handlers

import (
	"context"
	"errors"
	"strings"
	"time"

	validator "gopkg.in/validator.v2"

	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/entities/game_room"
	portsErrors "github.com/topfreegames/maestro/internal/core/ports/errors"
	"github.com/topfreegames/maestro/internal/core/services/scheduler_manager"
	api "github.com/topfreegames/maestro/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SchedulersHandler struct {
	schedulerManager *scheduler_manager.SchedulerManager
	api.UnimplementedSchedulersServiceServer
}

func ProvideSchedulersHandler(schedulerManager *scheduler_manager.SchedulerManager) *SchedulersHandler {
	return &SchedulersHandler{
		schedulerManager: schedulerManager,
	}
}

func (h *SchedulersHandler) ListSchedulers(ctx context.Context, message *api.ListSchedulersRequest) (*api.ListSchedulersResponse, error) {
	schedulerEntities, err := h.schedulerManager.GetAllSchedulers(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	schedulers := make([]*api.SchedulerWithoutSpec, len(schedulerEntities))
	for i, entity := range schedulerEntities {
		schedulers[i] = h.fromEntitySchedulerToListResponse(entity)
	}

	return &api.ListSchedulersResponse{
		Schedulers: schedulers,
	}, nil
}

func (h *SchedulersHandler) GetScheduler(ctx context.Context, request *api.GetSchedulerRequest) (*api.GetSchedulerResponse, error) {
	var scheduler *entities.Scheduler
	var err error
	scheduler, err = h.schedulerManager.GetScheduler(ctx, request.GetSchedulerName(), request.GetVersion())

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &api.GetSchedulerResponse{Scheduler: h.fromEntitySchedulerToResponse(scheduler)}, nil
}

func (h *SchedulersHandler) GetSchedulerVersions(ctx context.Context, request *api.GetSchedulerVersionsRequest) (*api.GetSchedulerVersionsResponse, error) {
	versions, err := h.schedulerManager.GetSchedulerVersions(ctx, request.GetSchedulerName())

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &api.GetSchedulerVersionsResponse{Versions: h.fromEntitySchedulerVersionListToResponse(versions)}, nil
}

func (h *SchedulersHandler) CreateScheduler(ctx context.Context, request *api.CreateSchedulerRequest) (*api.CreateSchedulerResponse, error) {
	scheduler := h.fromApiCreateSchedulerRequestToEntity(request)

	scheduler, err := h.schedulerManager.CreateScheduler(ctx, scheduler)
	if errors.Is(err, portsErrors.ErrAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	if err != nil {
		switch err.(type) {
		case validator.ErrorMap:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}

	return &api.CreateSchedulerResponse{
		Scheduler: h.fromEntitySchedulerToResponse(scheduler),
	}, nil
}

func (h *SchedulersHandler) AddRooms(ctx context.Context, request *api.AddRoomsRequest) (*api.AddRoomsResponse, error) {

	operation, err := h.schedulerManager.AddRooms(ctx, request.GetSchedulerName(), request.GetAmount())
	if errors.Is(err, portsErrors.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &api.AddRoomsResponse{
		OperationId: operation.ID,
	}, nil
}

func (h *SchedulersHandler) RemoveRooms(ctx context.Context, request *api.RemoveRoomsRequest) (*api.RemoveRoomsResponse, error) {

	operation, err := h.schedulerManager.RemoveRooms(ctx, request.GetSchedulerName(), int(request.GetAmount()))
	if errors.Is(err, portsErrors.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &api.RemoveRoomsResponse{
		OperationId: operation.ID,
	}, nil
}

func (h *SchedulersHandler) UpdateScheduler(ctx context.Context, request *api.UpdateSchedulerRequest) (*api.UpdateSchedulerResponse, error) {
	scheduler := h.fromApiUpdateSchedulerRequestToEntity(request)

	operation, err := h.schedulerManager.CreateUpdateSchedulerOperation(ctx, scheduler)
	if errors.Is(err, portsErrors.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &api.UpdateSchedulerResponse{
		OperationId: operation.ID,
	}, nil
}

func (h *SchedulersHandler) fromApiCreateSchedulerRequestToEntity(request *api.CreateSchedulerRequest) *entities.Scheduler {
	return &entities.Scheduler{
		Name:     request.GetName(),
		Game:     request.GetGame(),
		State:    entities.StateCreating,
		MaxSurge: request.GetMaxSurge(),
		PortRange: &entities.PortRange{
			Start: request.GetPortRange().GetStart(),
			End:   request.GetPortRange().GetEnd(),
		},
		Spec: game_room.Spec{
			Version:                request.GetVersion(),
			TerminationGracePeriod: time.Duration(request.GetTerminationGracePeriod()),
			Affinity:               request.GetAffinity(),
			Toleration:             request.GetToleration(),
			Containers:             h.fromApiContainers(request.GetContainers()),
		},
	}
}

func (h *SchedulersHandler) fromEntitySchedulerToListResponse(entity *entities.Scheduler) *api.SchedulerWithoutSpec {
	return &api.SchedulerWithoutSpec{
		Name:      entity.Name,
		Game:      entity.Game,
		State:     entity.State,
		Version:   entity.Spec.Version,
		PortRange: getPortRange(entity.PortRange),
		CreatedAt: timestamppb.New(entity.CreatedAt),
		MaxSurge:  entity.MaxSurge,
	}
}

func (h *SchedulersHandler) fromApiUpdateSchedulerRequestToEntity(request *api.UpdateSchedulerRequest) *entities.Scheduler {
	return &entities.Scheduler{
		Name:     request.GetName(),
		Game:     request.GetGame(),
		MaxSurge: request.GetMaxSurge(),
		PortRange: &entities.PortRange{
			Start: request.GetPortRange().GetStart(),
			End:   request.GetPortRange().GetEnd(),
		},
		Spec: game_room.Spec{
			TerminationGracePeriod: time.Duration(request.GetTerminationGracePeriod()),
			Affinity:               request.GetAffinity(),
			Toleration:             request.GetToleration(),
			Containers:             h.fromApiContainers(request.GetContainers()),
		},
	}
}

func (h *SchedulersHandler) fromEntitySchedulerToResponse(entity *entities.Scheduler) *api.Scheduler {
	return &api.Scheduler{
		Name:      entity.Name,
		Game:      entity.Game,
		State:     entity.State,
		Version:   entity.Spec.Version,
		PortRange: getPortRange(entity.PortRange),
		CreatedAt: timestamppb.New(entity.CreatedAt),
		MaxSurge:  entity.MaxSurge,
		Spec:      getSpec(entity.Spec),
	}
}

func (h *SchedulersHandler) fromEntitySchedulerVersionListToResponse(entity []*entities.SchedulerVersion) []*api.SchedulerVersion {
	versions := make([]*api.SchedulerVersion, len(entity))
	for i, version := range entity {
		versions[i] = &api.SchedulerVersion{
			Version:   version.Version,
			CreatedAt: timestamppb.New(version.CreatedAt),
		}
	}
	return versions
}

func (h *SchedulersHandler) fromApiContainers(apiContainers []*api.Container) []game_room.Container {
	var containers []game_room.Container
	for _, apiContainer := range apiContainers {
		container := game_room.Container{
			Name:            apiContainer.GetName(),
			Image:           apiContainer.GetImage(),
			ImagePullPolicy: apiContainer.GetImagePullPolicy(),
			Command:         apiContainer.GetCommand(),
			Ports:           h.fromApiContainerPorts(apiContainer.GetPorts()),
			Environment:     h.fromApiContainerEnvironments(apiContainer.GetEnvironment()),
			Requests: game_room.ContainerResources{
				CPU:    apiContainer.GetRequests().GetCpu(),
				Memory: apiContainer.GetRequests().GetMemory(),
			},
			Limits: game_room.ContainerResources{
				CPU:    apiContainer.GetLimits().GetCpu(),
				Memory: apiContainer.GetLimits().GetMemory(),
			},
		}
		containers = append(containers, container)
	}

	return containers
}

func (h *SchedulersHandler) fromApiContainerPorts(apiPorts []*api.ContainerPort) []game_room.ContainerPort {
	var ports []game_room.ContainerPort
	for _, apiPort := range apiPorts {
		port := game_room.ContainerPort{
			Name:     apiPort.GetName(),
			Port:     int(apiPort.GetPort()),
			Protocol: apiPort.GetProtocol(),
			HostPort: int(apiPort.GetHostPort()),
		}
		ports = append(ports, port)
	}

	return ports
}

func (h *SchedulersHandler) fromApiContainerEnvironments(apiEnvironments []*api.ContainerEnvironment) []game_room.ContainerEnvironment {
	var environments []game_room.ContainerEnvironment
	for _, apiEnvironment := range apiEnvironments {
		environment := game_room.ContainerEnvironment{
			Name:  apiEnvironment.GetName(),
			Value: apiEnvironment.GetValue(),
		}
		environments = append(environments, environment)
	}

	return environments
}

func getPortRange(portRange *entities.PortRange) *api.PortRange {
	if portRange != nil {
		return &api.PortRange{
			Start: portRange.Start,
			End:   portRange.End,
		}
	}

	return nil
}

func getSpec(spec game_room.Spec) *api.Spec {
	if spec.Version != "" {
		return &api.Spec{
			Version:                spec.Version,
			Toleration:             spec.Toleration,
			Containers:             fromEntityContainerToApiContainer(spec.Containers),
			TerminationGracePeriod: int64(spec.TerminationGracePeriod),
			Affinity:               spec.Affinity,
		}
	}

	return nil
}

func fromEntityContainerToApiContainer(containers []game_room.Container) []*api.Container {
	var convertedContainers []*api.Container
	for _, container := range containers {
		convertedContainers = append(convertedContainers, &api.Container{
			Name:            container.Name,
			Image:           container.Image,
			ImagePullPolicy: container.ImagePullPolicy,
			Command:         container.Command,
			Environment:     fromEntityContainerEnvironmentToApiContainerEnvironment(container.Environment),
			Requests:        fromEntityContainerResourcesToApiContainerResources(container.Requests),
			Limits:          fromEntityContainerResourcesToApiContainerResources(container.Limits),
			Ports:           fromEntityContainerPortsToApiContainerPorts(container.Ports),
		})
	}
	return convertedContainers
}

func fromEntityContainerEnvironmentToApiContainerEnvironment(environments []game_room.ContainerEnvironment) []*api.ContainerEnvironment {
	var convertedContainerEnvironment []*api.ContainerEnvironment
	for _, environment := range environments {
		convertedContainerEnvironment = append(convertedContainerEnvironment, &api.ContainerEnvironment{
			Name:  environment.Name,
			Value: environment.Value,
		})
	}
	return convertedContainerEnvironment
}

func fromEntityContainerResourcesToApiContainerResources(resources game_room.ContainerResources) *api.ContainerResources {
	return &api.ContainerResources{
		Memory: resources.Memory,
		Cpu:    resources.CPU,
	}
}

func fromEntityContainerPortsToApiContainerPorts(ports []game_room.ContainerPort) []*api.ContainerPort {
	var convertedContainerPort []*api.ContainerPort
	for _, port := range ports {
		convertedContainerPort = append(convertedContainerPort, &api.ContainerPort{
			Name:     port.Name,
			Protocol: port.Protocol,
			Port:     int32(port.Port),
			HostPort: int32(port.HostPort),
		})
	}
	return convertedContainerPort
}