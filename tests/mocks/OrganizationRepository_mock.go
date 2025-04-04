// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juicycleff/frank/internal/organization (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/juicycleff/frank/ent"
	"github.com/juicycleff/frank/organization"
)

// MockOrganizationRepository is a mock of Repository interface.
type MockOrganizationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationRepositoryMockRecorder
}

// MockOrganizationRepositoryMockRecorder is the mock recorder for MockOrganizationRepository.
type MockOrganizationRepositoryMockRecorder struct {
	mock *MockOrganizationRepository
}

// NewMockOrganizationRepository creates a new mock instance.
func NewMockOrganizationRepository(ctrl *gomock.Controller) *MockOrganizationRepository {
	mock := &MockOrganizationRepository{ctrl: ctrl}
	mock.recorder = &MockOrganizationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationRepository) EXPECT() *MockOrganizationRepositoryMockRecorder {
	return m.recorder
}

// AddMember mocks base method.
func (m *MockOrganizationRepository) AddMember(arg0 context.Context, arg1, arg2 string, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMember", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMember indicates an expected call of AddMember.
func (mr *MockOrganizationRepositoryMockRecorder) AddMember(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMember", reflect.TypeOf((*MockOrganizationRepository)(nil).AddMember), arg0, arg1, arg2, arg3)
}

// Create mocks base method.
func (m *MockOrganizationRepository) Create(arg0 context.Context, arg1 organization.RepositoryCreateInput) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrganizationRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganizationRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockOrganizationRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationRepository)(nil).Delete), arg0, arg1)
}

// DisableFeature mocks base method.
func (m *MockOrganizationRepository) DisableFeature(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableFeature", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableFeature indicates an expected call of DisableFeature.
func (mr *MockOrganizationRepositoryMockRecorder) DisableFeature(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableFeature", reflect.TypeOf((*MockOrganizationRepository)(nil).DisableFeature), arg0, arg1, arg2)
}

// EnableFeature mocks base method.
func (m *MockOrganizationRepository) EnableFeature(arg0 context.Context, arg1, arg2 string, arg3 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableFeature", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableFeature indicates an expected call of EnableFeature.
func (mr *MockOrganizationRepositoryMockRecorder) EnableFeature(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableFeature", reflect.TypeOf((*MockOrganizationRepository)(nil).EnableFeature), arg0, arg1, arg2, arg3)
}

// GetByDomain mocks base method.
func (m *MockOrganizationRepository) GetByDomain(arg0 context.Context, arg1 string) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDomain", arg0, arg1)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDomain indicates an expected call of GetByDomain.
func (mr *MockOrganizationRepositoryMockRecorder) GetByDomain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDomain", reflect.TypeOf((*MockOrganizationRepository)(nil).GetByDomain), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockOrganizationRepository) GetByID(arg0 context.Context, arg1 string) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockOrganizationRepositoryMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockOrganizationRepository)(nil).GetByID), arg0, arg1)
}

// GetBySlug mocks base method.
func (m *MockOrganizationRepository) GetBySlug(arg0 context.Context, arg1 string) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySlug", arg0, arg1)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug.
func (mr *MockOrganizationRepositoryMockRecorder) GetBySlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockOrganizationRepository)(nil).GetBySlug), arg0, arg1)
}

// GetFeatures mocks base method.
func (m *MockOrganizationRepository) GetFeatures(arg0 context.Context, arg1 string) ([]*ent.OrganizationFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatures", arg0, arg1)
	ret0, _ := ret[0].([]*ent.OrganizationFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatures indicates an expected call of GetFeatures.
func (mr *MockOrganizationRepositoryMockRecorder) GetFeatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatures", reflect.TypeOf((*MockOrganizationRepository)(nil).GetFeatures), arg0, arg1)
}

// GetMembers mocks base method.
func (m *MockOrganizationRepository) GetMembers(arg0 context.Context, arg1 string, arg2 organization.RepositoryListInput) ([]*ent.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*ent.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockOrganizationRepositoryMockRecorder) GetMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockOrganizationRepository)(nil).GetMembers), arg0, arg1, arg2)
}

// IsFeatureEnabled mocks base method.
func (m *MockOrganizationRepository) IsFeatureEnabled(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFeatureEnabled", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFeatureEnabled indicates an expected call of IsFeatureEnabled.
func (mr *MockOrganizationRepositoryMockRecorder) IsFeatureEnabled(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFeatureEnabled", reflect.TypeOf((*MockOrganizationRepository)(nil).IsFeatureEnabled), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockOrganizationRepository) List(arg0 context.Context, arg1 organization.RepositoryListInput) ([]*ent.Organization, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*ent.Organization)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockOrganizationRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrganizationRepository)(nil).List), arg0, arg1)
}

// RemoveMember mocks base method.
func (m *MockOrganizationRepository) RemoveMember(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMember indicates an expected call of RemoveMember.
func (mr *MockOrganizationRepositoryMockRecorder) RemoveMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMember", reflect.TypeOf((*MockOrganizationRepository)(nil).RemoveMember), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockOrganizationRepository) Update(arg0 context.Context, arg1 string, arg2 organization.RepositoryUpdateInput) (*ent.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ent.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationRepositoryMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationRepository)(nil).Update), arg0, arg1, arg2)
}

// UpdateMember mocks base method.
func (m *MockOrganizationRepository) UpdateMember(arg0 context.Context, arg1, arg2 string, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMember", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMember indicates an expected call of UpdateMember.
func (mr *MockOrganizationRepositoryMockRecorder) UpdateMember(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMember", reflect.TypeOf((*MockOrganizationRepository)(nil).UpdateMember), arg0, arg1, arg2, arg3)
}
