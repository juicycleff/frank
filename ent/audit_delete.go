// Copyright 2023-present XRaph LLC. All rights reserved.
// This source code is licensed under the XRaph LLC license found
// in the LICENSE file in the root directory of this source tree.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xraph/frank/ent/audit"
	"github.com/xraph/frank/ent/predicate"
)

// AuditDelete is the builder for deleting a Audit entity.
type AuditDelete struct {
	config
	hooks    []Hook
	mutation *AuditMutation
}

// Where appends a list predicates to the AuditDelete builder.
func (ad *AuditDelete) Where(ps ...predicate.Audit) *AuditDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AuditDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AuditDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AuditDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(audit.Table, sqlgraph.NewFieldSpec(audit.FieldID, field.TypeString))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AuditDeleteOne is the builder for deleting a single Audit entity.
type AuditDeleteOne struct {
	ad *AuditDelete
}

// Where appends a list predicates to the AuditDelete builder.
func (ado *AuditDeleteOne) Where(ps ...predicate.Audit) *AuditDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AuditDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{audit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AuditDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
