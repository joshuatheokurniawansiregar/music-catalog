// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go
//
// Generated by this command:
//
//	mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
//

// Package memberships is a generated GoMock package.
package memberships

import (
	reflect "reflect"

	memberships "github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
	gomock "go.uber.org/mock/gomock"
)

// Mockservice is a mock of service interface.
type Mockservice struct {
	ctrl     *gomock.Controller
	recorder *MockserviceMockRecorder
	isgomock struct{}
}

// MockserviceMockRecorder is the mock recorder for Mockservice.
type MockserviceMockRecorder struct {
	mock *Mockservice
}

// NewMockservice creates a new mock instance.
func NewMockservice(ctrl *gomock.Controller) *Mockservice {
	mock := &Mockservice{ctrl: ctrl}
	mock.recorder = &MockserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockservice) EXPECT() *MockserviceMockRecorder {
	return m.recorder
}

// SignUp mocks base method.
func (m *Mockservice) SignUp(request memberships.SignupRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockserviceMockRecorder) SignUp(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*Mockservice)(nil).SignUp), request)
}
