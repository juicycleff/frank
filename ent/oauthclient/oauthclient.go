// Copyright 2023-present XRaph LLC. All rights reserved.
// This source code is licensed under the XRaph LLC license found
// in the LICENSE file in the root directory of this source tree.
// Code generated by ent, DO NOT EDIT.

package oauthclient

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the oauthclient type in the database.
	Label = "oauth_client"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldClientSecret holds the string denoting the client_secret field in the database.
	FieldClientSecret = "client_secret"
	// FieldClientName holds the string denoting the client_name field in the database.
	FieldClientName = "client_name"
	// FieldClientDescription holds the string denoting the client_description field in the database.
	FieldClientDescription = "client_description"
	// FieldClientURI holds the string denoting the client_uri field in the database.
	FieldClientURI = "client_uri"
	// FieldLogoURI holds the string denoting the logo_uri field in the database.
	FieldLogoURI = "logo_uri"
	// FieldRedirectUris holds the string denoting the redirect_uris field in the database.
	FieldRedirectUris = "redirect_uris"
	// FieldPostLogoutRedirectUris holds the string denoting the post_logout_redirect_uris field in the database.
	FieldPostLogoutRedirectUris = "post_logout_redirect_uris"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldAllowedCorsOrigins holds the string denoting the allowed_cors_origins field in the database.
	FieldAllowedCorsOrigins = "allowed_cors_origins"
	// FieldAllowedGrantTypes holds the string denoting the allowed_grant_types field in the database.
	FieldAllowedGrantTypes = "allowed_grant_types"
	// FieldTokenExpirySeconds holds the string denoting the token_expiry_seconds field in the database.
	FieldTokenExpirySeconds = "token_expiry_seconds"
	// FieldRefreshTokenExpirySeconds holds the string denoting the refresh_token_expiry_seconds field in the database.
	FieldRefreshTokenExpirySeconds = "refresh_token_expiry_seconds"
	// FieldAuthCodeExpirySeconds holds the string denoting the auth_code_expiry_seconds field in the database.
	FieldAuthCodeExpirySeconds = "auth_code_expiry_seconds"
	// FieldRequiresPkce holds the string denoting the requires_pkce field in the database.
	FieldRequiresPkce = "requires_pkce"
	// FieldRequiresConsent holds the string denoting the requires_consent field in the database.
	FieldRequiresConsent = "requires_consent"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeTokens holds the string denoting the tokens edge name in mutations.
	EdgeTokens = "tokens"
	// EdgeAuthorizations holds the string denoting the authorizations edge name in mutations.
	EdgeAuthorizations = "authorizations"
	// EdgeScopes holds the string denoting the scopes edge name in mutations.
	EdgeScopes = "scopes"
	// Table holds the table name of the oauthclient in the database.
	Table = "oauth_clients"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "oauth_clients"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// TokensTable is the table that holds the tokens relation/edge.
	TokensTable = "oauth_tokens"
	// TokensInverseTable is the table name for the OAuthToken entity.
	// It exists in this package in order to avoid circular dependency with the "oauthtoken" package.
	TokensInverseTable = "oauth_tokens"
	// TokensColumn is the table column denoting the tokens relation/edge.
	TokensColumn = "client_id"
	// AuthorizationsTable is the table that holds the authorizations relation/edge.
	AuthorizationsTable = "oauth_authorizations"
	// AuthorizationsInverseTable is the table name for the OAuthAuthorization entity.
	// It exists in this package in order to avoid circular dependency with the "oauthauthorization" package.
	AuthorizationsInverseTable = "oauth_authorizations"
	// AuthorizationsColumn is the table column denoting the authorizations relation/edge.
	AuthorizationsColumn = "client_id"
	// ScopesTable is the table that holds the scopes relation/edge. The primary key declared below.
	ScopesTable = "oauth_client_scopes"
	// ScopesInverseTable is the table name for the OAuthScope entity.
	// It exists in this package in order to avoid circular dependency with the "oauthscope" package.
	ScopesInverseTable = "oauth_scopes"
)

// Columns holds all SQL columns for oauthclient fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldClientID,
	FieldClientSecret,
	FieldClientName,
	FieldClientDescription,
	FieldClientURI,
	FieldLogoURI,
	FieldRedirectUris,
	FieldPostLogoutRedirectUris,
	FieldOrganizationID,
	FieldPublic,
	FieldActive,
	FieldAllowedCorsOrigins,
	FieldAllowedGrantTypes,
	FieldTokenExpirySeconds,
	FieldRefreshTokenExpirySeconds,
	FieldAuthCodeExpirySeconds,
	FieldRequiresPkce,
	FieldRequiresConsent,
}

var (
	// ScopesPrimaryKey and ScopesColumn2 are the table columns denoting the
	// primary key for the scopes relation (M2M).
	ScopesPrimaryKey = []string{"oauth_client_id", "oauth_scope_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	ClientIDValidator func(string) error
	// ClientSecretValidator is a validator for the "client_secret" field. It is called by the builders before save.
	ClientSecretValidator func(string) error
	// ClientNameValidator is a validator for the "client_name" field. It is called by the builders before save.
	ClientNameValidator func(string) error
	// DefaultPublic holds the default value on creation for the "public" field.
	DefaultPublic bool
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultAllowedGrantTypes holds the default value on creation for the "allowed_grant_types" field.
	DefaultAllowedGrantTypes []string
	// DefaultTokenExpirySeconds holds the default value on creation for the "token_expiry_seconds" field.
	DefaultTokenExpirySeconds int
	// DefaultRefreshTokenExpirySeconds holds the default value on creation for the "refresh_token_expiry_seconds" field.
	DefaultRefreshTokenExpirySeconds int
	// DefaultAuthCodeExpirySeconds holds the default value on creation for the "auth_code_expiry_seconds" field.
	DefaultAuthCodeExpirySeconds int
	// DefaultRequiresPkce holds the default value on creation for the "requires_pkce" field.
	DefaultRequiresPkce bool
	// DefaultRequiresConsent holds the default value on creation for the "requires_consent" field.
	DefaultRequiresConsent bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
)

// OrderOption defines the ordering options for the OAuthClient queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByClientID orders the results by the client_id field.
func ByClientID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientID, opts...).ToFunc()
}

// ByClientSecret orders the results by the client_secret field.
func ByClientSecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientSecret, opts...).ToFunc()
}

// ByClientName orders the results by the client_name field.
func ByClientName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientName, opts...).ToFunc()
}

// ByClientDescription orders the results by the client_description field.
func ByClientDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientDescription, opts...).ToFunc()
}

// ByClientURI orders the results by the client_uri field.
func ByClientURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientURI, opts...).ToFunc()
}

// ByLogoURI orders the results by the logo_uri field.
func ByLogoURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogoURI, opts...).ToFunc()
}

// ByOrganizationID orders the results by the organization_id field.
func ByOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationID, opts...).ToFunc()
}

// ByPublic orders the results by the public field.
func ByPublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublic, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByTokenExpirySeconds orders the results by the token_expiry_seconds field.
func ByTokenExpirySeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenExpirySeconds, opts...).ToFunc()
}

// ByRefreshTokenExpirySeconds orders the results by the refresh_token_expiry_seconds field.
func ByRefreshTokenExpirySeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshTokenExpirySeconds, opts...).ToFunc()
}

// ByAuthCodeExpirySeconds orders the results by the auth_code_expiry_seconds field.
func ByAuthCodeExpirySeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthCodeExpirySeconds, opts...).ToFunc()
}

// ByRequiresPkce orders the results by the requires_pkce field.
func ByRequiresPkce(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequiresPkce, opts...).ToFunc()
}

// ByRequiresConsent orders the results by the requires_consent field.
func ByRequiresConsent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequiresConsent, opts...).ToFunc()
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByTokensCount orders the results by tokens count.
func ByTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTokensStep(), opts...)
	}
}

// ByTokens orders the results by tokens terms.
func ByTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAuthorizationsCount orders the results by authorizations count.
func ByAuthorizationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorizationsStep(), opts...)
	}
}

// ByAuthorizations orders the results by authorizations terms.
func ByAuthorizations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorizationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByScopesCount orders the results by scopes count.
func ByScopesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newScopesStep(), opts...)
	}
}

// ByScopes orders the results by scopes terms.
func ByScopes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScopesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
	)
}
func newTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TokensTable, TokensColumn),
	)
}
func newAuthorizationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorizationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthorizationsTable, AuthorizationsColumn),
	)
}
func newScopesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScopesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ScopesTable, ScopesPrimaryKey...),
	)
}
