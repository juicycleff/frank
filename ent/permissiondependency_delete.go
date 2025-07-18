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
	"github.com/xraph/frank/ent/permissiondependency"
	"github.com/xraph/frank/ent/predicate"
)

// PermissionDependencyDelete is the builder for deleting a PermissionDependency entity.
type PermissionDependencyDelete struct {
	config
	hooks    []Hook
	mutation *PermissionDependencyMutation
}

// Where appends a list predicates to the PermissionDependencyDelete builder.
func (pdd *PermissionDependencyDelete) Where(ps ...predicate.PermissionDependency) *PermissionDependencyDelete {
	pdd.mutation.Where(ps...)
	return pdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pdd *PermissionDependencyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pdd.sqlExec, pdd.mutation, pdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pdd *PermissionDependencyDelete) ExecX(ctx context.Context) int {
	n, err := pdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pdd *PermissionDependencyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(permissiondependency.Table, sqlgraph.NewFieldSpec(permissiondependency.FieldID, field.TypeString))
	if ps := pdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pdd.mutation.done = true
	return affected, err
}

// PermissionDependencyDeleteOne is the builder for deleting a single PermissionDependency entity.
type PermissionDependencyDeleteOne struct {
	pdd *PermissionDependencyDelete
}

// Where appends a list predicates to the PermissionDependencyDelete builder.
func (pddo *PermissionDependencyDeleteOne) Where(ps ...predicate.PermissionDependency) *PermissionDependencyDeleteOne {
	pddo.pdd.mutation.Where(ps...)
	return pddo
}

// Exec executes the deletion query.
func (pddo *PermissionDependencyDeleteOne) Exec(ctx context.Context) error {
	n, err := pddo.pdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{permissiondependency.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pddo *PermissionDependencyDeleteOne) ExecX(ctx context.Context) {
	if err := pddo.Exec(ctx); err != nil {
		panic(err)
	}
}
