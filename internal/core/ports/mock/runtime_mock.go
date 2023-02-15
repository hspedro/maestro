// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/core/ports/runtime.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/topfreegames/maestro/internal/core/entities"
	game_room "github.com/topfreegames/maestro/internal/core/entities/game_room"
	ports "github.com/topfreegames/maestro/internal/core/ports"
)

// MockRuntime is a mock of Runtime interface.
type MockRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeMockRecorder
}

// MockRuntimeMockRecorder is the mock recorder for MockRuntime.
type MockRuntimeMockRecorder struct {
	mock *MockRuntime
}

// NewMockRuntime creates a new mock instance.
func NewMockRuntime(ctrl *gomock.Controller) *MockRuntime {
	mock := &MockRuntime{ctrl: ctrl}
	mock.recorder = &MockRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntime) EXPECT() *MockRuntimeMockRecorder {
	return m.recorder
}

// CreateGameRoomInstance mocks base method.
func (m *MockRuntime) CreateGameRoomInstance(ctx context.Context, schedulerId, gameRoomName string, spec game_room.Spec) (*game_room.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGameRoomInstance", ctx, schedulerId, gameRoomName, spec)
	ret0, _ := ret[0].(*game_room.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGameRoomInstance indicates an expected call of CreateGameRoomInstance.
func (mr *MockRuntimeMockRecorder) CreateGameRoomInstance(ctx, schedulerId, gameRoomName, spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGameRoomInstance", reflect.TypeOf((*MockRuntime)(nil).CreateGameRoomInstance), ctx, schedulerId, gameRoomName, spec)
}

// CreateGameRoomName mocks base method.
func (m *MockRuntime) CreateGameRoomName(ctx context.Context, scheduler entities.Scheduler) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGameRoomName", ctx, scheduler)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGameRoomName indicates an expected call of CreateGameRoomName.
func (mr *MockRuntimeMockRecorder) CreateGameRoomName(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGameRoomName", reflect.TypeOf((*MockRuntime)(nil).CreateGameRoomName), ctx, scheduler)
}

// CreateScheduler mocks base method.
func (m *MockRuntime) CreateScheduler(ctx context.Context, scheduler *entities.Scheduler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScheduler", ctx, scheduler)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateScheduler indicates an expected call of CreateScheduler.
func (mr *MockRuntimeMockRecorder) CreateScheduler(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScheduler", reflect.TypeOf((*MockRuntime)(nil).CreateScheduler), ctx, scheduler)
}

// DeleteGameRoomInstance mocks base method.
func (m *MockRuntime) DeleteGameRoomInstance(ctx context.Context, gameRoomInstance *game_room.Instance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGameRoomInstance", ctx, gameRoomInstance)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGameRoomInstance indicates an expected call of DeleteGameRoomInstance.
func (mr *MockRuntimeMockRecorder) DeleteGameRoomInstance(ctx, gameRoomInstance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGameRoomInstance", reflect.TypeOf((*MockRuntime)(nil).DeleteGameRoomInstance), ctx, gameRoomInstance)
}

// DeleteScheduler mocks base method.
func (m *MockRuntime) DeleteScheduler(ctx context.Context, scheduler *entities.Scheduler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScheduler", ctx, scheduler)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScheduler indicates an expected call of DeleteScheduler.
func (mr *MockRuntimeMockRecorder) DeleteScheduler(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduler", reflect.TypeOf((*MockRuntime)(nil).DeleteScheduler), ctx, scheduler)
}

// WatchGameRoomInstances mocks base method.
func (m *MockRuntime) WatchGameRoomInstances(ctx context.Context, scheduler *entities.Scheduler) (ports.RuntimeWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchGameRoomInstances", ctx, scheduler)
	ret0, _ := ret[0].(ports.RuntimeWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchGameRoomInstances indicates an expected call of WatchGameRoomInstances.
func (mr *MockRuntimeMockRecorder) WatchGameRoomInstances(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchGameRoomInstances", reflect.TypeOf((*MockRuntime)(nil).WatchGameRoomInstances), ctx, scheduler)
}

// MockRuntimeWatcher is a mock of RuntimeWatcher interface.
type MockRuntimeWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeWatcherMockRecorder
}

// MockRuntimeWatcherMockRecorder is the mock recorder for MockRuntimeWatcher.
type MockRuntimeWatcherMockRecorder struct {
	mock *MockRuntimeWatcher
}

// NewMockRuntimeWatcher creates a new mock instance.
func NewMockRuntimeWatcher(ctrl *gomock.Controller) *MockRuntimeWatcher {
	mock := &MockRuntimeWatcher{ctrl: ctrl}
	mock.recorder = &MockRuntimeWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntimeWatcher) EXPECT() *MockRuntimeWatcherMockRecorder {
	return m.recorder
}

// Err mocks base method.
func (m *MockRuntimeWatcher) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockRuntimeWatcherMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockRuntimeWatcher)(nil).Err))
}

// ResultChan mocks base method.
func (m *MockRuntimeWatcher) ResultChan() chan game_room.InstanceEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResultChan")
	ret0, _ := ret[0].(chan game_room.InstanceEvent)
	return ret0
}

// ResultChan indicates an expected call of ResultChan.
func (mr *MockRuntimeWatcherMockRecorder) ResultChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResultChan", reflect.TypeOf((*MockRuntimeWatcher)(nil).ResultChan))
}

// Stop mocks base method.
func (m *MockRuntimeWatcher) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockRuntimeWatcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRuntimeWatcher)(nil).Stop))
}