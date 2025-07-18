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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xraph/frank/ent/mfa"
	"github.com/xraph/frank/ent/user"
	"github.com/rs/xid"
)

// MFACreate is the builder for creating a MFA entity.
type MFACreate struct {
	config
	mutation *MFAMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MFACreate) SetCreatedAt(t time.Time) *MFACreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MFACreate) SetNillableCreatedAt(t *time.Time) *MFACreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MFACreate) SetUpdatedAt(t time.Time) *MFACreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MFACreate) SetNillableUpdatedAt(t *time.Time) *MFACreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MFACreate) SetDeletedAt(t time.Time) *MFACreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MFACreate) SetNillableDeletedAt(t *time.Time) *MFACreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetUserID sets the "user_id" field.
func (mc *MFACreate) SetUserID(x xid.ID) *MFACreate {
	mc.mutation.SetUserID(x)
	return mc
}

// SetMethod sets the "method" field.
func (mc *MFACreate) SetMethod(s string) *MFACreate {
	mc.mutation.SetMethod(s)
	return mc
}

// SetSecret sets the "secret" field.
func (mc *MFACreate) SetSecret(s string) *MFACreate {
	mc.mutation.SetSecret(s)
	return mc
}

// SetVerified sets the "verified" field.
func (mc *MFACreate) SetVerified(b bool) *MFACreate {
	mc.mutation.SetVerified(b)
	return mc
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (mc *MFACreate) SetNillableVerified(b *bool) *MFACreate {
	if b != nil {
		mc.SetVerified(*b)
	}
	return mc
}

// SetActive sets the "active" field.
func (mc *MFACreate) SetActive(b bool) *MFACreate {
	mc.mutation.SetActive(b)
	return mc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (mc *MFACreate) SetNillableActive(b *bool) *MFACreate {
	if b != nil {
		mc.SetActive(*b)
	}
	return mc
}

// SetBackupCodes sets the "backup_codes" field.
func (mc *MFACreate) SetBackupCodes(s []string) *MFACreate {
	mc.mutation.SetBackupCodes(s)
	return mc
}

// SetPhoneNumber sets the "phone_number" field.
func (mc *MFACreate) SetPhoneNumber(s string) *MFACreate {
	mc.mutation.SetPhoneNumber(s)
	return mc
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (mc *MFACreate) SetNillablePhoneNumber(s *string) *MFACreate {
	if s != nil {
		mc.SetPhoneNumber(*s)
	}
	return mc
}

// SetEmail sets the "email" field.
func (mc *MFACreate) SetEmail(s string) *MFACreate {
	mc.mutation.SetEmail(s)
	return mc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mc *MFACreate) SetNillableEmail(s *string) *MFACreate {
	if s != nil {
		mc.SetEmail(*s)
	}
	return mc
}

// SetLastUsed sets the "last_used" field.
func (mc *MFACreate) SetLastUsed(t time.Time) *MFACreate {
	mc.mutation.SetLastUsed(t)
	return mc
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (mc *MFACreate) SetNillableLastUsed(t *time.Time) *MFACreate {
	if t != nil {
		mc.SetLastUsed(*t)
	}
	return mc
}

// SetMetadata sets the "metadata" field.
func (mc *MFACreate) SetMetadata(m map[string]interface{}) *MFACreate {
	mc.mutation.SetMetadata(m)
	return mc
}

// SetID sets the "id" field.
func (mc *MFACreate) SetID(x xid.ID) *MFACreate {
	mc.mutation.SetID(x)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MFACreate) SetNillableID(x *xid.ID) *MFACreate {
	if x != nil {
		mc.SetID(*x)
	}
	return mc
}

// SetUser sets the "user" edge to the User entity.
func (mc *MFACreate) SetUser(u *User) *MFACreate {
	return mc.SetUserID(u.ID)
}

// Mutation returns the MFAMutation object of the builder.
func (mc *MFACreate) Mutation() *MFAMutation {
	return mc.mutation
}

// Save creates the MFA in the database.
func (mc *MFACreate) Save(ctx context.Context) (*MFA, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MFACreate) SaveX(ctx context.Context) *MFA {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MFACreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MFACreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MFACreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := mfa.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := mfa.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Verified(); !ok {
		v := mfa.DefaultVerified
		mc.mutation.SetVerified(v)
	}
	if _, ok := mc.mutation.Active(); !ok {
		v := mfa.DefaultActive
		mc.mutation.SetActive(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := mfa.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MFACreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MFA.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MFA.updated_at"`)}
	}
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "MFA.user_id"`)}
	}
	if v, ok := mc.mutation.UserID(); ok {
		if err := mfa.UserIDValidator(v.String()); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "MFA.user_id": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "MFA.method"`)}
	}
	if v, ok := mc.mutation.Method(); ok {
		if err := mfa.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "MFA.method": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`ent: missing required field "MFA.secret"`)}
	}
	if v, ok := mc.mutation.Secret(); ok {
		if err := mfa.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`ent: validator failed for field "MFA.secret": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Verified(); !ok {
		return &ValidationError{Name: "verified", err: errors.New(`ent: missing required field "MFA.verified"`)}
	}
	if _, ok := mc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "MFA.active"`)}
	}
	if len(mc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "MFA.user"`)}
	}
	return nil
}

func (mc *MFACreate) sqlSave(ctx context.Context) (*MFA, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MFACreate) createSpec() (*MFA, *sqlgraph.CreateSpec) {
	var (
		_node = &MFA{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(mfa.Table, sqlgraph.NewFieldSpec(mfa.FieldID, field.TypeString))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(mfa.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(mfa.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.SetField(mfa.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := mc.mutation.Method(); ok {
		_spec.SetField(mfa.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	if value, ok := mc.mutation.Secret(); ok {
		_spec.SetField(mfa.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if value, ok := mc.mutation.Verified(); ok {
		_spec.SetField(mfa.FieldVerified, field.TypeBool, value)
		_node.Verified = value
	}
	if value, ok := mc.mutation.Active(); ok {
		_spec.SetField(mfa.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := mc.mutation.BackupCodes(); ok {
		_spec.SetField(mfa.FieldBackupCodes, field.TypeJSON, value)
		_node.BackupCodes = value
	}
	if value, ok := mc.mutation.PhoneNumber(); ok {
		_spec.SetField(mfa.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := mc.mutation.Email(); ok {
		_spec.SetField(mfa.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := mc.mutation.LastUsed(); ok {
		_spec.SetField(mfa.FieldLastUsed, field.TypeTime, value)
		_node.LastUsed = &value
	}
	if value, ok := mc.mutation.Metadata(); ok {
		_spec.SetField(mfa.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mfa.UserTable,
			Columns: []string{mfa.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MFA.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MFAUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mc *MFACreate) OnConflict(opts ...sql.ConflictOption) *MFAUpsertOne {
	mc.conflict = opts
	return &MFAUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MFA.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MFACreate) OnConflictColumns(columns ...string) *MFAUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MFAUpsertOne{
		create: mc,
	}
}

type (
	// MFAUpsertOne is the builder for "upsert"-ing
	//  one MFA node.
	MFAUpsertOne struct {
		create *MFACreate
	}

	// MFAUpsert is the "OnConflict" setter.
	MFAUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *MFAUpsert) SetUpdatedAt(v time.Time) *MFAUpsert {
	u.Set(mfa.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MFAUpsert) UpdateUpdatedAt() *MFAUpsert {
	u.SetExcluded(mfa.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MFAUpsert) SetDeletedAt(v time.Time) *MFAUpsert {
	u.Set(mfa.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MFAUpsert) UpdateDeletedAt() *MFAUpsert {
	u.SetExcluded(mfa.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MFAUpsert) ClearDeletedAt() *MFAUpsert {
	u.SetNull(mfa.FieldDeletedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *MFAUpsert) SetUserID(v xid.ID) *MFAUpsert {
	u.Set(mfa.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MFAUpsert) UpdateUserID() *MFAUpsert {
	u.SetExcluded(mfa.FieldUserID)
	return u
}

// SetMethod sets the "method" field.
func (u *MFAUpsert) SetMethod(v string) *MFAUpsert {
	u.Set(mfa.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *MFAUpsert) UpdateMethod() *MFAUpsert {
	u.SetExcluded(mfa.FieldMethod)
	return u
}

// SetSecret sets the "secret" field.
func (u *MFAUpsert) SetSecret(v string) *MFAUpsert {
	u.Set(mfa.FieldSecret, v)
	return u
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *MFAUpsert) UpdateSecret() *MFAUpsert {
	u.SetExcluded(mfa.FieldSecret)
	return u
}

// SetVerified sets the "verified" field.
func (u *MFAUpsert) SetVerified(v bool) *MFAUpsert {
	u.Set(mfa.FieldVerified, v)
	return u
}

// UpdateVerified sets the "verified" field to the value that was provided on create.
func (u *MFAUpsert) UpdateVerified() *MFAUpsert {
	u.SetExcluded(mfa.FieldVerified)
	return u
}

// SetActive sets the "active" field.
func (u *MFAUpsert) SetActive(v bool) *MFAUpsert {
	u.Set(mfa.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *MFAUpsert) UpdateActive() *MFAUpsert {
	u.SetExcluded(mfa.FieldActive)
	return u
}

// SetBackupCodes sets the "backup_codes" field.
func (u *MFAUpsert) SetBackupCodes(v []string) *MFAUpsert {
	u.Set(mfa.FieldBackupCodes, v)
	return u
}

// UpdateBackupCodes sets the "backup_codes" field to the value that was provided on create.
func (u *MFAUpsert) UpdateBackupCodes() *MFAUpsert {
	u.SetExcluded(mfa.FieldBackupCodes)
	return u
}

// ClearBackupCodes clears the value of the "backup_codes" field.
func (u *MFAUpsert) ClearBackupCodes() *MFAUpsert {
	u.SetNull(mfa.FieldBackupCodes)
	return u
}

// SetPhoneNumber sets the "phone_number" field.
func (u *MFAUpsert) SetPhoneNumber(v string) *MFAUpsert {
	u.Set(mfa.FieldPhoneNumber, v)
	return u
}

// UpdatePhoneNumber sets the "phone_number" field to the value that was provided on create.
func (u *MFAUpsert) UpdatePhoneNumber() *MFAUpsert {
	u.SetExcluded(mfa.FieldPhoneNumber)
	return u
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (u *MFAUpsert) ClearPhoneNumber() *MFAUpsert {
	u.SetNull(mfa.FieldPhoneNumber)
	return u
}

// SetEmail sets the "email" field.
func (u *MFAUpsert) SetEmail(v string) *MFAUpsert {
	u.Set(mfa.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *MFAUpsert) UpdateEmail() *MFAUpsert {
	u.SetExcluded(mfa.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *MFAUpsert) ClearEmail() *MFAUpsert {
	u.SetNull(mfa.FieldEmail)
	return u
}

// SetLastUsed sets the "last_used" field.
func (u *MFAUpsert) SetLastUsed(v time.Time) *MFAUpsert {
	u.Set(mfa.FieldLastUsed, v)
	return u
}

// UpdateLastUsed sets the "last_used" field to the value that was provided on create.
func (u *MFAUpsert) UpdateLastUsed() *MFAUpsert {
	u.SetExcluded(mfa.FieldLastUsed)
	return u
}

// ClearLastUsed clears the value of the "last_used" field.
func (u *MFAUpsert) ClearLastUsed() *MFAUpsert {
	u.SetNull(mfa.FieldLastUsed)
	return u
}

// SetMetadata sets the "metadata" field.
func (u *MFAUpsert) SetMetadata(v map[string]interface{}) *MFAUpsert {
	u.Set(mfa.FieldMetadata, v)
	return u
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *MFAUpsert) UpdateMetadata() *MFAUpsert {
	u.SetExcluded(mfa.FieldMetadata)
	return u
}

// ClearMetadata clears the value of the "metadata" field.
func (u *MFAUpsert) ClearMetadata() *MFAUpsert {
	u.SetNull(mfa.FieldMetadata)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MFA.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mfa.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MFAUpsertOne) UpdateNewValues() *MFAUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mfa.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(mfa.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MFA.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MFAUpsertOne) Ignore() *MFAUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MFAUpsertOne) DoNothing() *MFAUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MFACreate.OnConflict
// documentation for more info.
func (u *MFAUpsertOne) Update(set func(*MFAUpsert)) *MFAUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MFAUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MFAUpsertOne) SetUpdatedAt(v time.Time) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateUpdatedAt() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MFAUpsertOne) SetDeletedAt(v time.Time) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateDeletedAt() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MFAUpsertOne) ClearDeletedAt() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *MFAUpsertOne) SetUserID(v xid.ID) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateUserID() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateUserID()
	})
}

// SetMethod sets the "method" field.
func (u *MFAUpsertOne) SetMethod(v string) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateMethod() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateMethod()
	})
}

// SetSecret sets the "secret" field.
func (u *MFAUpsertOne) SetSecret(v string) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateSecret() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateSecret()
	})
}

// SetVerified sets the "verified" field.
func (u *MFAUpsertOne) SetVerified(v bool) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetVerified(v)
	})
}

// UpdateVerified sets the "verified" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateVerified() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateVerified()
	})
}

// SetActive sets the "active" field.
func (u *MFAUpsertOne) SetActive(v bool) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateActive() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateActive()
	})
}

// SetBackupCodes sets the "backup_codes" field.
func (u *MFAUpsertOne) SetBackupCodes(v []string) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetBackupCodes(v)
	})
}

// UpdateBackupCodes sets the "backup_codes" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateBackupCodes() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateBackupCodes()
	})
}

// ClearBackupCodes clears the value of the "backup_codes" field.
func (u *MFAUpsertOne) ClearBackupCodes() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.ClearBackupCodes()
	})
}

// SetPhoneNumber sets the "phone_number" field.
func (u *MFAUpsertOne) SetPhoneNumber(v string) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetPhoneNumber(v)
	})
}

// UpdatePhoneNumber sets the "phone_number" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdatePhoneNumber() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdatePhoneNumber()
	})
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (u *MFAUpsertOne) ClearPhoneNumber() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.ClearPhoneNumber()
	})
}

// SetEmail sets the "email" field.
func (u *MFAUpsertOne) SetEmail(v string) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateEmail() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *MFAUpsertOne) ClearEmail() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.ClearEmail()
	})
}

// SetLastUsed sets the "last_used" field.
func (u *MFAUpsertOne) SetLastUsed(v time.Time) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetLastUsed(v)
	})
}

// UpdateLastUsed sets the "last_used" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateLastUsed() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateLastUsed()
	})
}

// ClearLastUsed clears the value of the "last_used" field.
func (u *MFAUpsertOne) ClearLastUsed() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.ClearLastUsed()
	})
}

// SetMetadata sets the "metadata" field.
func (u *MFAUpsertOne) SetMetadata(v map[string]interface{}) *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *MFAUpsertOne) UpdateMetadata() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateMetadata()
	})
}

// ClearMetadata clears the value of the "metadata" field.
func (u *MFAUpsertOne) ClearMetadata() *MFAUpsertOne {
	return u.Update(func(s *MFAUpsert) {
		s.ClearMetadata()
	})
}

// Exec executes the query.
func (u *MFAUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MFACreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MFAUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MFAUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MFAUpsertOne.ID is not supported by MySQL driver. Use MFAUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MFAUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MFACreateBulk is the builder for creating many MFA entities in bulk.
type MFACreateBulk struct {
	config
	err      error
	builders []*MFACreate
	conflict []sql.ConflictOption
}

// Save creates the MFA entities in the database.
func (mcb *MFACreateBulk) Save(ctx context.Context) ([]*MFA, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*MFA, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MFAMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MFACreateBulk) SaveX(ctx context.Context) []*MFA {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MFACreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MFACreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MFA.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MFAUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcb *MFACreateBulk) OnConflict(opts ...sql.ConflictOption) *MFAUpsertBulk {
	mcb.conflict = opts
	return &MFAUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MFA.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MFACreateBulk) OnConflictColumns(columns ...string) *MFAUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MFAUpsertBulk{
		create: mcb,
	}
}

// MFAUpsertBulk is the builder for "upsert"-ing
// a bulk of MFA nodes.
type MFAUpsertBulk struct {
	create *MFACreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MFA.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mfa.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MFAUpsertBulk) UpdateNewValues() *MFAUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mfa.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(mfa.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MFA.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MFAUpsertBulk) Ignore() *MFAUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MFAUpsertBulk) DoNothing() *MFAUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MFACreateBulk.OnConflict
// documentation for more info.
func (u *MFAUpsertBulk) Update(set func(*MFAUpsert)) *MFAUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MFAUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MFAUpsertBulk) SetUpdatedAt(v time.Time) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateUpdatedAt() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MFAUpsertBulk) SetDeletedAt(v time.Time) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateDeletedAt() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MFAUpsertBulk) ClearDeletedAt() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *MFAUpsertBulk) SetUserID(v xid.ID) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateUserID() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateUserID()
	})
}

// SetMethod sets the "method" field.
func (u *MFAUpsertBulk) SetMethod(v string) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateMethod() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateMethod()
	})
}

// SetSecret sets the "secret" field.
func (u *MFAUpsertBulk) SetSecret(v string) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateSecret() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateSecret()
	})
}

// SetVerified sets the "verified" field.
func (u *MFAUpsertBulk) SetVerified(v bool) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetVerified(v)
	})
}

// UpdateVerified sets the "verified" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateVerified() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateVerified()
	})
}

// SetActive sets the "active" field.
func (u *MFAUpsertBulk) SetActive(v bool) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateActive() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateActive()
	})
}

// SetBackupCodes sets the "backup_codes" field.
func (u *MFAUpsertBulk) SetBackupCodes(v []string) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetBackupCodes(v)
	})
}

// UpdateBackupCodes sets the "backup_codes" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateBackupCodes() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateBackupCodes()
	})
}

// ClearBackupCodes clears the value of the "backup_codes" field.
func (u *MFAUpsertBulk) ClearBackupCodes() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.ClearBackupCodes()
	})
}

// SetPhoneNumber sets the "phone_number" field.
func (u *MFAUpsertBulk) SetPhoneNumber(v string) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetPhoneNumber(v)
	})
}

// UpdatePhoneNumber sets the "phone_number" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdatePhoneNumber() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdatePhoneNumber()
	})
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (u *MFAUpsertBulk) ClearPhoneNumber() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.ClearPhoneNumber()
	})
}

// SetEmail sets the "email" field.
func (u *MFAUpsertBulk) SetEmail(v string) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateEmail() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *MFAUpsertBulk) ClearEmail() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.ClearEmail()
	})
}

// SetLastUsed sets the "last_used" field.
func (u *MFAUpsertBulk) SetLastUsed(v time.Time) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetLastUsed(v)
	})
}

// UpdateLastUsed sets the "last_used" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateLastUsed() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateLastUsed()
	})
}

// ClearLastUsed clears the value of the "last_used" field.
func (u *MFAUpsertBulk) ClearLastUsed() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.ClearLastUsed()
	})
}

// SetMetadata sets the "metadata" field.
func (u *MFAUpsertBulk) SetMetadata(v map[string]interface{}) *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *MFAUpsertBulk) UpdateMetadata() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.UpdateMetadata()
	})
}

// ClearMetadata clears the value of the "metadata" field.
func (u *MFAUpsertBulk) ClearMetadata() *MFAUpsertBulk {
	return u.Update(func(s *MFAUpsert) {
		s.ClearMetadata()
	})
}

// Exec executes the query.
func (u *MFAUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MFACreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MFACreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MFAUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
