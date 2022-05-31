// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/core/ports/autoscaler/policy.go

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
