// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/juicycleff/frank/ent/identityprovider"
	"github.com/juicycleff/frank/ent/organization"
)

// IdentityProvider holds the schema definition for the IdentityProvider entity.
type IdentityProvider struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID string `json:"organization_id,omitempty"`
	// Type of IdP: 'oauth2', 'oidc', 'saml'
	ProviderType string `json:"provider_type,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret holds the value of the "client_secret" field.
	ClientSecret string `json:"-"`
	// Issuer holds the value of the "issuer" field.
	Issuer string `json:"issuer,omitempty"`
	// AuthorizationEndpoint holds the value of the "authorization_endpoint" field.
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty"`
	// TokenEndpoint holds the value of the "token_endpoint" field.
	TokenEndpoint string `json:"token_endpoint,omitempty"`
	// UserinfoEndpoint holds the value of the "userinfo_endpoint" field.
	UserinfoEndpoint string `json:"userinfo_endpoint,omitempty"`
	// JwksURI holds the value of the "jwks_uri" field.
	JwksURI string `json:"jwks_uri,omitempty"`
	// MetadataURL holds the value of the "metadata_url" field.
	MetadataURL string `json:"metadata_url,omitempty"`
	// RedirectURI holds the value of the "redirect_uri" field.
	RedirectURI string `json:"redirect_uri,omitempty"`
	// Certificate holds the value of the "certificate" field.
	Certificate string `json:"-"`
	// PrivateKey holds the value of the "private_key" field.
	PrivateKey string `json:"-"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Primary holds the value of the "primary" field.
	Primary bool `json:"primary,omitempty"`
	// Domains holds the value of the "domains" field.
	Domains []string `json:"domains,omitempty"`
	// AttributesMapping holds the value of the "attributes_mapping" field.
	AttributesMapping map[string]string `json:"attributes_mapping,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IdentityProviderQuery when eager-loading is set.
	Edges        IdentityProviderEdges `json:"edges"`
	selectValues sql.SelectValues
} //@name IdentityProvider


// IdentityProviderEdges holds the relations/edges for other nodes in the graph.
type IdentityProviderEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IdentityProviderEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IdentityProvider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case identityprovider.FieldDomains, identityprovider.FieldAttributesMapping, identityprovider.FieldMetadata:
			values[i] = new([]byte)
		case identityprovider.FieldActive, identityprovider.FieldPrimary:
			values[i] = new(sql.NullBool)
		case identityprovider.FieldID, identityprovider.FieldName, identityprovider.FieldOrganizationID, identityprovider.FieldProviderType, identityprovider.FieldClientID, identityprovider.FieldClientSecret, identityprovider.FieldIssuer, identityprovider.FieldAuthorizationEndpoint, identityprovider.FieldTokenEndpoint, identityprovider.FieldUserinfoEndpoint, identityprovider.FieldJwksURI, identityprovider.FieldMetadataURL, identityprovider.FieldRedirectURI, identityprovider.FieldCertificate, identityprovider.FieldPrivateKey:
			values[i] = new(sql.NullString)
		case identityprovider.FieldCreatedAt, identityprovider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IdentityProvider fields.
func (ip *IdentityProvider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case identityprovider.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ip.ID = value.String
			}
		case identityprovider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ip.CreatedAt = value.Time
			}
		case identityprovider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ip.UpdatedAt = value.Time
			}
		case identityprovider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ip.Name = value.String
			}
		case identityprovider.FieldOrganizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				ip.OrganizationID = value.String
			}
		case identityprovider.FieldProviderType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_type", values[i])
			} else if value.Valid {
				ip.ProviderType = value.String
			}
		case identityprovider.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				ip.ClientID = value.String
			}
		case identityprovider.FieldClientSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_secret", values[i])
			} else if value.Valid {
				ip.ClientSecret = value.String
			}
		case identityprovider.FieldIssuer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field issuer", values[i])
			} else if value.Valid {
				ip.Issuer = value.String
			}
		case identityprovider.FieldAuthorizationEndpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authorization_endpoint", values[i])
			} else if value.Valid {
				ip.AuthorizationEndpoint = value.String
			}
		case identityprovider.FieldTokenEndpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_endpoint", values[i])
			} else if value.Valid {
				ip.TokenEndpoint = value.String
			}
		case identityprovider.FieldUserinfoEndpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userinfo_endpoint", values[i])
			} else if value.Valid {
				ip.UserinfoEndpoint = value.String
			}
		case identityprovider.FieldJwksURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jwks_uri", values[i])
			} else if value.Valid {
				ip.JwksURI = value.String
			}
		case identityprovider.FieldMetadataURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field metadata_url", values[i])
			} else if value.Valid {
				ip.MetadataURL = value.String
			}
		case identityprovider.FieldRedirectURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_uri", values[i])
			} else if value.Valid {
				ip.RedirectURI = value.String
			}
		case identityprovider.FieldCertificate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field certificate", values[i])
			} else if value.Valid {
				ip.Certificate = value.String
			}
		case identityprovider.FieldPrivateKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key", values[i])
			} else if value.Valid {
				ip.PrivateKey = value.String
			}
		case identityprovider.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				ip.Active = value.Bool
			}
		case identityprovider.FieldPrimary:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field primary", values[i])
			} else if value.Valid {
				ip.Primary = value.Bool
			}
		case identityprovider.FieldDomains:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field domains", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ip.Domains); err != nil {
					return fmt.Errorf("unmarshal field domains: %w", err)
				}
			}
		case identityprovider.FieldAttributesMapping:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field attributes_mapping", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ip.AttributesMapping); err != nil {
					return fmt.Errorf("unmarshal field attributes_mapping: %w", err)
				}
			}
		case identityprovider.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ip.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		default:
			ip.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IdentityProvider.
// This includes values selected through modifiers, order, etc.
func (ip *IdentityProvider) Value(name string) (ent.Value, error) {
	return ip.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the IdentityProvider entity.
func (ip *IdentityProvider) QueryOrganization() *OrganizationQuery {
	return NewIdentityProviderClient(ip.config).QueryOrganization(ip)
}

// Update returns a builder for updating this IdentityProvider.
// Note that you need to call IdentityProvider.Unwrap() before calling this method if this IdentityProvider
// was returned from a transaction, and the transaction was committed or rolled back.
func (ip *IdentityProvider) Update() *IdentityProviderUpdateOne {
	return NewIdentityProviderClient(ip.config).UpdateOne(ip)
}

// Unwrap unwraps the IdentityProvider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ip *IdentityProvider) Unwrap() *IdentityProvider {
	_tx, ok := ip.config.driver.(*txDriver)
	if !ok {
		panic("ent: IdentityProvider is not a transactional entity")
	}
	ip.config.driver = _tx.drv
	return ip
}

// String implements the fmt.Stringer.
func (ip *IdentityProvider) String() string {
	var builder strings.Builder
	builder.WriteString("IdentityProvider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ip.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ip.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ip.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ip.Name)
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(ip.OrganizationID)
	builder.WriteString(", ")
	builder.WriteString("provider_type=")
	builder.WriteString(ip.ProviderType)
	builder.WriteString(", ")
	builder.WriteString("client_id=")
	builder.WriteString(ip.ClientID)
	builder.WriteString(", ")
	builder.WriteString("client_secret=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("issuer=")
	builder.WriteString(ip.Issuer)
	builder.WriteString(", ")
	builder.WriteString("authorization_endpoint=")
	builder.WriteString(ip.AuthorizationEndpoint)
	builder.WriteString(", ")
	builder.WriteString("token_endpoint=")
	builder.WriteString(ip.TokenEndpoint)
	builder.WriteString(", ")
	builder.WriteString("userinfo_endpoint=")
	builder.WriteString(ip.UserinfoEndpoint)
	builder.WriteString(", ")
	builder.WriteString("jwks_uri=")
	builder.WriteString(ip.JwksURI)
	builder.WriteString(", ")
	builder.WriteString("metadata_url=")
	builder.WriteString(ip.MetadataURL)
	builder.WriteString(", ")
	builder.WriteString("redirect_uri=")
	builder.WriteString(ip.RedirectURI)
	builder.WriteString(", ")
	builder.WriteString("certificate=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("private_key=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", ip.Active))
	builder.WriteString(", ")
	builder.WriteString("primary=")
	builder.WriteString(fmt.Sprintf("%v", ip.Primary))
	builder.WriteString(", ")
	builder.WriteString("domains=")
	builder.WriteString(fmt.Sprintf("%v", ip.Domains))
	builder.WriteString(", ")
	builder.WriteString("attributes_mapping=")
	builder.WriteString(fmt.Sprintf("%v", ip.AttributesMapping))
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", ip.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// IdentityProviders is a parsable slice of IdentityProvider.
type IdentityProviders []*IdentityProvider
