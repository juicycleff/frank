// Copyright 2023-present XRaph LLC. All rights reserved.
// This source code is licensed under the XRaph LLC license found
// in the LICENSE file in the root directory of this source tree.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/xraph/frank/ent/oauthclient"
	"github.com/xraph/frank/ent/oauthtoken"
	"github.com/xraph/frank/ent/user"
	"github.com/rs/xid"
)

// OAuthToken is the model entity for the OAuthToken schema.
type OAuthToken struct {
	config `json:"-"`
	// ID of the ent.
	// ID of the entity
	ID xid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"-"`
	// RefreshToken holds the value of the "refresh_token" field.
	RefreshToken string `json:"-"`
	// TokenType holds the value of the "token_type" field.
	TokenType string `json:"token_type,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID xid.ID `json:"client_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID xid.ID `json:"user_id,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID xid.ID `json:"organization_id,omitempty"`
	// Scope names as strings for quick access
	ScopeNames []string `json:"scope_names,omitempty"`
	// ExpiresIn holds the value of the "expires_in" field.
	ExpiresIn int `json:"expires_in,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// RefreshTokenExpiresAt holds the value of the "refresh_token_expires_at" field.
	RefreshTokenExpiresAt *time.Time `json:"refresh_token_expires_at,omitempty"`
	// Revoked holds the value of the "revoked" field.
	Revoked bool `json:"revoked,omitempty"`
	// RevokedAt holds the value of the "revoked_at" field.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	// IPAddress holds the value of the "ip_address" field.
	IPAddress string `json:"ip_address,omitempty"`
	// UserAgent holds the value of the "user_agent" field.
	UserAgent string `json:"user_agent,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OAuthTokenQuery when eager-loading is set.
	Edges        OAuthTokenEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OAuthTokenEdges holds the relations/edges for other nodes in the graph.
type OAuthTokenEdges struct {
	// Client holds the value of the client edge.
	Client *OAuthClient `json:"client,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Scopes holds the value of the scopes edge.
	Scopes []*OAuthScope `json:"scopes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	namedScopes map[string][]*OAuthScope
}

// ClientOrErr returns the Client value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OAuthTokenEdges) ClientOrErr() (*OAuthClient, error) {
	if e.Client != nil {
		return e.Client, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: oauthclient.Label}
	}
	return nil, &NotLoadedError{edge: "client"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OAuthTokenEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ScopesOrErr returns the Scopes value or an error if the edge
// was not loaded in eager-loading.
func (e OAuthTokenEdges) ScopesOrErr() ([]*OAuthScope, error) {
	if e.loadedTypes[2] {
		return e.Scopes, nil
	}
	return nil, &NotLoadedError{edge: "scopes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OAuthToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauthtoken.FieldScopeNames:
			values[i] = new([]byte)
		case oauthtoken.FieldRevoked:
			values[i] = new(sql.NullBool)
		case oauthtoken.FieldExpiresIn:
			values[i] = new(sql.NullInt64)
		case oauthtoken.FieldAccessToken, oauthtoken.FieldRefreshToken, oauthtoken.FieldTokenType, oauthtoken.FieldIPAddress, oauthtoken.FieldUserAgent:
			values[i] = new(sql.NullString)
		case oauthtoken.FieldCreatedAt, oauthtoken.FieldUpdatedAt, oauthtoken.FieldDeletedAt, oauthtoken.FieldExpiresAt, oauthtoken.FieldRefreshTokenExpiresAt, oauthtoken.FieldRevokedAt:
			values[i] = new(sql.NullTime)
		case oauthtoken.FieldID, oauthtoken.FieldClientID, oauthtoken.FieldUserID, oauthtoken.FieldOrganizationID:
			values[i] = new(xid.ID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OAuthToken fields.
func (ot *OAuthToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauthtoken.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ot.ID = *value
			}
		case oauthtoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ot.CreatedAt = value.Time
			}
		case oauthtoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ot.UpdatedAt = value.Time
			}
		case oauthtoken.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ot.DeletedAt = value.Time
			}
		case oauthtoken.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				ot.AccessToken = value.String
			}
		case oauthtoken.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				ot.RefreshToken = value.String
			}
		case oauthtoken.FieldTokenType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_type", values[i])
			} else if value.Valid {
				ot.TokenType = value.String
			}
		case oauthtoken.FieldClientID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value != nil {
				ot.ClientID = *value
			}
		case oauthtoken.FieldUserID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ot.UserID = *value
			}
		case oauthtoken.FieldOrganizationID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				ot.OrganizationID = *value
			}
		case oauthtoken.FieldScopeNames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field scope_names", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ot.ScopeNames); err != nil {
					return fmt.Errorf("unmarshal field scope_names: %w", err)
				}
			}
		case oauthtoken.FieldExpiresIn:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expires_in", values[i])
			} else if value.Valid {
				ot.ExpiresIn = int(value.Int64)
			}
		case oauthtoken.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				ot.ExpiresAt = value.Time
			}
		case oauthtoken.FieldRefreshTokenExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token_expires_at", values[i])
			} else if value.Valid {
				ot.RefreshTokenExpiresAt = new(time.Time)
				*ot.RefreshTokenExpiresAt = value.Time
			}
		case oauthtoken.FieldRevoked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field revoked", values[i])
			} else if value.Valid {
				ot.Revoked = value.Bool
			}
		case oauthtoken.FieldRevokedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field revoked_at", values[i])
			} else if value.Valid {
				ot.RevokedAt = new(time.Time)
				*ot.RevokedAt = value.Time
			}
		case oauthtoken.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				ot.IPAddress = value.String
			}
		case oauthtoken.FieldUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_agent", values[i])
			} else if value.Valid {
				ot.UserAgent = value.String
			}
		default:
			ot.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OAuthToken.
// This includes values selected through modifiers, order, etc.
func (ot *OAuthToken) Value(name string) (ent.Value, error) {
	return ot.selectValues.Get(name)
}

// QueryClient queries the "client" edge of the OAuthToken entity.
func (ot *OAuthToken) QueryClient() *OAuthClientQuery {
	return NewOAuthTokenClient(ot.config).QueryClient(ot)
}

// QueryUser queries the "user" edge of the OAuthToken entity.
func (ot *OAuthToken) QueryUser() *UserQuery {
	return NewOAuthTokenClient(ot.config).QueryUser(ot)
}

// QueryScopes queries the "scopes" edge of the OAuthToken entity.
func (ot *OAuthToken) QueryScopes() *OAuthScopeQuery {
	return NewOAuthTokenClient(ot.config).QueryScopes(ot)
}

// Update returns a builder for updating this OAuthToken.
// Note that you need to call OAuthToken.Unwrap() before calling this method if this OAuthToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (ot *OAuthToken) Update() *OAuthTokenUpdateOne {
	return NewOAuthTokenClient(ot.config).UpdateOne(ot)
}

// Unwrap unwraps the OAuthToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ot *OAuthToken) Unwrap() *OAuthToken {
	_tx, ok := ot.config.driver.(*txDriver)
	if !ok {
		panic("ent: OAuthToken is not a transactional entity")
	}
	ot.config.driver = _tx.drv
	return ot
}

// String implements the fmt.Stringer.
func (ot *OAuthToken) String() string {
	var builder strings.Builder
	builder.WriteString("OAuthToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ot.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ot.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ot.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ot.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("access_token=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("refresh_token=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("token_type=")
	builder.WriteString(ot.TokenType)
	builder.WriteString(", ")
	builder.WriteString("client_id=")
	builder.WriteString(fmt.Sprintf("%v", ot.ClientID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ot.UserID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", ot.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("scope_names=")
	builder.WriteString(fmt.Sprintf("%v", ot.ScopeNames))
	builder.WriteString(", ")
	builder.WriteString("expires_in=")
	builder.WriteString(fmt.Sprintf("%v", ot.ExpiresIn))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(ot.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ot.RefreshTokenExpiresAt; v != nil {
		builder.WriteString("refresh_token_expires_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("revoked=")
	builder.WriteString(fmt.Sprintf("%v", ot.Revoked))
	builder.WriteString(", ")
	if v := ot.RevokedAt; v != nil {
		builder.WriteString("revoked_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("ip_address=")
	builder.WriteString(ot.IPAddress)
	builder.WriteString(", ")
	builder.WriteString("user_agent=")
	builder.WriteString(ot.UserAgent)
	builder.WriteByte(')')
	return builder.String()
}

// NamedScopes returns the Scopes named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ot *OAuthToken) NamedScopes(name string) ([]*OAuthScope, error) {
	if ot.Edges.namedScopes == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ot.Edges.namedScopes[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ot *OAuthToken) appendNamedScopes(name string, edges ...*OAuthScope) {
	if ot.Edges.namedScopes == nil {
		ot.Edges.namedScopes = make(map[string][]*OAuthScope)
	}
	if len(edges) == 0 {
		ot.Edges.namedScopes[name] = []*OAuthScope{}
	} else {
		ot.Edges.namedScopes[name] = append(ot.Edges.namedScopes[name], edges...)
	}
}

// OAuthTokens is a parsable slice of OAuthToken.
type OAuthTokens []*OAuthToken
