// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/operation_storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	operation "github.com/topfreegames/maestro/internal/core/entities/operation"
)

// MockOperationStorage is a mock of OperationStorage interface.
type MockOperationStorage struct {
	ctrl     *gomock.Controller
	recorder *MockOperationStorageMockRecorder
}

// MockOperationStorageMockRecorder is the mock recorder for MockOperationStorage.
type MockOperationStorageMockRecorder struct {
	mock *MockOperationStorage
}

// NewMockOperationStorage creates a new mock instance.
func NewMockOperationStorage(ctrl *gomock.Controller) *MockOperationStorage {
	mock := &MockOperationStorage{ctrl: ctrl}
	mock.recorder = &MockOperationStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationStorage) EXPECT() *MockOperationStorageMockRecorder {
	return m.recorder
}

// CreateOperation mocks base method.
func (m *MockOperationStorage) CreateOperation(ctx context.Context, operation *operation.Operation, definitionContent []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOperation", ctx, operation, definitionContent)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOperation indicates an expected call of CreateOperation.
func (mr *MockOperationStorageMockRecorder) CreateOperation(ctx, operation, definitionContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOperation", reflect.TypeOf((*MockOperationStorage)(nil).CreateOperation), ctx, operation, definitionContent)
}

// GetOperation mocks base method.
func (m *MockOperationStorage) GetOperation(ctx context.Context, schedulerName, operationID string) (*operation.Operation, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", ctx, schedulerName, operationID)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockOperationStorageMockRecorder) GetOperation(ctx, schedulerName, operationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockOperationStorage)(nil).GetOperation), ctx, schedulerName, operationID)
}

// ListSchedulerActiveOperations mocks base method.
func (m *MockOperationStorage) ListSchedulerActiveOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulerActiveOperations", ctx, schedulerName)
	ret0, _ := ret[0].([]*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulerActiveOperations indicates an expected call of ListSchedulerActiveOperations.
func (mr *MockOperationStorageMockRecorder) ListSchedulerActiveOperations(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulerActiveOperations", reflect.TypeOf((*MockOperationStorage)(nil).ListSchedulerActiveOperations), ctx, schedulerName)
}

// NextSchedulerOperationID mocks base method.
func (m *MockOperationStorage) NextSchedulerOperationID(ctx context.Context, schedulerName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextSchedulerOperationID", ctx, schedulerName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextSchedulerOperationID indicates an expected call of NextSchedulerOperationID.
func (mr *MockOperationStorageMockRecorder) NextSchedulerOperationID(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextSchedulerOperationID", reflect.TypeOf((*MockOperationStorage)(nil).NextSchedulerOperationID), ctx, schedulerName)
}

// SetOperationActive mocks base method.
func (m *MockOperationStorage) SetOperationActive(ctx context.Context, operation *operation.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOperationActive", ctx, operation)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOperationActive indicates an expected call of SetOperationActive.
func (mr *MockOperationStorageMockRecorder) SetOperationActive(ctx, operation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOperationActive", reflect.TypeOf((*MockOperationStorage)(nil).SetOperationActive), ctx, operation)
}