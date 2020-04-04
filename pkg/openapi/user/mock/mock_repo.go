// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/openapi/user/repositry.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/goodrain/rainbond-operator/pkg/openapi/model"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateIfNotExist mocks base method
func (m *MockRepository) CreateIfNotExist(user *model.User) error {
	ret := m.ctrl.Call(m, "CreateIfNotExist", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIfNotExist indicates an expected call of CreateIfNotExist
func (mr *MockRepositoryMockRecorder) CreateIfNotExist(user interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIfNotExist", reflect.TypeOf((*MockRepository)(nil).CreateIfNotExist), user)
}

// GetByUsername mocks base method
func (m *MockRepository) GetByUsername(username string) (*model.User, error) {
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername
func (mr *MockRepositoryMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockRepository)(nil).GetByUsername), username)
}

// Listusers mocks base method
func (m *MockRepository) Listusers() ([]*model.User, error) {
	ret := m.ctrl.Call(m, "Listusers")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Listusers indicates an expected call of Listusers
func (mr *MockRepositoryMockRecorder) Listusers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listusers", reflect.TypeOf((*MockRepository)(nil).Listusers))
}
