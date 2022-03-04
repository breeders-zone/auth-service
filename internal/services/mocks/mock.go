// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/services.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	domain "github.com/breeders-zone/auth-service/internal/domain"
	services "github.com/breeders-zone/auth-service/internal/services"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// FirstOrCreateByEmail mocks base method.
func (m *MockUserService) FirstOrCreateByEmail(input services.FirstOrCreateByEmailInput) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstOrCreateByEmail", input)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstOrCreateByEmail indicates an expected call of FirstOrCreateByEmail.
func (mr *MockUserServiceMockRecorder) FirstOrCreateByEmail(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrCreateByEmail", reflect.TypeOf((*MockUserService)(nil).FirstOrCreateByEmail), input)
}

// Login mocks base method.
func (m *MockUserService) Login(input services.UserLoginInput) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", input)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), input)
}
