// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juicycleff/frank/internal/auth/session (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	session "github.com/juicycleff/frank/internal/auth/session"
)

// MockSessionStore is a mock of Store interface.
type MockSessionStore struct {
	ctrl     *gomock.Controller
	recorder *MockSessionStoreMockRecorder
}

// MockSessionStoreMockRecorder is the mock recorder for MockSessionStore.
type MockSessionStoreMockRecorder struct {
	mock *MockSessionStore
}

// NewMockSessionStore creates a new mock instance.
func NewMockSessionStore(ctrl *gomock.Controller) *MockSessionStore {
	mock := &MockSessionStore{ctrl: ctrl}
	mock.recorder = &MockSessionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionStore) EXPECT() *MockSessionStoreMockRecorder {
	return m.recorder
}

// DeleteSession mocks base method.
func (m *MockSessionStore) DeleteSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionStoreMockRecorder) DeleteSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionStore)(nil).DeleteSession), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockSessionStore) GetSession(arg0 context.Context, arg1 string) (*session.SessionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(*session.SessionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockSessionStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockSessionStore)(nil).GetSession), arg0, arg1)
}

// StoreSession mocks base method.
func (m *MockSessionStore) StoreSession(arg0 context.Context, arg1 string, arg2 *session.SessionInfo, arg3 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSession", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreSession indicates an expected call of StoreSession.
func (mr *MockSessionStoreMockRecorder) StoreSession(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSession", reflect.TypeOf((*MockSessionStore)(nil).StoreSession), arg0, arg1, arg2, arg3)
}

// UpdateSession mocks base method.
func (m *MockSessionStore) UpdateSession(arg0 context.Context, arg1 string, arg2 *session.SessionInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSession indicates an expected call of UpdateSession.
func (mr *MockSessionStoreMockRecorder) UpdateSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockSessionStore)(nil).UpdateSession), arg0, arg1, arg2)
}
