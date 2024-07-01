// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/core/ports/room_ports.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	sync "sync"
	time "time"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/topfreegames/maestro/internal/core/entities"
	game_room "github.com/topfreegames/maestro/internal/core/entities/game_room"
	ports "github.com/topfreegames/maestro/internal/core/ports"
)

// MockRoomManager is a mock of RoomManager interface.
type MockRoomManager struct {
	ctrl     *gomock.Controller
	recorder *MockRoomManagerMockRecorder
}

// MockRoomManagerMockRecorder is the mock recorder for MockRoomManager.
type MockRoomManagerMockRecorder struct {
	mock *MockRoomManager
}

// NewMockRoomManager creates a new mock instance.
func NewMockRoomManager(ctrl *gomock.Controller) *MockRoomManager {
	mock := &MockRoomManager{ctrl: ctrl}
	mock.recorder = &MockRoomManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomManager) EXPECT() *MockRoomManagerMockRecorder {
	return m.recorder
}

// CleanRoomState mocks base method.
func (m *MockRoomManager) CleanRoomState(ctx context.Context, schedulerName, roomId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanRoomState", ctx, schedulerName, roomId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanRoomState indicates an expected call of CleanRoomState.
func (mr *MockRoomManagerMockRecorder) CleanRoomState(ctx, schedulerName, roomId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanRoomState", reflect.TypeOf((*MockRoomManager)(nil).CleanRoomState), ctx, schedulerName, roomId)
}

// CreateRoom mocks base method.
func (m *MockRoomManager) CreateRoom(ctx context.Context, scheduler entities.Scheduler, isValidationRoom bool) (*game_room.GameRoom, *game_room.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, scheduler, isValidationRoom)
	ret0, _ := ret[0].(*game_room.GameRoom)
	ret1, _ := ret[1].(*game_room.Instance)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockRoomManagerMockRecorder) CreateRoom(ctx, scheduler, isValidationRoom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockRoomManager)(nil).CreateRoom), ctx, scheduler, isValidationRoom)
}

// DeleteRoom mocks base method.
func (m *MockRoomManager) DeleteRoom(ctx context.Context, gameRoom *game_room.GameRoom, reason string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", ctx, gameRoom, reason)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom.
func (mr *MockRoomManagerMockRecorder) DeleteRoom(ctx, gameRoom, reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockRoomManager)(nil).DeleteRoom), ctx, gameRoom, reason)
}

// GetRoomInstance mocks base method.
func (m *MockRoomManager) GetRoomInstance(ctx context.Context, scheduler, roomID string) (*game_room.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomInstance", ctx, scheduler, roomID)
	ret0, _ := ret[0].(*game_room.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomInstance indicates an expected call of GetRoomInstance.
func (mr *MockRoomManagerMockRecorder) GetRoomInstance(ctx, scheduler, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomInstance", reflect.TypeOf((*MockRoomManager)(nil).GetRoomInstance), ctx, scheduler, roomID)
}

// ListRoomsWithDeletionPriority mocks base method.
func (m *MockRoomManager) ListRoomsWithDeletionPriority(ctx context.Context, schedulerName, ignoredVersion string, amount int, roomsBeingReplaced *sync.Map) ([]*game_room.GameRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoomsWithDeletionPriority", ctx, schedulerName, ignoredVersion, amount, roomsBeingReplaced)
	ret0, _ := ret[0].([]*game_room.GameRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoomsWithDeletionPriority indicates an expected call of ListRoomsWithDeletionPriority.
func (mr *MockRoomManagerMockRecorder) ListRoomsWithDeletionPriority(ctx, schedulerName, ignoredVersion, amount, roomsBeingReplaced interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoomsWithDeletionPriority", reflect.TypeOf((*MockRoomManager)(nil).ListRoomsWithDeletionPriority), ctx, schedulerName, ignoredVersion, amount, roomsBeingReplaced)
}

// SchedulerMaxSurge mocks base method.
func (m *MockRoomManager) SchedulerMaxSurge(ctx context.Context, scheduler *entities.Scheduler) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulerMaxSurge", ctx, scheduler)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchedulerMaxSurge indicates an expected call of SchedulerMaxSurge.
func (mr *MockRoomManagerMockRecorder) SchedulerMaxSurge(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulerMaxSurge", reflect.TypeOf((*MockRoomManager)(nil).SchedulerMaxSurge), ctx, scheduler)
}

// UpdateGameRoomStatus mocks base method.
func (m *MockRoomManager) UpdateGameRoomStatus(ctx context.Context, schedulerId, gameRoomId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGameRoomStatus", ctx, schedulerId, gameRoomId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGameRoomStatus indicates an expected call of UpdateGameRoomStatus.
func (mr *MockRoomManagerMockRecorder) UpdateGameRoomStatus(ctx, schedulerId, gameRoomId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGameRoomStatus", reflect.TypeOf((*MockRoomManager)(nil).UpdateGameRoomStatus), ctx, schedulerId, gameRoomId)
}

// UpdateRoom mocks base method.
func (m *MockRoomManager) UpdateRoom(ctx context.Context, gameRoom *game_room.GameRoom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", ctx, gameRoom)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoom indicates an expected call of UpdateRoom.
func (mr *MockRoomManagerMockRecorder) UpdateRoom(ctx, gameRoom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockRoomManager)(nil).UpdateRoom), ctx, gameRoom)
}

// UpdateRoomInstance mocks base method.
func (m *MockRoomManager) UpdateRoomInstance(ctx context.Context, gameRoomInstance *game_room.Instance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoomInstance", ctx, gameRoomInstance)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoomInstance indicates an expected call of UpdateRoomInstance.
func (mr *MockRoomManagerMockRecorder) UpdateRoomInstance(ctx, gameRoomInstance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoomInstance", reflect.TypeOf((*MockRoomManager)(nil).UpdateRoomInstance), ctx, gameRoomInstance)
}

// WaitRoomStatus mocks base method.
func (m *MockRoomManager) WaitRoomStatus(ctx context.Context, gameRoom *game_room.GameRoom, status []game_room.GameRoomStatus) (game_room.GameRoomStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitRoomStatus", ctx, gameRoom, status)
	ret0, _ := ret[0].(game_room.GameRoomStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitRoomStatus indicates an expected call of WaitRoomStatus.
func (mr *MockRoomManagerMockRecorder) WaitRoomStatus(ctx, gameRoom, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitRoomStatus", reflect.TypeOf((*MockRoomManager)(nil).WaitRoomStatus), ctx, gameRoom, status)
}

// MockRoomStorage is a mock of RoomStorage interface.
type MockRoomStorage struct {
	ctrl     *gomock.Controller
	recorder *MockRoomStorageMockRecorder
}

// MockRoomStorageMockRecorder is the mock recorder for MockRoomStorage.
type MockRoomStorageMockRecorder struct {
	mock *MockRoomStorage
}

// NewMockRoomStorage creates a new mock instance.
func NewMockRoomStorage(ctrl *gomock.Controller) *MockRoomStorage {
	mock := &MockRoomStorage{ctrl: ctrl}
	mock.recorder = &MockRoomStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomStorage) EXPECT() *MockRoomStorageMockRecorder {
	return m.recorder
}

// CreateRoom mocks base method.
func (m *MockRoomStorage) CreateRoom(ctx context.Context, room *game_room.GameRoom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, room)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockRoomStorageMockRecorder) CreateRoom(ctx, room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockRoomStorage)(nil).CreateRoom), ctx, room)
}

// DeleteRoom mocks base method.
func (m *MockRoomStorage) DeleteRoom(ctx context.Context, scheduler, roomID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", ctx, scheduler, roomID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom.
func (mr *MockRoomStorageMockRecorder) DeleteRoom(ctx, scheduler, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockRoomStorage)(nil).DeleteRoom), ctx, scheduler, roomID)
}

// GetAllRoomIDs mocks base method.
func (m *MockRoomStorage) GetAllRoomIDs(ctx context.Context, scheduler string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRoomIDs", ctx, scheduler)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRoomIDs indicates an expected call of GetAllRoomIDs.
func (mr *MockRoomStorageMockRecorder) GetAllRoomIDs(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoomIDs", reflect.TypeOf((*MockRoomStorage)(nil).GetAllRoomIDs), ctx, scheduler)
}

// GetRoom mocks base method.
func (m *MockRoomStorage) GetRoom(ctx context.Context, scheduler, roomID string) (*game_room.GameRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, scheduler, roomID)
	ret0, _ := ret[0].(*game_room.GameRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockRoomStorageMockRecorder) GetRoom(ctx, scheduler, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockRoomStorage)(nil).GetRoom), ctx, scheduler, roomID)
}

// GetRoomCount mocks base method.
func (m *MockRoomStorage) GetRoomCount(ctx context.Context, scheduler string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomCount", ctx, scheduler)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomCount indicates an expected call of GetRoomCount.
func (mr *MockRoomStorageMockRecorder) GetRoomCount(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomCount", reflect.TypeOf((*MockRoomStorage)(nil).GetRoomCount), ctx, scheduler)
}

// GetRoomCountByStatus mocks base method.
func (m *MockRoomStorage) GetRoomCountByStatus(ctx context.Context, scheduler string, status game_room.GameRoomStatus) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomCountByStatus", ctx, scheduler, status)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomCountByStatus indicates an expected call of GetRoomCountByStatus.
func (mr *MockRoomStorageMockRecorder) GetRoomCountByStatus(ctx, scheduler, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomCountByStatus", reflect.TypeOf((*MockRoomStorage)(nil).GetRoomCountByStatus), ctx, scheduler, status)
}

// GetRoomIDsByLastPing mocks base method.
func (m *MockRoomStorage) GetRoomIDsByLastPing(ctx context.Context, scheduler string, threshold time.Time) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomIDsByLastPing", ctx, scheduler, threshold)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomIDsByLastPing indicates an expected call of GetRoomIDsByLastPing.
func (mr *MockRoomStorageMockRecorder) GetRoomIDsByLastPing(ctx, scheduler, threshold interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomIDsByLastPing", reflect.TypeOf((*MockRoomStorage)(nil).GetRoomIDsByLastPing), ctx, scheduler, threshold)
}

// GetRoomIDsByStatus mocks base method.
func (m *MockRoomStorage) GetRoomIDsByStatus(ctx context.Context, scheduler string, status game_room.GameRoomStatus) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomIDsByStatus", ctx, scheduler, status)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomIDsByStatus indicates an expected call of GetRoomIDsByStatus.
func (mr *MockRoomStorageMockRecorder) GetRoomIDsByStatus(ctx, scheduler, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomIDsByStatus", reflect.TypeOf((*MockRoomStorage)(nil).GetRoomIDsByStatus), ctx, scheduler, status)
}

// UpdateRoom mocks base method.
func (m *MockRoomStorage) UpdateRoom(ctx context.Context, room *game_room.GameRoom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", ctx, room)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoom indicates an expected call of UpdateRoom.
func (mr *MockRoomStorageMockRecorder) UpdateRoom(ctx, room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockRoomStorage)(nil).UpdateRoom), ctx, room)
}

// UpdateRoomStatus mocks base method.
func (m *MockRoomStorage) UpdateRoomStatus(ctx context.Context, scheduler, roomId string, status game_room.GameRoomStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoomStatus", ctx, scheduler, roomId, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoomStatus indicates an expected call of UpdateRoomStatus.
func (mr *MockRoomStorageMockRecorder) UpdateRoomStatus(ctx, scheduler, roomId, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoomStatus", reflect.TypeOf((*MockRoomStorage)(nil).UpdateRoomStatus), ctx, scheduler, roomId, status)
}

// WatchRoomStatus mocks base method.
func (m *MockRoomStorage) WatchRoomStatus(ctx context.Context, room *game_room.GameRoom) (ports.RoomStorageStatusWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRoomStatus", ctx, room)
	ret0, _ := ret[0].(ports.RoomStorageStatusWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRoomStatus indicates an expected call of WatchRoomStatus.
func (mr *MockRoomStorageMockRecorder) WatchRoomStatus(ctx, room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRoomStatus", reflect.TypeOf((*MockRoomStorage)(nil).WatchRoomStatus), ctx, room)
}

// MockRoomStorageStatusWatcher is a mock of RoomStorageStatusWatcher interface.
type MockRoomStorageStatusWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockRoomStorageStatusWatcherMockRecorder
}

// MockRoomStorageStatusWatcherMockRecorder is the mock recorder for MockRoomStorageStatusWatcher.
type MockRoomStorageStatusWatcherMockRecorder struct {
	mock *MockRoomStorageStatusWatcher
}

// NewMockRoomStorageStatusWatcher creates a new mock instance.
func NewMockRoomStorageStatusWatcher(ctrl *gomock.Controller) *MockRoomStorageStatusWatcher {
	mock := &MockRoomStorageStatusWatcher{ctrl: ctrl}
	mock.recorder = &MockRoomStorageStatusWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomStorageStatusWatcher) EXPECT() *MockRoomStorageStatusWatcherMockRecorder {
	return m.recorder
}

// ResultChan mocks base method.
func (m *MockRoomStorageStatusWatcher) ResultChan() chan game_room.StatusEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResultChan")
	ret0, _ := ret[0].(chan game_room.StatusEvent)
	return ret0
}

// ResultChan indicates an expected call of ResultChan.
func (mr *MockRoomStorageStatusWatcherMockRecorder) ResultChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResultChan", reflect.TypeOf((*MockRoomStorageStatusWatcher)(nil).ResultChan))
}

// Stop mocks base method.
func (m *MockRoomStorageStatusWatcher) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockRoomStorageStatusWatcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRoomStorageStatusWatcher)(nil).Stop))
}
