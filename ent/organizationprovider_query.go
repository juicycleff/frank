// Copyright 2023-present XRaph LLC. All rights reserved.
// This source code is licensed under the XRaph LLC license found
// in the LICENSE file in the root directory of this source tree.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xraph/frank/ent/identityprovider"
	"github.com/xraph/frank/ent/organization"
	"github.com/xraph/frank/ent/organizationprovider"
	"github.com/xraph/frank/ent/predicate"
	"github.com/xraph/frank/ent/providertemplate"
	"github.com/rs/xid"
)

// OrganizationProviderQuery is the builder for querying OrganizationProvider entities.
type OrganizationProviderQuery struct {
	config
	ctx              *QueryContext
	order            []organizationprovider.OrderOption
	inters           []Interceptor
	predicates       []predicate.OrganizationProvider
	withOrganization *OrganizationQuery
	withProvider     *IdentityProviderQuery
	withTemplate     *ProviderTemplateQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationProviderQuery builder.
func (opq *OrganizationProviderQuery) Where(ps ...predicate.OrganizationProvider) *OrganizationProviderQuery {
	opq.predicates = append(opq.predicates, ps...)
	return opq
}

// Limit the number of records to be returned by this query.
func (opq *OrganizationProviderQuery) Limit(limit int) *OrganizationProviderQuery {
	opq.ctx.Limit = &limit
	return opq
}

// Offset to start from.
func (opq *OrganizationProviderQuery) Offset(offset int) *OrganizationProviderQuery {
	opq.ctx.Offset = &offset
	return opq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (opq *OrganizationProviderQuery) Unique(unique bool) *OrganizationProviderQuery {
	opq.ctx.Unique = &unique
	return opq
}

// Order specifies how the records should be ordered.
func (opq *OrganizationProviderQuery) Order(o ...organizationprovider.OrderOption) *OrganizationProviderQuery {
	opq.order = append(opq.order, o...)
	return opq
}

// QueryOrganization chains the current query on the "organization" edge.
func (opq *OrganizationProviderQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: opq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationprovider.Table, organizationprovider.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, organizationprovider.OrganizationTable, organizationprovider.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProvider chains the current query on the "provider" edge.
func (opq *OrganizationProviderQuery) QueryProvider() *IdentityProviderQuery {
	query := (&IdentityProviderClient{config: opq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationprovider.Table, organizationprovider.FieldID, selector),
			sqlgraph.To(identityprovider.Table, identityprovider.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, organizationprovider.ProviderTable, organizationprovider.ProviderColumn),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTemplate chains the current query on the "template" edge.
func (opq *OrganizationProviderQuery) QueryTemplate() *ProviderTemplateQuery {
	query := (&ProviderTemplateClient{config: opq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationprovider.Table, organizationprovider.FieldID, selector),
			sqlgraph.To(providertemplate.Table, providertemplate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, organizationprovider.TemplateTable, organizationprovider.TemplateColumn),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrganizationProvider entity from the query.
// Returns a *NotFoundError when no OrganizationProvider was found.
func (opq *OrganizationProviderQuery) First(ctx context.Context) (*OrganizationProvider, error) {
	nodes, err := opq.Limit(1).All(setContextOp(ctx, opq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationprovider.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (opq *OrganizationProviderQuery) FirstX(ctx context.Context) *OrganizationProvider {
	node, err := opq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationProvider ID from the query.
// Returns a *NotFoundError when no OrganizationProvider ID was found.
func (opq *OrganizationProviderQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = opq.Limit(1).IDs(setContextOp(ctx, opq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationprovider.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (opq *OrganizationProviderQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := opq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationProvider entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationProvider entity is found.
// Returns a *NotFoundError when no OrganizationProvider entities are found.
func (opq *OrganizationProviderQuery) Only(ctx context.Context) (*OrganizationProvider, error) {
	nodes, err := opq.Limit(2).All(setContextOp(ctx, opq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationprovider.Label}
	default:
		return nil, &NotSingularError{organizationprovider.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (opq *OrganizationProviderQuery) OnlyX(ctx context.Context) *OrganizationProvider {
	node, err := opq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationProvider ID in the query.
// Returns a *NotSingularError when more than one OrganizationProvider ID is found.
// Returns a *NotFoundError when no entities are found.
func (opq *OrganizationProviderQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = opq.Limit(2).IDs(setContextOp(ctx, opq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationprovider.Label}
	default:
		err = &NotSingularError{organizationprovider.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (opq *OrganizationProviderQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := opq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationProviders.
func (opq *OrganizationProviderQuery) All(ctx context.Context) ([]*OrganizationProvider, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryAll)
	if err := opq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrganizationProvider, *OrganizationProviderQuery]()
	return withInterceptors[[]*OrganizationProvider](ctx, opq, qr, opq.inters)
}

// AllX is like All, but panics if an error occurs.
func (opq *OrganizationProviderQuery) AllX(ctx context.Context) []*OrganizationProvider {
	nodes, err := opq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationProvider IDs.
func (opq *OrganizationProviderQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if opq.ctx.Unique == nil && opq.path != nil {
		opq.Unique(true)
	}
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryIDs)
	if err = opq.Select(organizationprovider.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (opq *OrganizationProviderQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := opq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (opq *OrganizationProviderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryCount)
	if err := opq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, opq, querierCount[*OrganizationProviderQuery](), opq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (opq *OrganizationProviderQuery) CountX(ctx context.Context) int {
	count, err := opq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (opq *OrganizationProviderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryExist)
	switch _, err := opq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (opq *OrganizationProviderQuery) ExistX(ctx context.Context) bool {
	exist, err := opq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationProviderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (opq *OrganizationProviderQuery) Clone() *OrganizationProviderQuery {
	if opq == nil {
		return nil
	}
	return &OrganizationProviderQuery{
		config:           opq.config,
		ctx:              opq.ctx.Clone(),
		order:            append([]organizationprovider.OrderOption{}, opq.order...),
		inters:           append([]Interceptor{}, opq.inters...),
		predicates:       append([]predicate.OrganizationProvider{}, opq.predicates...),
		withOrganization: opq.withOrganization.Clone(),
		withProvider:     opq.withProvider.Clone(),
		withTemplate:     opq.withTemplate.Clone(),
		// clone intermediate query.
		sql:       opq.sql.Clone(),
		path:      opq.path,
		modifiers: append([]func(*sql.Selector){}, opq.modifiers...),
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrganizationProviderQuery) WithOrganization(opts ...func(*OrganizationQuery)) *OrganizationProviderQuery {
	query := (&OrganizationClient{config: opq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	opq.withOrganization = query
	return opq
}

// WithProvider tells the query-builder to eager-load the nodes that are connected to
// the "provider" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrganizationProviderQuery) WithProvider(opts ...func(*IdentityProviderQuery)) *OrganizationProviderQuery {
	query := (&IdentityProviderClient{config: opq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	opq.withProvider = query
	return opq
}

// WithTemplate tells the query-builder to eager-load the nodes that are connected to
// the "template" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrganizationProviderQuery) WithTemplate(opts ...func(*ProviderTemplateQuery)) *OrganizationProviderQuery {
	query := (&ProviderTemplateClient{config: opq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	opq.withTemplate = query
	return opq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrganizationProvider.Query().
//		GroupBy(organizationprovider.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (opq *OrganizationProviderQuery) GroupBy(field string, fields ...string) *OrganizationProviderGroupBy {
	opq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationProviderGroupBy{build: opq}
	grbuild.flds = &opq.ctx.Fields
	grbuild.label = organizationprovider.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.OrganizationProvider.Query().
//		Select(organizationprovider.FieldCreatedAt).
//		Scan(ctx, &v)
func (opq *OrganizationProviderQuery) Select(fields ...string) *OrganizationProviderSelect {
	opq.ctx.Fields = append(opq.ctx.Fields, fields...)
	sbuild := &OrganizationProviderSelect{OrganizationProviderQuery: opq}
	sbuild.label = organizationprovider.Label
	sbuild.flds, sbuild.scan = &opq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationProviderSelect configured with the given aggregations.
func (opq *OrganizationProviderQuery) Aggregate(fns ...AggregateFunc) *OrganizationProviderSelect {
	return opq.Select().Aggregate(fns...)
}

func (opq *OrganizationProviderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range opq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, opq); err != nil {
				return err
			}
		}
	}
	for _, f := range opq.ctx.Fields {
		if !organizationprovider.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if opq.path != nil {
		prev, err := opq.path(ctx)
		if err != nil {
			return err
		}
		opq.sql = prev
	}
	return nil
}

func (opq *OrganizationProviderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationProvider, error) {
	var (
		nodes       = []*OrganizationProvider{}
		_spec       = opq.querySpec()
		loadedTypes = [3]bool{
			opq.withOrganization != nil,
			opq.withProvider != nil,
			opq.withTemplate != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrganizationProvider).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrganizationProvider{config: opq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(opq.modifiers) > 0 {
		_spec.Modifiers = opq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, opq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := opq.withOrganization; query != nil {
		if err := opq.loadOrganization(ctx, query, nodes, nil,
			func(n *OrganizationProvider, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := opq.withProvider; query != nil {
		if err := opq.loadProvider(ctx, query, nodes, nil,
			func(n *OrganizationProvider, e *IdentityProvider) { n.Edges.Provider = e }); err != nil {
			return nil, err
		}
	}
	if query := opq.withTemplate; query != nil {
		if err := opq.loadTemplate(ctx, query, nodes, nil,
			func(n *OrganizationProvider, e *ProviderTemplate) { n.Edges.Template = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (opq *OrganizationProviderQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*OrganizationProvider, init func(*OrganizationProvider), assign func(*OrganizationProvider, *Organization)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*OrganizationProvider)
	for i := range nodes {
		fk := nodes[i].OrganizationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (opq *OrganizationProviderQuery) loadProvider(ctx context.Context, query *IdentityProviderQuery, nodes []*OrganizationProvider, init func(*OrganizationProvider), assign func(*OrganizationProvider, *IdentityProvider)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*OrganizationProvider)
	for i := range nodes {
		fk := nodes[i].ProviderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(identityprovider.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "provider_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (opq *OrganizationProviderQuery) loadTemplate(ctx context.Context, query *ProviderTemplateQuery, nodes []*OrganizationProvider, init func(*OrganizationProvider), assign func(*OrganizationProvider, *ProviderTemplate)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*OrganizationProvider)
	for i := range nodes {
		fk := nodes[i].TemplateID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(providertemplate.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "template_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (opq *OrganizationProviderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := opq.querySpec()
	if len(opq.modifiers) > 0 {
		_spec.Modifiers = opq.modifiers
	}
	_spec.Node.Columns = opq.ctx.Fields
	if len(opq.ctx.Fields) > 0 {
		_spec.Unique = opq.ctx.Unique != nil && *opq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, opq.driver, _spec)
}

func (opq *OrganizationProviderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organizationprovider.Table, organizationprovider.Columns, sqlgraph.NewFieldSpec(organizationprovider.FieldID, field.TypeString))
	_spec.From = opq.sql
	if unique := opq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if opq.path != nil {
		_spec.Unique = true
	}
	if fields := opq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationprovider.FieldID)
		for i := range fields {
			if fields[i] != organizationprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if opq.withOrganization != nil {
			_spec.Node.AddColumnOnce(organizationprovider.FieldOrganizationID)
		}
		if opq.withProvider != nil {
			_spec.Node.AddColumnOnce(organizationprovider.FieldProviderID)
		}
		if opq.withTemplate != nil {
			_spec.Node.AddColumnOnce(organizationprovider.FieldTemplateID)
		}
	}
	if ps := opq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := opq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := opq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := opq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (opq *OrganizationProviderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(opq.driver.Dialect())
	t1 := builder.Table(organizationprovider.Table)
	columns := opq.ctx.Fields
	if len(columns) == 0 {
		columns = organizationprovider.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if opq.sql != nil {
		selector = opq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if opq.ctx.Unique != nil && *opq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range opq.modifiers {
		m(selector)
	}
	for _, p := range opq.predicates {
		p(selector)
	}
	for _, p := range opq.order {
		p(selector)
	}
	if offset := opq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := opq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (opq *OrganizationProviderQuery) ForUpdate(opts ...sql.LockOption) *OrganizationProviderQuery {
	if opq.driver.Dialect() == dialect.Postgres {
		opq.Unique(false)
	}
	opq.modifiers = append(opq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return opq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (opq *OrganizationProviderQuery) ForShare(opts ...sql.LockOption) *OrganizationProviderQuery {
	if opq.driver.Dialect() == dialect.Postgres {
		opq.Unique(false)
	}
	opq.modifiers = append(opq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return opq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (opq *OrganizationProviderQuery) Modify(modifiers ...func(s *sql.Selector)) *OrganizationProviderSelect {
	opq.modifiers = append(opq.modifiers, modifiers...)
	return opq.Select()
}

// OrganizationProviderGroupBy is the group-by builder for OrganizationProvider entities.
type OrganizationProviderGroupBy struct {
	selector
	build *OrganizationProviderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (opgb *OrganizationProviderGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationProviderGroupBy {
	opgb.fns = append(opgb.fns, fns...)
	return opgb
}

// Scan applies the selector query and scans the result into the given value.
func (opgb *OrganizationProviderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, opgb.build.ctx, ent.OpQueryGroupBy)
	if err := opgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationProviderQuery, *OrganizationProviderGroupBy](ctx, opgb.build, opgb, opgb.build.inters, v)
}

func (opgb *OrganizationProviderGroupBy) sqlScan(ctx context.Context, root *OrganizationProviderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(opgb.fns))
	for _, fn := range opgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*opgb.flds)+len(opgb.fns))
		for _, f := range *opgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*opgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := opgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationProviderSelect is the builder for selecting fields of OrganizationProvider entities.
type OrganizationProviderSelect struct {
	*OrganizationProviderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ops *OrganizationProviderSelect) Aggregate(fns ...AggregateFunc) *OrganizationProviderSelect {
	ops.fns = append(ops.fns, fns...)
	return ops
}

// Scan applies the selector query and scans the result into the given value.
func (ops *OrganizationProviderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ops.ctx, ent.OpQuerySelect)
	if err := ops.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationProviderQuery, *OrganizationProviderSelect](ctx, ops.OrganizationProviderQuery, ops, ops.inters, v)
}

func (ops *OrganizationProviderSelect) sqlScan(ctx context.Context, root *OrganizationProviderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ops.fns))
	for _, fn := range ops.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ops.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ops.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ops *OrganizationProviderSelect) Modify(modifiers ...func(s *sql.Selector)) *OrganizationProviderSelect {
	ops.modifiers = append(ops.modifiers, modifiers...)
	return ops
}
