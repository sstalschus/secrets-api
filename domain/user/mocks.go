// Code generated by MockGen. DO NOT EDIT.
// Source: ./contracts.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	domain "github.com/SamStalschus/secrets-api/domain"
	errors "github.com/SamStalschus/secrets-api/infra/errors"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIService) CreateUser(ctx context.Context, user *domain.User) *errors.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*errors.Message)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIServiceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIService)(nil).CreateUser), ctx, user)
}

// GetUserByEmail mocks base method.
func (m *MockIService) GetUserByEmail(ctx context.Context, userEmail string) (*domain.User, *errors.Message) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, userEmail)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(*errors.Message)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockIServiceMockRecorder) GetUserByEmail(ctx, userEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockIService)(nil).GetUserByEmail), ctx, userEmail)
}
