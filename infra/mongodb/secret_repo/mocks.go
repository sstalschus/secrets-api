// Code generated by MockGen. DO NOT EDIT.
// Source: ./contracts.go

// Package secret_repo is a generated GoMock package.
package secret_repo

import (
	context "context"
	reflect "reflect"

	internal "github.com/SamStalschus/secrets-api/internal"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method.
func (m *MockIRepository) CreateSecret(ctx context.Context, secret *internal.Secret, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", ctx, secret, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockIRepositoryMockRecorder) CreateSecret(ctx, secret, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockIRepository)(nil).CreateSecret), ctx, secret, userID)
}

// GenerateID mocks base method.
func (m *MockIRepository) GenerateID() primitive.ObjectID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateID")
	ret0, _ := ret[0].(primitive.ObjectID)
	return ret0
}

// GenerateID indicates an expected call of GenerateID.
func (mr *MockIRepositoryMockRecorder) GenerateID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateID", reflect.TypeOf((*MockIRepository)(nil).GenerateID))
}