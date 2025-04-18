// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juicycleff/frank/internal/apikeys (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/juicycleff/frank/ent"
	apikeys "github.com/juicycleff/frank/internal/apikeys"
)

// MockApikeysRepository is a mock of Repository interface.
type MockApikeysRepository struct {
	ctrl     *gomock.Controller
	recorder *MockApikeysRepositoryMockRecorder
}

// MockApikeysRepositoryMockRecorder is the mock recorder for MockApikeysRepository.
type MockApikeysRepositoryMockRecorder struct {
	mock *MockApikeysRepository
}

// NewMockApikeysRepository creates a new mock instance.
func NewMockApikeysRepository(ctrl *gomock.Controller) *MockApikeysRepository {
	mock := &MockApikeysRepository{ctrl: ctrl}
	mock.recorder = &MockApikeysRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApikeysRepository) EXPECT() *MockApikeysRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockApikeysRepository) Create(arg0 context.Context, arg1 apikeys.RepositoryCreateInput) (*ent.ApiKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*ent.ApiKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockApikeysRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockApikeysRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockApikeysRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockApikeysRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApikeysRepository)(nil).Delete), arg0, arg1)
}

// GetByHashedKey mocks base method.
func (m *MockApikeysRepository) GetByHashedKey(arg0 context.Context, arg1 string) (*ent.ApiKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHashedKey", arg0, arg1)
	ret0, _ := ret[0].(*ent.ApiKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHashedKey indicates an expected call of GetByHashedKey.
func (mr *MockApikeysRepositoryMockRecorder) GetByHashedKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHashedKey", reflect.TypeOf((*MockApikeysRepository)(nil).GetByHashedKey), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockApikeysRepository) GetByID(arg0 context.Context, arg1 string) (*ent.ApiKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*ent.ApiKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockApikeysRepositoryMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockApikeysRepository)(nil).GetByID), arg0, arg1)
}

// List mocks base method.
func (m *MockApikeysRepository) List(arg0 context.Context, arg1 apikeys.RepositoryListInput) ([]*ent.ApiKey, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*ent.ApiKey)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockApikeysRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockApikeysRepository)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockApikeysRepository) Update(arg0 context.Context, arg1 string, arg2 apikeys.RepositoryUpdateInput) (*ent.ApiKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ent.ApiKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockApikeysRepositoryMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockApikeysRepository)(nil).Update), arg0, arg1, arg2)
}
