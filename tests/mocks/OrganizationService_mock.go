// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juicycleff/frank/organization (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/juicycleff/frank/ent"
	"github.com/juicycleff/frank/organization"
)

// MockOrganizationService is a mock of Service interface.
type MockOrganizationService struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationServiceMockRecorder
}

// MockOrganizationServiceMockRecorder is the mock recorder for MockOrganizationService.
type MockOrganizationServiceMockRecorder struct {
	mock *MockOrganizationService
}

// NewMockOrganizationService creates a new mock instance.
func NewMockOrganizationService(ctrl *gomock.Controller) *MockOrganizationService {
	mock := &MockOrganizationService{ctrl: ctrl}
	mock.recorder = &MockOrganizationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationService) EXPECT() *MockOrganizationServiceMockRecorder {
	return m.recorder
}

// AddMember mocks base method.
func (m *MockOrganizationService) AddMember(arg0 context.Context, arg1, arg2 string, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMember", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMember indicates an expected call of AddMember.
func (mr *MockOrganizationServiceMockRecorder) AddMember(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMember", reflect.TypeOf((*MockOrganizationService)(nil).AddMember), arg0, arg1, arg2, arg3)
}

// Create mocks base method.
func (m *MockOrganizationService) Create(arg0 context.Context, arg1 organization.CreateOrganizationInput) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrganizationServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganizationService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockOrganizationService) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationService)(nil).Delete), arg0, arg1)
}

// DisableFeature mocks base method.
func (m *MockOrganizationService) DisableFeature(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableFeature", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableFeature indicates an expected call of DisableFeature.
func (mr *MockOrganizationServiceMockRecorder) DisableFeature(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableFeature", reflect.TypeOf((*MockOrganizationService)(nil).DisableFeature), arg0, arg1, arg2)
}

// EnableFeature mocks base method.
func (m *MockOrganizationService) EnableFeature(arg0 context.Context, arg1, arg2 string, arg3 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableFeature", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableFeature indicates an expected call of EnableFeature.
func (mr *MockOrganizationServiceMockRecorder) EnableFeature(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableFeature", reflect.TypeOf((*MockOrganizationService)(nil).EnableFeature), arg0, arg1, arg2, arg3)
}

// Get mocks base method.
func (m *MockOrganizationService) Get(arg0 context.Context, arg1 string) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrganizationServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrganizationService)(nil).Get), arg0, arg1)
}

// GetByDomain mocks base method.
func (m *MockOrganizationService) GetByDomain(arg0 context.Context, arg1 string) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDomain", arg0, arg1)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDomain indicates an expected call of GetByDomain.
func (mr *MockOrganizationServiceMockRecorder) GetByDomain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDomain", reflect.TypeOf((*MockOrganizationService)(nil).GetByDomain), arg0, arg1)
}

// GetBySlug mocks base method.
func (m *MockOrganizationService) GetBySlug(arg0 context.Context, arg1 string) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySlug", arg0, arg1)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug.
func (mr *MockOrganizationServiceMockRecorder) GetBySlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockOrganizationService)(nil).GetBySlug), arg0, arg1)
}

// GetFeatures mocks base method.
func (m *MockOrganizationService) GetFeatures(arg0 context.Context, arg1 string) ([]*ent.OrganizationFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatures", arg0, arg1)
	ret0, _ := ret[0].([]*ent.OrganizationFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatures indicates an expected call of GetFeatures.
func (mr *MockOrganizationServiceMockRecorder) GetFeatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatures", reflect.TypeOf((*MockOrganizationService)(nil).GetFeatures), arg0, arg1)
}

// GetMembers mocks base method.
func (m *MockOrganizationService) GetMembers(arg0 context.Context, arg1 string, arg2 organization.ListParams) ([]*ent.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*ent.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockOrganizationServiceMockRecorder) GetMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockOrganizationService)(nil).GetMembers), arg0, arg1, arg2)
}

// IsFeatureEnabled mocks base method.
func (m *MockOrganizationService) IsFeatureEnabled(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFeatureEnabled", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFeatureEnabled indicates an expected call of IsFeatureEnabled.
func (mr *MockOrganizationServiceMockRecorder) IsFeatureEnabled(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFeatureEnabled", reflect.TypeOf((*MockOrganizationService)(nil).IsFeatureEnabled), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockOrganizationService) List(arg0 context.Context, arg1 organization.ListParams) ([]*ent.Organization, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*ent.Organization)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockOrganizationServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrganizationService)(nil).List), arg0, arg1)
}

// RemoveMember mocks base method.
func (m *MockOrganizationService) RemoveMember(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMember indicates an expected call of RemoveMember.
func (mr *MockOrganizationServiceMockRecorder) RemoveMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMember", reflect.TypeOf((*MockOrganizationService)(nil).RemoveMember), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockOrganizationService) Update(arg0 context.Context, arg1 string, arg2 organization.UpdateOrganizationInput) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationServiceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationService)(nil).Update), arg0, arg1, arg2)
}

// UpdateMember mocks base method.
func (m *MockOrganizationService) UpdateMember(arg0 context.Context, arg1, arg2 string, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMember", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMember indicates an expected call of UpdateMember.
func (mr *MockOrganizationServiceMockRecorder) UpdateMember(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMember", reflect.TypeOf((*MockOrganizationService)(nil).UpdateMember), arg0, arg1, arg2, arg3)
}
