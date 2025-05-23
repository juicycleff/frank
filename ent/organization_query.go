// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/juicycleff/frank/ent/apikey"
	"github.com/juicycleff/frank/ent/identityprovider"
	"github.com/juicycleff/frank/ent/oauthclient"
	"github.com/juicycleff/frank/ent/organization"
	"github.com/juicycleff/frank/ent/organizationfeature"
	"github.com/juicycleff/frank/ent/predicate"
	"github.com/juicycleff/frank/ent/user"
	"github.com/juicycleff/frank/ent/webhook"
)

// OrganizationQuery is the builder for querying Organization entities.
type OrganizationQuery struct {
	config
	ctx                   *QueryContext
	order                 []organization.OrderOption
	inters                []Interceptor
	predicates            []predicate.Organization
	withUsers             *UserQuery
	withAPIKeys           *ApiKeyQuery
	withWebhooks          *WebhookQuery
	withFeatureFlags      *OrganizationFeatureQuery
	withIdentityProviders *IdentityProviderQuery
	withOauthClients      *OAuthClientQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationQuery builder.
func (oq *OrganizationQuery) Where(ps ...predicate.Organization) *OrganizationQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OrganizationQuery) Limit(limit int) *OrganizationQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OrganizationQuery) Offset(offset int) *OrganizationQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OrganizationQuery) Unique(unique bool) *OrganizationQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OrganizationQuery) Order(o ...organization.OrderOption) *OrganizationQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryUsers chains the current query on the "users" edge.
func (oq *OrganizationQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, organization.UsersTable, organization.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAPIKeys chains the current query on the "api_keys" edge.
func (oq *OrganizationQuery) QueryAPIKeys() *ApiKeyQuery {
	query := (&ApiKeyClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(apikey.Table, apikey.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.APIKeysTable, organization.APIKeysColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWebhooks chains the current query on the "webhooks" edge.
func (oq *OrganizationQuery) QueryWebhooks() *WebhookQuery {
	query := (&WebhookClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(webhook.Table, webhook.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.WebhooksTable, organization.WebhooksColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeatureFlags chains the current query on the "feature_flags" edge.
func (oq *OrganizationQuery) QueryFeatureFlags() *OrganizationFeatureQuery {
	query := (&OrganizationFeatureClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(organizationfeature.Table, organizationfeature.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.FeatureFlagsTable, organization.FeatureFlagsColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIdentityProviders chains the current query on the "identity_providers" edge.
func (oq *OrganizationQuery) QueryIdentityProviders() *IdentityProviderQuery {
	query := (&IdentityProviderClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(identityprovider.Table, identityprovider.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.IdentityProvidersTable, organization.IdentityProvidersColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOauthClients chains the current query on the "oauth_clients" edge.
func (oq *OrganizationQuery) QueryOauthClients() *OAuthClientQuery {
	query := (&OAuthClientClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(oauthclient.Table, oauthclient.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.OauthClientsTable, organization.OauthClientsColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Organization entity from the query.
// Returns a *NotFoundError when no Organization was found.
func (oq *OrganizationQuery) First(ctx context.Context) (*Organization, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organization.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OrganizationQuery) FirstX(ctx context.Context) *Organization {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Organization ID from the query.
// Returns a *NotFoundError when no Organization ID was found.
func (oq *OrganizationQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organization.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OrganizationQuery) FirstIDX(ctx context.Context) string {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Organization entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Organization entity is found.
// Returns a *NotFoundError when no Organization entities are found.
func (oq *OrganizationQuery) Only(ctx context.Context) (*Organization, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organization.Label}
	default:
		return nil, &NotSingularError{organization.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OrganizationQuery) OnlyX(ctx context.Context) *Organization {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Organization ID in the query.
// Returns a *NotSingularError when more than one Organization ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OrganizationQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organization.Label}
	default:
		err = &NotSingularError{organization.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OrganizationQuery) OnlyIDX(ctx context.Context) string {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Organizations.
func (oq *OrganizationQuery) All(ctx context.Context) ([]*Organization, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryAll)
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Organization, *OrganizationQuery]()
	return withInterceptors[[]*Organization](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OrganizationQuery) AllX(ctx context.Context) []*Organization {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Organization IDs.
func (oq *OrganizationQuery) IDs(ctx context.Context) (ids []string, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryIDs)
	if err = oq.Select(organization.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OrganizationQuery) IDsX(ctx context.Context) []string {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OrganizationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryCount)
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OrganizationQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OrganizationQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OrganizationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryExist)
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OrganizationQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OrganizationQuery) Clone() *OrganizationQuery {
	if oq == nil {
		return nil
	}
	return &OrganizationQuery{
		config:                oq.config,
		ctx:                   oq.ctx.Clone(),
		order:                 append([]organization.OrderOption{}, oq.order...),
		inters:                append([]Interceptor{}, oq.inters...),
		predicates:            append([]predicate.Organization{}, oq.predicates...),
		withUsers:             oq.withUsers.Clone(),
		withAPIKeys:           oq.withAPIKeys.Clone(),
		withWebhooks:          oq.withWebhooks.Clone(),
		withFeatureFlags:      oq.withFeatureFlags.Clone(),
		withIdentityProviders: oq.withIdentityProviders.Clone(),
		withOauthClients:      oq.withOauthClients.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithUsers(opts ...func(*UserQuery)) *OrganizationQuery {
	query := (&UserClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withUsers = query
	return oq
}

// WithAPIKeys tells the query-builder to eager-load the nodes that are connected to
// the "api_keys" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithAPIKeys(opts ...func(*ApiKeyQuery)) *OrganizationQuery {
	query := (&ApiKeyClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withAPIKeys = query
	return oq
}

// WithWebhooks tells the query-builder to eager-load the nodes that are connected to
// the "webhooks" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithWebhooks(opts ...func(*WebhookQuery)) *OrganizationQuery {
	query := (&WebhookClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withWebhooks = query
	return oq
}

// WithFeatureFlags tells the query-builder to eager-load the nodes that are connected to
// the "feature_flags" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithFeatureFlags(opts ...func(*OrganizationFeatureQuery)) *OrganizationQuery {
	query := (&OrganizationFeatureClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withFeatureFlags = query
	return oq
}

// WithIdentityProviders tells the query-builder to eager-load the nodes that are connected to
// the "identity_providers" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithIdentityProviders(opts ...func(*IdentityProviderQuery)) *OrganizationQuery {
	query := (&IdentityProviderClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withIdentityProviders = query
	return oq
}

// WithOauthClients tells the query-builder to eager-load the nodes that are connected to
// the "oauth_clients" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithOauthClients(opts ...func(*OAuthClientQuery)) *OrganizationQuery {
	query := (&OAuthClientClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withOauthClients = query
	return oq
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
//	client.Organization.Query().
//		GroupBy(organization.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OrganizationQuery) GroupBy(field string, fields ...string) *OrganizationGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = organization.Label
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
//	client.Organization.Query().
//		Select(organization.FieldCreatedAt).
//		Scan(ctx, &v)
func (oq *OrganizationQuery) Select(fields ...string) *OrganizationSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OrganizationSelect{OrganizationQuery: oq}
	sbuild.label = organization.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationSelect configured with the given aggregations.
func (oq *OrganizationQuery) Aggregate(fns ...AggregateFunc) *OrganizationSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OrganizationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !organization.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OrganizationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Organization, error) {
	var (
		nodes       = []*Organization{}
		_spec       = oq.querySpec()
		loadedTypes = [6]bool{
			oq.withUsers != nil,
			oq.withAPIKeys != nil,
			oq.withWebhooks != nil,
			oq.withFeatureFlags != nil,
			oq.withIdentityProviders != nil,
			oq.withOauthClients != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Organization).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Organization{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withUsers; query != nil {
		if err := oq.loadUsers(ctx, query, nodes,
			func(n *Organization) { n.Edges.Users = []*User{} },
			func(n *Organization, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withAPIKeys; query != nil {
		if err := oq.loadAPIKeys(ctx, query, nodes,
			func(n *Organization) { n.Edges.APIKeys = []*ApiKey{} },
			func(n *Organization, e *ApiKey) { n.Edges.APIKeys = append(n.Edges.APIKeys, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withWebhooks; query != nil {
		if err := oq.loadWebhooks(ctx, query, nodes,
			func(n *Organization) { n.Edges.Webhooks = []*Webhook{} },
			func(n *Organization, e *Webhook) { n.Edges.Webhooks = append(n.Edges.Webhooks, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withFeatureFlags; query != nil {
		if err := oq.loadFeatureFlags(ctx, query, nodes,
			func(n *Organization) { n.Edges.FeatureFlags = []*OrganizationFeature{} },
			func(n *Organization, e *OrganizationFeature) { n.Edges.FeatureFlags = append(n.Edges.FeatureFlags, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withIdentityProviders; query != nil {
		if err := oq.loadIdentityProviders(ctx, query, nodes,
			func(n *Organization) { n.Edges.IdentityProviders = []*IdentityProvider{} },
			func(n *Organization, e *IdentityProvider) {
				n.Edges.IdentityProviders = append(n.Edges.IdentityProviders, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := oq.withOauthClients; query != nil {
		if err := oq.loadOauthClients(ctx, query, nodes,
			func(n *Organization) { n.Edges.OauthClients = []*OAuthClient{} },
			func(n *Organization, e *OAuthClient) { n.Edges.OauthClients = append(n.Edges.OauthClients, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OrganizationQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Organization)
	nids := make(map[string]map[*Organization]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(organization.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(organization.UsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(organization.UsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(organization.UsersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Organization]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (oq *OrganizationQuery) loadAPIKeys(ctx context.Context, query *ApiKeyQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *ApiKey)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Organization)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(apikey.FieldOrganizationID)
	}
	query.Where(predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organization.APIKeysColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrganizationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrganizationQuery) loadWebhooks(ctx context.Context, query *WebhookQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *Webhook)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Organization)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(webhook.FieldOrganizationID)
	}
	query.Where(predicate.Webhook(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organization.WebhooksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrganizationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrganizationQuery) loadFeatureFlags(ctx context.Context, query *OrganizationFeatureQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *OrganizationFeature)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Organization)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(organizationfeature.FieldOrganizationID)
	}
	query.Where(predicate.OrganizationFeature(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organization.FeatureFlagsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrganizationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrganizationQuery) loadIdentityProviders(ctx context.Context, query *IdentityProviderQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *IdentityProvider)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Organization)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(identityprovider.FieldOrganizationID)
	}
	query.Where(predicate.IdentityProvider(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organization.IdentityProvidersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrganizationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrganizationQuery) loadOauthClients(ctx context.Context, query *OAuthClientQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *OAuthClient)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Organization)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(oauthclient.FieldOrganizationID)
	}
	query.Where(predicate.OAuthClient(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organization.OauthClientsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrganizationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (oq *OrganizationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OrganizationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for i := range fields {
			if fields[i] != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OrganizationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(organization.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = organization.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrganizationGroupBy is the group-by builder for Organization entities.
type OrganizationGroupBy struct {
	selector
	build *OrganizationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OrganizationGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OrganizationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, ent.OpQueryGroupBy)
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationQuery, *OrganizationGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OrganizationGroupBy) sqlScan(ctx context.Context, root *OrganizationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationSelect is the builder for selecting fields of Organization entities.
type OrganizationSelect struct {
	*OrganizationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OrganizationSelect) Aggregate(fns ...AggregateFunc) *OrganizationSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OrganizationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, ent.OpQuerySelect)
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationQuery, *OrganizationSelect](ctx, os.OrganizationQuery, os, os.inters, v)
}

func (os *OrganizationSelect) sqlScan(ctx context.Context, root *OrganizationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
