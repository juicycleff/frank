// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/juicycleff/frank/ent/user"
)

// User holds the schema definition for the User entity.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash string `json:"-"`
	// EmailVerified holds the value of the "email_verified" field.
	EmailVerified bool `json:"email_verified,omitempty"`
	// PhoneVerified holds the value of the "phone_verified" field.
	PhoneVerified bool `json:"phone_verified,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// LastLogin holds the value of the "last_login" field.
	LastLogin *time.Time `json:"last_login,omitempty"`
	// LastPasswordChange holds the value of the "last_password_change" field.
	LastPasswordChange *time.Time `json:"last_password_change,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// ProfileImageURL holds the value of the "profile_image_url" field.
	ProfileImageURL string `json:"profile_image_url,omitempty"`
	// PrimaryOrganizationID holds the value of the "primary_organization_id" field.
	PrimaryOrganizationID string `json:"primary_organization_id,omitempty"`
	// Locale holds the value of the "locale" field.
	Locale string `json:"locale,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
} //@name User


// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Sessions holds the value of the sessions edge.
	Sessions []*Session `json:"sessions,omitempty"`
	// APIKeys holds the value of the api_keys edge.
	APIKeys []*ApiKey `json:"api_keys,omitempty"`
	// Organizations holds the value of the organizations edge.
	Organizations []*Organization `json:"organizations,omitempty"`
	// MfaMethods holds the value of the mfa_methods edge.
	MfaMethods []*MFA `json:"mfa_methods,omitempty"`
	// Passkeys holds the value of the passkeys edge.
	Passkeys []*Passkey `json:"passkeys,omitempty"`
	// OauthTokens holds the value of the oauth_tokens edge.
	OauthTokens []*OAuthToken `json:"oauth_tokens,omitempty"`
	// OauthAuthorizations holds the value of the oauth_authorizations edge.
	OauthAuthorizations []*OAuthAuthorization `json:"oauth_authorizations,omitempty"`
	// Verifications holds the value of the verifications edge.
	Verifications []*Verification `json:"verifications,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// SessionsOrErr returns the Sessions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SessionsOrErr() ([]*Session, error) {
	if e.loadedTypes[0] {
		return e.Sessions, nil
	}
	return nil, &NotLoadedError{edge: "sessions"}
}

// APIKeysOrErr returns the APIKeys value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) APIKeysOrErr() ([]*ApiKey, error) {
	if e.loadedTypes[1] {
		return e.APIKeys, nil
	}
	return nil, &NotLoadedError{edge: "api_keys"}
}

// OrganizationsOrErr returns the Organizations value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrganizationsOrErr() ([]*Organization, error) {
	if e.loadedTypes[2] {
		return e.Organizations, nil
	}
	return nil, &NotLoadedError{edge: "organizations"}
}

// MfaMethodsOrErr returns the MfaMethods value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MfaMethodsOrErr() ([]*MFA, error) {
	if e.loadedTypes[3] {
		return e.MfaMethods, nil
	}
	return nil, &NotLoadedError{edge: "mfa_methods"}
}

// PasskeysOrErr returns the Passkeys value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PasskeysOrErr() ([]*Passkey, error) {
	if e.loadedTypes[4] {
		return e.Passkeys, nil
	}
	return nil, &NotLoadedError{edge: "passkeys"}
}

// OauthTokensOrErr returns the OauthTokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OauthTokensOrErr() ([]*OAuthToken, error) {
	if e.loadedTypes[5] {
		return e.OauthTokens, nil
	}
	return nil, &NotLoadedError{edge: "oauth_tokens"}
}

// OauthAuthorizationsOrErr returns the OauthAuthorizations value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OauthAuthorizationsOrErr() ([]*OAuthAuthorization, error) {
	if e.loadedTypes[6] {
		return e.OauthAuthorizations, nil
	}
	return nil, &NotLoadedError{edge: "oauth_authorizations"}
}

// VerificationsOrErr returns the Verifications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VerificationsOrErr() ([]*Verification, error) {
	if e.loadedTypes[7] {
		return e.Verifications, nil
	}
	return nil, &NotLoadedError{edge: "verifications"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[8] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldMetadata:
			values[i] = new([]byte)
		case user.FieldEmailVerified, user.FieldPhoneVerified, user.FieldActive:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldEmail, user.FieldPhoneNumber, user.FieldFirstName, user.FieldLastName, user.FieldPasswordHash, user.FieldProfileImageURL, user.FieldPrimaryOrganizationID, user.FieldLocale:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldLastLogin, user.FieldLastPasswordChange:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				u.PhoneNumber = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				u.PasswordHash = value.String
			}
		case user.FieldEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified", values[i])
			} else if value.Valid {
				u.EmailVerified = value.Bool
			}
		case user.FieldPhoneVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field phone_verified", values[i])
			} else if value.Valid {
				u.PhoneVerified = value.Bool
			}
		case user.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				u.Active = value.Bool
			}
		case user.FieldLastLogin:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login", values[i])
			} else if value.Valid {
				u.LastLogin = new(time.Time)
				*u.LastLogin = value.Time
			}
		case user.FieldLastPasswordChange:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_password_change", values[i])
			} else if value.Valid {
				u.LastPasswordChange = new(time.Time)
				*u.LastPasswordChange = value.Time
			}
		case user.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case user.FieldProfileImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_image_url", values[i])
			} else if value.Valid {
				u.ProfileImageURL = value.String
			}
		case user.FieldPrimaryOrganizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_organization_id", values[i])
			} else if value.Valid {
				u.PrimaryOrganizationID = value.String
			}
		case user.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				u.Locale = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QuerySessions queries the "sessions" edge of the User entity.
func (u *User) QuerySessions() *SessionQuery {
	return NewUserClient(u.config).QuerySessions(u)
}

// QueryAPIKeys queries the "api_keys" edge of the User entity.
func (u *User) QueryAPIKeys() *ApiKeyQuery {
	return NewUserClient(u.config).QueryAPIKeys(u)
}

// QueryOrganizations queries the "organizations" edge of the User entity.
func (u *User) QueryOrganizations() *OrganizationQuery {
	return NewUserClient(u.config).QueryOrganizations(u)
}

// QueryMfaMethods queries the "mfa_methods" edge of the User entity.
func (u *User) QueryMfaMethods() *MFAQuery {
	return NewUserClient(u.config).QueryMfaMethods(u)
}

// QueryPasskeys queries the "passkeys" edge of the User entity.
func (u *User) QueryPasskeys() *PasskeyQuery {
	return NewUserClient(u.config).QueryPasskeys(u)
}

// QueryOauthTokens queries the "oauth_tokens" edge of the User entity.
func (u *User) QueryOauthTokens() *OAuthTokenQuery {
	return NewUserClient(u.config).QueryOauthTokens(u)
}

// QueryOauthAuthorizations queries the "oauth_authorizations" edge of the User entity.
func (u *User) QueryOauthAuthorizations() *OAuthAuthorizationQuery {
	return NewUserClient(u.config).QueryOauthAuthorizations(u)
}

// QueryVerifications queries the "verifications" edge of the User entity.
func (u *User) QueryVerifications() *VerificationQuery {
	return NewUserClient(u.config).QueryVerifications(u)
}

// QueryRoles queries the "roles" edge of the User entity.
func (u *User) QueryRoles() *RoleQuery {
	return NewUserClient(u.config).QueryRoles(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(u.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("password_hash=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("email_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.EmailVerified))
	builder.WriteString(", ")
	builder.WriteString("phone_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.PhoneVerified))
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", u.Active))
	builder.WriteString(", ")
	if v := u.LastLogin; v != nil {
		builder.WriteString("last_login=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.LastPasswordChange; v != nil {
		builder.WriteString("last_password_change=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", u.Metadata))
	builder.WriteString(", ")
	builder.WriteString("profile_image_url=")
	builder.WriteString(u.ProfileImageURL)
	builder.WriteString(", ")
	builder.WriteString("primary_organization_id=")
	builder.WriteString(u.PrimaryOrganizationID)
	builder.WriteString(", ")
	builder.WriteString("locale=")
	builder.WriteString(u.Locale)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
