// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juicycleff/frank/internal/auth/passwordless (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/juicycleff/frank/ent"
	passwordless "github.com/juicycleff/frank/internal/auth/passwordless"
)

// MockPasswordlessService is a mock of Service interface.
type MockPasswordlessService struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordlessServiceMockRecorder
}

// MockPasswordlessServiceMockRecorder is the mock recorder for MockPasswordlessService.
type MockPasswordlessServiceMockRecorder struct {
	mock *MockPasswordlessService
}

// NewMockPasswordlessService creates a new mock instance.
func NewMockPasswordlessService(ctrl *gomock.Controller) *MockPasswordlessService {
	mock := &MockPasswordlessService{ctrl: ctrl}
	mock.recorder = &MockPasswordlessServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordlessService) EXPECT() *MockPasswordlessServiceMockRecorder {
	return m.recorder
}

// GenerateMagicLink mocks base method.
func (m *MockPasswordlessService) GenerateMagicLink(arg0 context.Context, arg1, arg2, arg3 string, arg4 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateMagicLink", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateMagicLink indicates an expected call of GenerateMagicLink.
func (mr *MockPasswordlessServiceMockRecorder) GenerateMagicLink(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateMagicLink", reflect.TypeOf((*MockPasswordlessService)(nil).GenerateMagicLink), arg0, arg1, arg2, arg3, arg4)
}

// GetMagicLinkData mocks base method.
func (m *MockPasswordlessService) GetMagicLinkData(arg0 context.Context, arg1 string) (*ent.Verification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMagicLinkData", arg0, arg1)
	ret0, _ := ret[0].(*ent.Verification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMagicLinkData indicates an expected call of GetMagicLinkData.
func (mr *MockPasswordlessServiceMockRecorder) GetMagicLinkData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMagicLinkData", reflect.TypeOf((*MockPasswordlessService)(nil).GetMagicLinkData), arg0, arg1)
}

// GetSupportedMethods mocks base method.
func (m *MockPasswordlessService) GetSupportedMethods() []passwordless.AuthType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportedMethods")
	ret0, _ := ret[0].([]passwordless.AuthType)
	return ret0
}

// GetSupportedMethods indicates an expected call of GetSupportedMethods.
func (mr *MockPasswordlessServiceMockRecorder) GetSupportedMethods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedMethods", reflect.TypeOf((*MockPasswordlessService)(nil).GetSupportedMethods))
}

// HandleDeprecatedMethods mocks base method.
func (m *MockPasswordlessService) HandleDeprecatedMethods() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDeprecatedMethods")
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleDeprecatedMethods indicates an expected call of HandleDeprecatedMethods.
func (mr *MockPasswordlessServiceMockRecorder) HandleDeprecatedMethods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDeprecatedMethods", reflect.TypeOf((*MockPasswordlessService)(nil).HandleDeprecatedMethods))
}

// InvalidateUserMagicLinks mocks base method.
func (m *MockPasswordlessService) InvalidateUserMagicLinks(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateUserMagicLinks", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateUserMagicLinks indicates an expected call of InvalidateUserMagicLinks.
func (mr *MockPasswordlessServiceMockRecorder) InvalidateUserMagicLinks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateUserMagicLinks", reflect.TypeOf((*MockPasswordlessService)(nil).InvalidateUserMagicLinks), arg0, arg1)
}

// IsConfigured mocks base method.
func (m *MockPasswordlessService) IsConfigured() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConfigured")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConfigured indicates an expected call of IsConfigured.
func (mr *MockPasswordlessServiceMockRecorder) IsConfigured() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConfigured", reflect.TypeOf((*MockPasswordlessService)(nil).IsConfigured))
}

// Login mocks base method.
func (m *MockPasswordlessService) Login(arg0 context.Context, arg1 passwordless.LoginRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockPasswordlessServiceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockPasswordlessService)(nil).Login), arg0, arg1)
}

// MarkMagicLinkAsUsed mocks base method.
func (m *MockPasswordlessService) MarkMagicLinkAsUsed(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkMagicLinkAsUsed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkMagicLinkAsUsed indicates an expected call of MarkMagicLinkAsUsed.
func (mr *MockPasswordlessServiceMockRecorder) MarkMagicLinkAsUsed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkMagicLinkAsUsed", reflect.TypeOf((*MockPasswordlessService)(nil).MarkMagicLinkAsUsed), arg0, arg1)
}

// VerifyLogin mocks base method.
func (m *MockPasswordlessService) VerifyLogin(arg0 context.Context, arg1 passwordless.VerifyRequest) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyLogin", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VerifyLogin indicates an expected call of VerifyLogin.
func (mr *MockPasswordlessServiceMockRecorder) VerifyLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyLogin", reflect.TypeOf((*MockPasswordlessService)(nil).VerifyLogin), arg0, arg1)
}

// VerifyPhoneOTP mocks base method.
func (m *MockPasswordlessService) VerifyPhoneOTP(arg0 context.Context, arg1, arg2, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPhoneOTP", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyPhoneOTP indicates an expected call of VerifyPhoneOTP.
func (mr *MockPasswordlessServiceMockRecorder) VerifyPhoneOTP(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPhoneOTP", reflect.TypeOf((*MockPasswordlessService)(nil).VerifyPhoneOTP), arg0, arg1, arg2, arg3)
}
