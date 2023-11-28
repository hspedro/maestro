// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/core/ports/autoscaler.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/topfreegames/maestro/internal/core/entities"
	autoscaling "github.com/topfreegames/maestro/internal/core/entities/autoscaling"
	policies "github.com/topfreegames/maestro/internal/core/services/autoscaler/policies"
)

// MockAutoscaler is a mock of Autoscaler interface.
type MockAutoscaler struct {
	ctrl     *gomock.Controller
	recorder *MockAutoscalerMockRecorder
}

// MockAutoscalerMockRecorder is the mock recorder for MockAutoscaler.
type MockAutoscalerMockRecorder struct {
	mock *MockAutoscaler
}

// NewMockAutoscaler creates a new mock instance.
func NewMockAutoscaler(ctrl *gomock.Controller) *MockAutoscaler {
	mock := &MockAutoscaler{ctrl: ctrl}
	mock.recorder = &MockAutoscalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutoscaler) EXPECT() *MockAutoscalerMockRecorder {
	return m.recorder
}

// CalculateDesiredNumberOfRooms mocks base method.
func (m *MockAutoscaler) CalculateDesiredNumberOfRooms(ctx context.Context, scheduler *entities.Scheduler) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateDesiredNumberOfRooms", ctx, scheduler)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateDesiredNumberOfRooms indicates an expected call of CalculateDesiredNumberOfRooms.
func (mr *MockAutoscalerMockRecorder) CalculateDesiredNumberOfRooms(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateDesiredNumberOfRooms", reflect.TypeOf((*MockAutoscaler)(nil).CalculateDesiredNumberOfRooms), ctx, scheduler)
}

// CanDownscale mocks base method.
func (m *MockAutoscaler) CanDownscale(ctx context.Context, scheduler *entities.Scheduler) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDownscale", ctx, scheduler)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanDownscale indicates an expected call of CanDownscale.
func (mr *MockAutoscalerMockRecorder) CanDownscale(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDownscale", reflect.TypeOf((*MockAutoscaler)(nil).CanDownscale), ctx, scheduler)
}

// MockPolicy is a mock of Policy interface.
type MockPolicy struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyMockRecorder
}

// MockPolicyMockRecorder is the mock recorder for MockPolicy.
type MockPolicyMockRecorder struct {
	mock *MockPolicy
}

// NewMockPolicy creates a new mock instance.
func NewMockPolicy(ctrl *gomock.Controller) *MockPolicy {
	mock := &MockPolicy{ctrl: ctrl}
	mock.recorder = &MockPolicyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicy) EXPECT() *MockPolicyMockRecorder {
	return m.recorder
}

// CalculateDesiredNumberOfRooms mocks base method.
func (m *MockPolicy) CalculateDesiredNumberOfRooms(policyParameters autoscaling.PolicyParameters, currentState policies.CurrentState) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateDesiredNumberOfRooms", policyParameters, currentState)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateDesiredNumberOfRooms indicates an expected call of CalculateDesiredNumberOfRooms.
func (mr *MockPolicyMockRecorder) CalculateDesiredNumberOfRooms(policyParameters, currentState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateDesiredNumberOfRooms", reflect.TypeOf((*MockPolicy)(nil).CalculateDesiredNumberOfRooms), policyParameters, currentState)
}

// CanDownscale mocks base method.
func (m *MockPolicy) CanDownscale(policyParameters autoscaling.PolicyParameters, currentState policies.CurrentState) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDownscale", policyParameters, currentState)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanDownscale indicates an expected call of CanDownscale.
func (mr *MockPolicyMockRecorder) CanDownscale(policyParameters, currentState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDownscale", reflect.TypeOf((*MockPolicy)(nil).CanDownscale), policyParameters, currentState)
}

// CurrentStateBuilder mocks base method.
func (m *MockPolicy) CurrentStateBuilder(ctx context.Context, scheduler *entities.Scheduler) (policies.CurrentState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentStateBuilder", ctx, scheduler)
	ret0, _ := ret[0].(policies.CurrentState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentStateBuilder indicates an expected call of CurrentStateBuilder.
func (mr *MockPolicyMockRecorder) CurrentStateBuilder(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentStateBuilder", reflect.TypeOf((*MockPolicy)(nil).CurrentStateBuilder), ctx, scheduler)
}
