// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juicycleff/frank/internal/auth/sso (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/juicycleff/frank/ent"
	sso "github.com/juicycleff/frank/internal/auth/sso"
)

// MockSsoService is a mock of Service interface.
type MockSsoService struct {
	ctrl     *gomock.Controller
	recorder *MockSsoServiceMockRecorder
}

// MockSsoServiceMockRecorder is the mock recorder for MockSsoService.
type MockSsoServiceMockRecorder struct {
	mock *MockSsoService
}

// NewMockSsoService creates a new mock instance.
func NewMockSsoService(ctrl *gomock.Controller) *MockSsoService {
	mock := &MockSsoService{ctrl: ctrl}
	mock.recorder = &MockSsoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSsoService) EXPECT() *MockSsoServiceMockRecorder {
	return m.recorder
}

// CompleteSSO mocks base method.
func (m *MockSsoService) CompleteSSO(arg0 context.Context, arg1, arg2 string) (*sso.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteSSO", arg0, arg1, arg2)
	ret0, _ := ret[0].(*sso.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteSSO indicates an expected call of CompleteSSO.
func (mr *MockSsoServiceMockRecorder) CompleteSSO(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteSSO", reflect.TypeOf((*MockSsoService)(nil).CompleteSSO), arg0, arg1, arg2)
}

// CreateIdentityProvider mocks base method.
func (m *MockSsoService) CreateIdentityProvider(arg0 context.Context, arg1 *sso.IdentityProviderInput, arg2 string) (*ent.IdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIdentityProvider", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ent.IdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIdentityProvider indicates an expected call of CreateIdentityProvider.
func (mr *MockSsoServiceMockRecorder) CreateIdentityProvider(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIdentityProvider", reflect.TypeOf((*MockSsoService)(nil).CreateIdentityProvider), arg0, arg1, arg2)
}

// DeleteIdentityProvider mocks base method.
func (m *MockSsoService) DeleteIdentityProvider(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIdentityProvider", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIdentityProvider indicates an expected call of DeleteIdentityProvider.
func (mr *MockSsoServiceMockRecorder) DeleteIdentityProvider(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIdentityProvider", reflect.TypeOf((*MockSsoService)(nil).DeleteIdentityProvider), arg0, arg1)
}

// FindOrCreateUser mocks base method.
func (m *MockSsoService) FindOrCreateUser(arg0 context.Context, arg1 *sso.UserInfo) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrCreateUser", arg0, arg1)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrCreateUser indicates an expected call of FindOrCreateUser.
func (mr *MockSsoServiceMockRecorder) FindOrCreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrCreateUser", reflect.TypeOf((*MockSsoService)(nil).FindOrCreateUser), arg0, arg1)
}

// GetProvider mocks base method.
func (m *MockSsoService) GetProvider(arg0 context.Context, arg1 string) (*ent.IdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProvider", arg0, arg1)
	ret0, _ := ret[0].(*ent.IdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProvider indicates an expected call of GetProvider.
func (mr *MockSsoServiceMockRecorder) GetProvider(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProvider", reflect.TypeOf((*MockSsoService)(nil).GetProvider), arg0, arg1)
}

// GetProviders mocks base method.
func (m *MockSsoService) GetProviders(arg0 context.Context, arg1 string) ([]*ent.IdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProviders", arg0, arg1)
	ret0, _ := ret[0].([]*ent.IdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProviders indicates an expected call of GetProviders.
func (mr *MockSsoServiceMockRecorder) GetProviders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProviders", reflect.TypeOf((*MockSsoService)(nil).GetProviders), arg0, arg1)
}

// Initialize mocks base method.
func (m *MockSsoService) Initialize(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockSsoServiceMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockSsoService)(nil).Initialize), arg0)
}

// InitiateSSO mocks base method.
func (m *MockSsoService) InitiateSSO(arg0 context.Context, arg1, arg2 string, arg3 map[string]interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiateSSO", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitiateSSO indicates an expected call of InitiateSSO.
func (mr *MockSsoServiceMockRecorder) InitiateSSO(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiateSSO", reflect.TypeOf((*MockSsoService)(nil).InitiateSSO), arg0, arg1, arg2, arg3)
}

// UpdateIdentityProvider mocks base method.
func (m *MockSsoService) UpdateIdentityProvider(arg0 context.Context, arg1 string, arg2 *sso.IdentityProviderInput) (*ent.IdentityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIdentityProvider", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ent.IdentityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIdentityProvider indicates an expected call of UpdateIdentityProvider.
func (mr *MockSsoServiceMockRecorder) UpdateIdentityProvider(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIdentityProvider", reflect.TypeOf((*MockSsoService)(nil).UpdateIdentityProvider), arg0, arg1, arg2)
}

// ValidateAndRefreshToken mocks base method.
func (m *MockSsoService) ValidateAndRefreshToken(arg0 context.Context, arg1, arg2 string) (*sso.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAndRefreshToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(*sso.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateAndRefreshToken indicates an expected call of ValidateAndRefreshToken.
func (mr *MockSsoServiceMockRecorder) ValidateAndRefreshToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAndRefreshToken", reflect.TypeOf((*MockSsoService)(nil).ValidateAndRefreshToken), arg0, arg1, arg2)
}
