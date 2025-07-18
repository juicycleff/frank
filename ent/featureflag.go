// Copyright 2023-present XRaph LLC. All rights reserved.
// This source code is licensed under the XRaph LLC license found
// in the LICENSE file in the root directory of this source tree.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/xraph/frank/ent/featureflag"
	"github.com/rs/xid"
)

// FeatureFlag is the model entity for the FeatureFlag schema.
type FeatureFlag struct {
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
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Enabled holds the value of the "enabled" field.
	Enabled bool `json:"enabled,omitempty"`
	// Whether this feature is only available for premium plans
	IsPremium bool `json:"is_premium,omitempty"`
	// Which component of the auth system this feature belongs to
	Component featureflag.Component `json:"component,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FeatureFlagQuery when eager-loading is set.
	Edges        FeatureFlagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FeatureFlagEdges holds the relations/edges for other nodes in the graph.
type FeatureFlagEdges struct {
	// OrganizationFeatures holds the value of the organization_features edge.
	OrganizationFeatures []*OrganizationFeature `json:"organization_features,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes               [1]bool
	namedOrganizationFeatures map[string][]*OrganizationFeature
}

// OrganizationFeaturesOrErr returns the OrganizationFeatures value or an error if the edge
// was not loaded in eager-loading.
func (e FeatureFlagEdges) OrganizationFeaturesOrErr() ([]*OrganizationFeature, error) {
	if e.loadedTypes[0] {
		return e.OrganizationFeatures, nil
	}
	return nil, &NotLoadedError{edge: "organization_features"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeatureFlag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case featureflag.FieldEnabled, featureflag.FieldIsPremium:
			values[i] = new(sql.NullBool)
		case featureflag.FieldName, featureflag.FieldKey, featureflag.FieldDescription, featureflag.FieldComponent:
			values[i] = new(sql.NullString)
		case featureflag.FieldCreatedAt, featureflag.FieldUpdatedAt, featureflag.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case featureflag.FieldID:
			values[i] = new(xid.ID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeatureFlag fields.
func (ff *FeatureFlag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case featureflag.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ff.ID = *value
			}
		case featureflag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ff.CreatedAt = value.Time
			}
		case featureflag.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ff.UpdatedAt = value.Time
			}
		case featureflag.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ff.DeletedAt = value.Time
			}
		case featureflag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ff.Name = value.String
			}
		case featureflag.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				ff.Key = value.String
			}
		case featureflag.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ff.Description = value.String
			}
		case featureflag.FieldEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled", values[i])
			} else if value.Valid {
				ff.Enabled = value.Bool
			}
		case featureflag.FieldIsPremium:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_premium", values[i])
			} else if value.Valid {
				ff.IsPremium = value.Bool
			}
		case featureflag.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				ff.Component = featureflag.Component(value.String)
			}
		default:
			ff.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FeatureFlag.
// This includes values selected through modifiers, order, etc.
func (ff *FeatureFlag) Value(name string) (ent.Value, error) {
	return ff.selectValues.Get(name)
}

// QueryOrganizationFeatures queries the "organization_features" edge of the FeatureFlag entity.
func (ff *FeatureFlag) QueryOrganizationFeatures() *OrganizationFeatureQuery {
	return NewFeatureFlagClient(ff.config).QueryOrganizationFeatures(ff)
}

// Update returns a builder for updating this FeatureFlag.
// Note that you need to call FeatureFlag.Unwrap() before calling this method if this FeatureFlag
// was returned from a transaction, and the transaction was committed or rolled back.
func (ff *FeatureFlag) Update() *FeatureFlagUpdateOne {
	return NewFeatureFlagClient(ff.config).UpdateOne(ff)
}

// Unwrap unwraps the FeatureFlag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ff *FeatureFlag) Unwrap() *FeatureFlag {
	_tx, ok := ff.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeatureFlag is not a transactional entity")
	}
	ff.config.driver = _tx.drv
	return ff
}

// String implements the fmt.Stringer.
func (ff *FeatureFlag) String() string {
	var builder strings.Builder
	builder.WriteString("FeatureFlag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ff.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ff.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ff.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ff.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ff.Name)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(ff.Key)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ff.Description)
	builder.WriteString(", ")
	builder.WriteString("enabled=")
	builder.WriteString(fmt.Sprintf("%v", ff.Enabled))
	builder.WriteString(", ")
	builder.WriteString("is_premium=")
	builder.WriteString(fmt.Sprintf("%v", ff.IsPremium))
	builder.WriteString(", ")
	builder.WriteString("component=")
	builder.WriteString(fmt.Sprintf("%v", ff.Component))
	builder.WriteByte(')')
	return builder.String()
}

// NamedOrganizationFeatures returns the OrganizationFeatures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ff *FeatureFlag) NamedOrganizationFeatures(name string) ([]*OrganizationFeature, error) {
	if ff.Edges.namedOrganizationFeatures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ff.Edges.namedOrganizationFeatures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ff *FeatureFlag) appendNamedOrganizationFeatures(name string, edges ...*OrganizationFeature) {
	if ff.Edges.namedOrganizationFeatures == nil {
		ff.Edges.namedOrganizationFeatures = make(map[string][]*OrganizationFeature)
	}
	if len(edges) == 0 {
		ff.Edges.namedOrganizationFeatures[name] = []*OrganizationFeature{}
	} else {
		ff.Edges.namedOrganizationFeatures[name] = append(ff.Edges.namedOrganizationFeatures[name], edges...)
	}
}

// FeatureFlags is a parsable slice of FeatureFlag.
type FeatureFlags []*FeatureFlag
