// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/smallretardedfish/go-chat/internal/domains/user (interfaces: UserCredentialsRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user_cred_repo "github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
)

// MockUserCredentialsRepo is a mock of UserCredentialsRepo interface.
type MockUserCredentialsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserCredentialsRepoMockRecorder
}

// MockUserCredentialsRepoMockRecorder is the mock recorder for MockUserCredentialsRepo.
type MockUserCredentialsRepoMockRecorder struct {
	mock *MockUserCredentialsRepo
}

// NewMockUserCredentialsRepo creates a new mock instance.
func NewMockUserCredentialsRepo(ctrl *gomock.Controller) *MockUserCredentialsRepo {
	mock := &MockUserCredentialsRepo{ctrl: ctrl}
	mock.recorder = &MockUserCredentialsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCredentialsRepo) EXPECT() *MockUserCredentialsRepoMockRecorder {
	return m.recorder
}

// CreateUserCredentials mocks base method.
func (m *MockUserCredentialsRepo) CreateUserCredentials(arg0 user_cred_repo.UserCredentials) (*user_cred_repo.UserCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserCredentials", arg0)
	ret0, _ := ret[0].(*user_cred_repo.UserCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserCredentials indicates an expected call of CreateUserCredentials.
func (mr *MockUserCredentialsRepoMockRecorder) CreateUserCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserCredentials", reflect.TypeOf((*MockUserCredentialsRepo)(nil).CreateUserCredentials), arg0)
}

// GetUserCredentials mocks base method.
func (m *MockUserCredentialsRepo) GetUserCredentials(arg0 string) (*user_cred_repo.UserCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCredentials", arg0)
	ret0, _ := ret[0].(*user_cred_repo.UserCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCredentials indicates an expected call of GetUserCredentials.
func (mr *MockUserCredentialsRepoMockRecorder) GetUserCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCredentials", reflect.TypeOf((*MockUserCredentialsRepo)(nil).GetUserCredentials), arg0)
}

// UpdateUserCredentials mocks base method.
func (m *MockUserCredentialsRepo) UpdateUserCredentials(arg0 user_cred_repo.UserCredentials) (*user_cred_repo.UserCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserCredentials", arg0)
	ret0, _ := ret[0].(*user_cred_repo.UserCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserCredentials indicates an expected call of UpdateUserCredentials.
func (mr *MockUserCredentialsRepoMockRecorder) UpdateUserCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserCredentials", reflect.TypeOf((*MockUserCredentialsRepo)(nil).UpdateUserCredentials), arg0)
}
