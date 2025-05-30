// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/juicycleff/frank/ent/oauthscope"
)

// OAuthScope holds the schema definition for the OAuthScope entity.
type OAuthScope struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Whether this scope is included by default
	DefaultScope bool `json:"default_scope,omitempty"`
	// Whether this scope can be requested by any client
	Public bool `json:"public,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OAuthScopeQuery when eager-loading is set.
	Edges        OAuthScopeEdges `json:"edges"`
	selectValues sql.SelectValues
} //@name OAuthScope


// OAuthScopeEdges holds the relations/edges for other nodes in the graph.
type OAuthScopeEdges struct {
	// Clients holds the value of the clients edge.
	Clients []*OAuthClient `json:"clients,omitempty"`
	// Tokens holds the value of the tokens edge.
	Tokens []*OAuthToken `json:"tokens,omitempty"`
	// Authorizations holds the value of the authorizations edge.
	Authorizations []*OAuthAuthorization `json:"authorizations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ClientsOrErr returns the Clients value or an error if the edge
// was not loaded in eager-loading.
func (e OAuthScopeEdges) ClientsOrErr() ([]*OAuthClient, error) {
	if e.loadedTypes[0] {
		return e.Clients, nil
	}
	return nil, &NotLoadedError{edge: "clients"}
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e OAuthScopeEdges) TokensOrErr() ([]*OAuthToken, error) {
	if e.loadedTypes[1] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// AuthorizationsOrErr returns the Authorizations value or an error if the edge
// was not loaded in eager-loading.
func (e OAuthScopeEdges) AuthorizationsOrErr() ([]*OAuthAuthorization, error) {
	if e.loadedTypes[2] {
		return e.Authorizations, nil
	}
	return nil, &NotLoadedError{edge: "authorizations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OAuthScope) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauthscope.FieldDefaultScope, oauthscope.FieldPublic:
			values[i] = new(sql.NullBool)
		case oauthscope.FieldID, oauthscope.FieldName, oauthscope.FieldDescription:
			values[i] = new(sql.NullString)
		case oauthscope.FieldCreatedAt, oauthscope.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OAuthScope fields.
func (os *OAuthScope) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauthscope.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				os.ID = value.String
			}
		case oauthscope.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				os.CreatedAt = value.Time
			}
		case oauthscope.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				os.UpdatedAt = value.Time
			}
		case oauthscope.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				os.Name = value.String
			}
		case oauthscope.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				os.Description = value.String
			}
		case oauthscope.FieldDefaultScope:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field default_scope", values[i])
			} else if value.Valid {
				os.DefaultScope = value.Bool
			}
		case oauthscope.FieldPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public", values[i])
			} else if value.Valid {
				os.Public = value.Bool
			}
		default:
			os.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OAuthScope.
// This includes values selected through modifiers, order, etc.
func (os *OAuthScope) Value(name string) (ent.Value, error) {
	return os.selectValues.Get(name)
}

// QueryClients queries the "clients" edge of the OAuthScope entity.
func (os *OAuthScope) QueryClients() *OAuthClientQuery {
	return NewOAuthScopeClient(os.config).QueryClients(os)
}

// QueryTokens queries the "tokens" edge of the OAuthScope entity.
func (os *OAuthScope) QueryTokens() *OAuthTokenQuery {
	return NewOAuthScopeClient(os.config).QueryTokens(os)
}

// QueryAuthorizations queries the "authorizations" edge of the OAuthScope entity.
func (os *OAuthScope) QueryAuthorizations() *OAuthAuthorizationQuery {
	return NewOAuthScopeClient(os.config).QueryAuthorizations(os)
}

// Update returns a builder for updating this OAuthScope.
// Note that you need to call OAuthScope.Unwrap() before calling this method if this OAuthScope
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OAuthScope) Update() *OAuthScopeUpdateOne {
	return NewOAuthScopeClient(os.config).UpdateOne(os)
}

// Unwrap unwraps the OAuthScope entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *OAuthScope) Unwrap() *OAuthScope {
	_tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("ent: OAuthScope is not a transactional entity")
	}
	os.config.driver = _tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OAuthScope) String() string {
	var builder strings.Builder
	builder.WriteString("OAuthScope(")
	builder.WriteString(fmt.Sprintf("id=%v, ", os.ID))
	builder.WriteString("created_at=")
	builder.WriteString(os.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(os.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(os.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(os.Description)
	builder.WriteString(", ")
	builder.WriteString("default_scope=")
	builder.WriteString(fmt.Sprintf("%v", os.DefaultScope))
	builder.WriteString(", ")
	builder.WriteString("public=")
	builder.WriteString(fmt.Sprintf("%v", os.Public))
	builder.WriteByte(')')
	return builder.String()
}

// OAuthScopes is a parsable slice of OAuthScope.
type OAuthScopes []*OAuthScope
