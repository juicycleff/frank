// Copyright 2023-present XRaph LLC. All rights reserved.
// This source code is licensed under the XRaph LLC license found
// in the LICENSE file in the root directory of this source tree.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xraph/frank/ent/oauthauthorization"
	"github.com/xraph/frank/ent/oauthclient"
	"github.com/xraph/frank/ent/oauthscope"
	"github.com/xraph/frank/ent/oauthtoken"
	"github.com/xraph/frank/ent/predicate"
	"github.com/rs/xid"
)

// OAuthScopeUpdate is the builder for updating OAuthScope entities.
type OAuthScopeUpdate struct {
	config
	hooks     []Hook
	mutation  *OAuthScopeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OAuthScopeUpdate builder.
func (osu *OAuthScopeUpdate) Where(ps ...predicate.OAuthScope) *OAuthScopeUpdate {
	osu.mutation.Where(ps...)
	return osu
}

// SetUpdatedAt sets the "updated_at" field.
func (osu *OAuthScopeUpdate) SetUpdatedAt(t time.Time) *OAuthScopeUpdate {
	osu.mutation.SetUpdatedAt(t)
	return osu
}

// SetDeletedAt sets the "deleted_at" field.
func (osu *OAuthScopeUpdate) SetDeletedAt(t time.Time) *OAuthScopeUpdate {
	osu.mutation.SetDeletedAt(t)
	return osu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (osu *OAuthScopeUpdate) SetNillableDeletedAt(t *time.Time) *OAuthScopeUpdate {
	if t != nil {
		osu.SetDeletedAt(*t)
	}
	return osu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (osu *OAuthScopeUpdate) ClearDeletedAt() *OAuthScopeUpdate {
	osu.mutation.ClearDeletedAt()
	return osu
}

// SetName sets the "name" field.
func (osu *OAuthScopeUpdate) SetName(s string) *OAuthScopeUpdate {
	osu.mutation.SetName(s)
	return osu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (osu *OAuthScopeUpdate) SetNillableName(s *string) *OAuthScopeUpdate {
	if s != nil {
		osu.SetName(*s)
	}
	return osu
}

// SetDescription sets the "description" field.
func (osu *OAuthScopeUpdate) SetDescription(s string) *OAuthScopeUpdate {
	osu.mutation.SetDescription(s)
	return osu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (osu *OAuthScopeUpdate) SetNillableDescription(s *string) *OAuthScopeUpdate {
	if s != nil {
		osu.SetDescription(*s)
	}
	return osu
}

// SetDefaultScope sets the "default_scope" field.
func (osu *OAuthScopeUpdate) SetDefaultScope(b bool) *OAuthScopeUpdate {
	osu.mutation.SetDefaultScope(b)
	return osu
}

// SetNillableDefaultScope sets the "default_scope" field if the given value is not nil.
func (osu *OAuthScopeUpdate) SetNillableDefaultScope(b *bool) *OAuthScopeUpdate {
	if b != nil {
		osu.SetDefaultScope(*b)
	}
	return osu
}

// SetPublic sets the "public" field.
func (osu *OAuthScopeUpdate) SetPublic(b bool) *OAuthScopeUpdate {
	osu.mutation.SetPublic(b)
	return osu
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (osu *OAuthScopeUpdate) SetNillablePublic(b *bool) *OAuthScopeUpdate {
	if b != nil {
		osu.SetPublic(*b)
	}
	return osu
}

// AddClientIDs adds the "clients" edge to the OAuthClient entity by IDs.
func (osu *OAuthScopeUpdate) AddClientIDs(ids ...xid.ID) *OAuthScopeUpdate {
	osu.mutation.AddClientIDs(ids...)
	return osu
}

// AddClients adds the "clients" edges to the OAuthClient entity.
func (osu *OAuthScopeUpdate) AddClients(o ...*OAuthClient) *OAuthScopeUpdate {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osu.AddClientIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the OAuthToken entity by IDs.
func (osu *OAuthScopeUpdate) AddTokenIDs(ids ...xid.ID) *OAuthScopeUpdate {
	osu.mutation.AddTokenIDs(ids...)
	return osu
}

// AddTokens adds the "tokens" edges to the OAuthToken entity.
func (osu *OAuthScopeUpdate) AddTokens(o ...*OAuthToken) *OAuthScopeUpdate {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osu.AddTokenIDs(ids...)
}

// AddAuthorizationIDs adds the "authorizations" edge to the OAuthAuthorization entity by IDs.
func (osu *OAuthScopeUpdate) AddAuthorizationIDs(ids ...xid.ID) *OAuthScopeUpdate {
	osu.mutation.AddAuthorizationIDs(ids...)
	return osu
}

// AddAuthorizations adds the "authorizations" edges to the OAuthAuthorization entity.
func (osu *OAuthScopeUpdate) AddAuthorizations(o ...*OAuthAuthorization) *OAuthScopeUpdate {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osu.AddAuthorizationIDs(ids...)
}

// Mutation returns the OAuthScopeMutation object of the builder.
func (osu *OAuthScopeUpdate) Mutation() *OAuthScopeMutation {
	return osu.mutation
}

// ClearClients clears all "clients" edges to the OAuthClient entity.
func (osu *OAuthScopeUpdate) ClearClients() *OAuthScopeUpdate {
	osu.mutation.ClearClients()
	return osu
}

// RemoveClientIDs removes the "clients" edge to OAuthClient entities by IDs.
func (osu *OAuthScopeUpdate) RemoveClientIDs(ids ...xid.ID) *OAuthScopeUpdate {
	osu.mutation.RemoveClientIDs(ids...)
	return osu
}

// RemoveClients removes "clients" edges to OAuthClient entities.
func (osu *OAuthScopeUpdate) RemoveClients(o ...*OAuthClient) *OAuthScopeUpdate {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osu.RemoveClientIDs(ids...)
}

// ClearTokens clears all "tokens" edges to the OAuthToken entity.
func (osu *OAuthScopeUpdate) ClearTokens() *OAuthScopeUpdate {
	osu.mutation.ClearTokens()
	return osu
}

// RemoveTokenIDs removes the "tokens" edge to OAuthToken entities by IDs.
func (osu *OAuthScopeUpdate) RemoveTokenIDs(ids ...xid.ID) *OAuthScopeUpdate {
	osu.mutation.RemoveTokenIDs(ids...)
	return osu
}

// RemoveTokens removes "tokens" edges to OAuthToken entities.
func (osu *OAuthScopeUpdate) RemoveTokens(o ...*OAuthToken) *OAuthScopeUpdate {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osu.RemoveTokenIDs(ids...)
}

// ClearAuthorizations clears all "authorizations" edges to the OAuthAuthorization entity.
func (osu *OAuthScopeUpdate) ClearAuthorizations() *OAuthScopeUpdate {
	osu.mutation.ClearAuthorizations()
	return osu
}

// RemoveAuthorizationIDs removes the "authorizations" edge to OAuthAuthorization entities by IDs.
func (osu *OAuthScopeUpdate) RemoveAuthorizationIDs(ids ...xid.ID) *OAuthScopeUpdate {
	osu.mutation.RemoveAuthorizationIDs(ids...)
	return osu
}

// RemoveAuthorizations removes "authorizations" edges to OAuthAuthorization entities.
func (osu *OAuthScopeUpdate) RemoveAuthorizations(o ...*OAuthAuthorization) *OAuthScopeUpdate {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osu.RemoveAuthorizationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *OAuthScopeUpdate) Save(ctx context.Context) (int, error) {
	osu.defaults()
	return withHooks(ctx, osu.sqlSave, osu.mutation, osu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (osu *OAuthScopeUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *OAuthScopeUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *OAuthScopeUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osu *OAuthScopeUpdate) defaults() {
	if _, ok := osu.mutation.UpdatedAt(); !ok {
		v := oauthscope.UpdateDefaultUpdatedAt()
		osu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osu *OAuthScopeUpdate) check() error {
	if v, ok := osu.mutation.Name(); ok {
		if err := oauthscope.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OAuthScope.name": %w`, err)}
		}
	}
	if v, ok := osu.mutation.Description(); ok {
		if err := oauthscope.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "OAuthScope.description": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (osu *OAuthScopeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OAuthScopeUpdate {
	osu.modifiers = append(osu.modifiers, modifiers...)
	return osu
}

func (osu *OAuthScopeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := osu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauthscope.Table, oauthscope.Columns, sqlgraph.NewFieldSpec(oauthscope.FieldID, field.TypeString))
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthscope.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := osu.mutation.DeletedAt(); ok {
		_spec.SetField(oauthscope.FieldDeletedAt, field.TypeTime, value)
	}
	if osu.mutation.DeletedAtCleared() {
		_spec.ClearField(oauthscope.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := osu.mutation.Name(); ok {
		_spec.SetField(oauthscope.FieldName, field.TypeString, value)
	}
	if value, ok := osu.mutation.Description(); ok {
		_spec.SetField(oauthscope.FieldDescription, field.TypeString, value)
	}
	if value, ok := osu.mutation.DefaultScope(); ok {
		_spec.SetField(oauthscope.FieldDefaultScope, field.TypeBool, value)
	}
	if value, ok := osu.mutation.Public(); ok {
		_spec.SetField(oauthscope.FieldPublic, field.TypeBool, value)
	}
	if osu.mutation.ClientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.ClientsTable,
			Columns: oauthscope.ClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.RemovedClientsIDs(); len(nodes) > 0 && !osu.mutation.ClientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.ClientsTable,
			Columns: oauthscope.ClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.ClientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.ClientsTable,
			Columns: oauthscope.ClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if osu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.TokensTable,
			Columns: oauthscope.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthtoken.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.RemovedTokensIDs(); len(nodes) > 0 && !osu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.TokensTable,
			Columns: oauthscope.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthtoken.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.TokensTable,
			Columns: oauthscope.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthtoken.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if osu.mutation.AuthorizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.AuthorizationsTable,
			Columns: oauthscope.AuthorizationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthauthorization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.RemovedAuthorizationsIDs(); len(nodes) > 0 && !osu.mutation.AuthorizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.AuthorizationsTable,
			Columns: oauthscope.AuthorizationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthauthorization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.AuthorizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.AuthorizationsTable,
			Columns: oauthscope.AuthorizationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthauthorization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(osu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthscope.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	osu.mutation.done = true
	return n, nil
}

// OAuthScopeUpdateOne is the builder for updating a single OAuthScope entity.
type OAuthScopeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OAuthScopeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (osuo *OAuthScopeUpdateOne) SetUpdatedAt(t time.Time) *OAuthScopeUpdateOne {
	osuo.mutation.SetUpdatedAt(t)
	return osuo
}

// SetDeletedAt sets the "deleted_at" field.
func (osuo *OAuthScopeUpdateOne) SetDeletedAt(t time.Time) *OAuthScopeUpdateOne {
	osuo.mutation.SetDeletedAt(t)
	return osuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (osuo *OAuthScopeUpdateOne) SetNillableDeletedAt(t *time.Time) *OAuthScopeUpdateOne {
	if t != nil {
		osuo.SetDeletedAt(*t)
	}
	return osuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (osuo *OAuthScopeUpdateOne) ClearDeletedAt() *OAuthScopeUpdateOne {
	osuo.mutation.ClearDeletedAt()
	return osuo
}

// SetName sets the "name" field.
func (osuo *OAuthScopeUpdateOne) SetName(s string) *OAuthScopeUpdateOne {
	osuo.mutation.SetName(s)
	return osuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (osuo *OAuthScopeUpdateOne) SetNillableName(s *string) *OAuthScopeUpdateOne {
	if s != nil {
		osuo.SetName(*s)
	}
	return osuo
}

// SetDescription sets the "description" field.
func (osuo *OAuthScopeUpdateOne) SetDescription(s string) *OAuthScopeUpdateOne {
	osuo.mutation.SetDescription(s)
	return osuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (osuo *OAuthScopeUpdateOne) SetNillableDescription(s *string) *OAuthScopeUpdateOne {
	if s != nil {
		osuo.SetDescription(*s)
	}
	return osuo
}

// SetDefaultScope sets the "default_scope" field.
func (osuo *OAuthScopeUpdateOne) SetDefaultScope(b bool) *OAuthScopeUpdateOne {
	osuo.mutation.SetDefaultScope(b)
	return osuo
}

// SetNillableDefaultScope sets the "default_scope" field if the given value is not nil.
func (osuo *OAuthScopeUpdateOne) SetNillableDefaultScope(b *bool) *OAuthScopeUpdateOne {
	if b != nil {
		osuo.SetDefaultScope(*b)
	}
	return osuo
}

// SetPublic sets the "public" field.
func (osuo *OAuthScopeUpdateOne) SetPublic(b bool) *OAuthScopeUpdateOne {
	osuo.mutation.SetPublic(b)
	return osuo
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (osuo *OAuthScopeUpdateOne) SetNillablePublic(b *bool) *OAuthScopeUpdateOne {
	if b != nil {
		osuo.SetPublic(*b)
	}
	return osuo
}

// AddClientIDs adds the "clients" edge to the OAuthClient entity by IDs.
func (osuo *OAuthScopeUpdateOne) AddClientIDs(ids ...xid.ID) *OAuthScopeUpdateOne {
	osuo.mutation.AddClientIDs(ids...)
	return osuo
}

// AddClients adds the "clients" edges to the OAuthClient entity.
func (osuo *OAuthScopeUpdateOne) AddClients(o ...*OAuthClient) *OAuthScopeUpdateOne {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osuo.AddClientIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the OAuthToken entity by IDs.
func (osuo *OAuthScopeUpdateOne) AddTokenIDs(ids ...xid.ID) *OAuthScopeUpdateOne {
	osuo.mutation.AddTokenIDs(ids...)
	return osuo
}

// AddTokens adds the "tokens" edges to the OAuthToken entity.
func (osuo *OAuthScopeUpdateOne) AddTokens(o ...*OAuthToken) *OAuthScopeUpdateOne {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osuo.AddTokenIDs(ids...)
}

// AddAuthorizationIDs adds the "authorizations" edge to the OAuthAuthorization entity by IDs.
func (osuo *OAuthScopeUpdateOne) AddAuthorizationIDs(ids ...xid.ID) *OAuthScopeUpdateOne {
	osuo.mutation.AddAuthorizationIDs(ids...)
	return osuo
}

// AddAuthorizations adds the "authorizations" edges to the OAuthAuthorization entity.
func (osuo *OAuthScopeUpdateOne) AddAuthorizations(o ...*OAuthAuthorization) *OAuthScopeUpdateOne {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osuo.AddAuthorizationIDs(ids...)
}

// Mutation returns the OAuthScopeMutation object of the builder.
func (osuo *OAuthScopeUpdateOne) Mutation() *OAuthScopeMutation {
	return osuo.mutation
}

// ClearClients clears all "clients" edges to the OAuthClient entity.
func (osuo *OAuthScopeUpdateOne) ClearClients() *OAuthScopeUpdateOne {
	osuo.mutation.ClearClients()
	return osuo
}

// RemoveClientIDs removes the "clients" edge to OAuthClient entities by IDs.
func (osuo *OAuthScopeUpdateOne) RemoveClientIDs(ids ...xid.ID) *OAuthScopeUpdateOne {
	osuo.mutation.RemoveClientIDs(ids...)
	return osuo
}

// RemoveClients removes "clients" edges to OAuthClient entities.
func (osuo *OAuthScopeUpdateOne) RemoveClients(o ...*OAuthClient) *OAuthScopeUpdateOne {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osuo.RemoveClientIDs(ids...)
}

// ClearTokens clears all "tokens" edges to the OAuthToken entity.
func (osuo *OAuthScopeUpdateOne) ClearTokens() *OAuthScopeUpdateOne {
	osuo.mutation.ClearTokens()
	return osuo
}

// RemoveTokenIDs removes the "tokens" edge to OAuthToken entities by IDs.
func (osuo *OAuthScopeUpdateOne) RemoveTokenIDs(ids ...xid.ID) *OAuthScopeUpdateOne {
	osuo.mutation.RemoveTokenIDs(ids...)
	return osuo
}

// RemoveTokens removes "tokens" edges to OAuthToken entities.
func (osuo *OAuthScopeUpdateOne) RemoveTokens(o ...*OAuthToken) *OAuthScopeUpdateOne {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osuo.RemoveTokenIDs(ids...)
}

// ClearAuthorizations clears all "authorizations" edges to the OAuthAuthorization entity.
func (osuo *OAuthScopeUpdateOne) ClearAuthorizations() *OAuthScopeUpdateOne {
	osuo.mutation.ClearAuthorizations()
	return osuo
}

// RemoveAuthorizationIDs removes the "authorizations" edge to OAuthAuthorization entities by IDs.
func (osuo *OAuthScopeUpdateOne) RemoveAuthorizationIDs(ids ...xid.ID) *OAuthScopeUpdateOne {
	osuo.mutation.RemoveAuthorizationIDs(ids...)
	return osuo
}

// RemoveAuthorizations removes "authorizations" edges to OAuthAuthorization entities.
func (osuo *OAuthScopeUpdateOne) RemoveAuthorizations(o ...*OAuthAuthorization) *OAuthScopeUpdateOne {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return osuo.RemoveAuthorizationIDs(ids...)
}

// Where appends a list predicates to the OAuthScopeUpdate builder.
func (osuo *OAuthScopeUpdateOne) Where(ps ...predicate.OAuthScope) *OAuthScopeUpdateOne {
	osuo.mutation.Where(ps...)
	return osuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *OAuthScopeUpdateOne) Select(field string, fields ...string) *OAuthScopeUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated OAuthScope entity.
func (osuo *OAuthScopeUpdateOne) Save(ctx context.Context) (*OAuthScope, error) {
	osuo.defaults()
	return withHooks(ctx, osuo.sqlSave, osuo.mutation, osuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *OAuthScopeUpdateOne) SaveX(ctx context.Context) *OAuthScope {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *OAuthScopeUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *OAuthScopeUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osuo *OAuthScopeUpdateOne) defaults() {
	if _, ok := osuo.mutation.UpdatedAt(); !ok {
		v := oauthscope.UpdateDefaultUpdatedAt()
		osuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osuo *OAuthScopeUpdateOne) check() error {
	if v, ok := osuo.mutation.Name(); ok {
		if err := oauthscope.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OAuthScope.name": %w`, err)}
		}
	}
	if v, ok := osuo.mutation.Description(); ok {
		if err := oauthscope.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "OAuthScope.description": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (osuo *OAuthScopeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OAuthScopeUpdateOne {
	osuo.modifiers = append(osuo.modifiers, modifiers...)
	return osuo
}

func (osuo *OAuthScopeUpdateOne) sqlSave(ctx context.Context) (_node *OAuthScope, err error) {
	if err := osuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauthscope.Table, oauthscope.Columns, sqlgraph.NewFieldSpec(oauthscope.FieldID, field.TypeString))
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OAuthScope.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauthscope.FieldID)
		for _, f := range fields {
			if !oauthscope.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != oauthscope.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthscope.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := osuo.mutation.DeletedAt(); ok {
		_spec.SetField(oauthscope.FieldDeletedAt, field.TypeTime, value)
	}
	if osuo.mutation.DeletedAtCleared() {
		_spec.ClearField(oauthscope.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := osuo.mutation.Name(); ok {
		_spec.SetField(oauthscope.FieldName, field.TypeString, value)
	}
	if value, ok := osuo.mutation.Description(); ok {
		_spec.SetField(oauthscope.FieldDescription, field.TypeString, value)
	}
	if value, ok := osuo.mutation.DefaultScope(); ok {
		_spec.SetField(oauthscope.FieldDefaultScope, field.TypeBool, value)
	}
	if value, ok := osuo.mutation.Public(); ok {
		_spec.SetField(oauthscope.FieldPublic, field.TypeBool, value)
	}
	if osuo.mutation.ClientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.ClientsTable,
			Columns: oauthscope.ClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.RemovedClientsIDs(); len(nodes) > 0 && !osuo.mutation.ClientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.ClientsTable,
			Columns: oauthscope.ClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.ClientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.ClientsTable,
			Columns: oauthscope.ClientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if osuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.TokensTable,
			Columns: oauthscope.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthtoken.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.RemovedTokensIDs(); len(nodes) > 0 && !osuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.TokensTable,
			Columns: oauthscope.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthtoken.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.TokensTable,
			Columns: oauthscope.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthtoken.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if osuo.mutation.AuthorizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.AuthorizationsTable,
			Columns: oauthscope.AuthorizationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthauthorization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.RemovedAuthorizationsIDs(); len(nodes) > 0 && !osuo.mutation.AuthorizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.AuthorizationsTable,
			Columns: oauthscope.AuthorizationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthauthorization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.AuthorizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oauthscope.AuthorizationsTable,
			Columns: oauthscope.AuthorizationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthauthorization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(osuo.modifiers...)
	_node = &OAuthScope{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthscope.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	osuo.mutation.done = true
	return _node, nil
}
