// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/juicycleff/frank/ent/emailtemplate"
)

// EmailTemplateCreate is the builder for creating a EmailTemplate entity.
type EmailTemplateCreate struct {
	config
	mutation *EmailTemplateMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (etc *EmailTemplateCreate) SetCreatedAt(t time.Time) *EmailTemplateCreate {
	etc.mutation.SetCreatedAt(t)
	return etc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableCreatedAt(t *time.Time) *EmailTemplateCreate {
	if t != nil {
		etc.SetCreatedAt(*t)
	}
	return etc
}

// SetUpdatedAt sets the "updated_at" field.
func (etc *EmailTemplateCreate) SetUpdatedAt(t time.Time) *EmailTemplateCreate {
	etc.mutation.SetUpdatedAt(t)
	return etc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableUpdatedAt(t *time.Time) *EmailTemplateCreate {
	if t != nil {
		etc.SetUpdatedAt(*t)
	}
	return etc
}

// SetName sets the "name" field.
func (etc *EmailTemplateCreate) SetName(s string) *EmailTemplateCreate {
	etc.mutation.SetName(s)
	return etc
}

// SetSubject sets the "subject" field.
func (etc *EmailTemplateCreate) SetSubject(s string) *EmailTemplateCreate {
	etc.mutation.SetSubject(s)
	return etc
}

// SetType sets the "type" field.
func (etc *EmailTemplateCreate) SetType(s string) *EmailTemplateCreate {
	etc.mutation.SetType(s)
	return etc
}

// SetHTMLContent sets the "html_content" field.
func (etc *EmailTemplateCreate) SetHTMLContent(s string) *EmailTemplateCreate {
	etc.mutation.SetHTMLContent(s)
	return etc
}

// SetTextContent sets the "text_content" field.
func (etc *EmailTemplateCreate) SetTextContent(s string) *EmailTemplateCreate {
	etc.mutation.SetTextContent(s)
	return etc
}

// SetNillableTextContent sets the "text_content" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableTextContent(s *string) *EmailTemplateCreate {
	if s != nil {
		etc.SetTextContent(*s)
	}
	return etc
}

// SetOrganizationID sets the "organization_id" field.
func (etc *EmailTemplateCreate) SetOrganizationID(s string) *EmailTemplateCreate {
	etc.mutation.SetOrganizationID(s)
	return etc
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableOrganizationID(s *string) *EmailTemplateCreate {
	if s != nil {
		etc.SetOrganizationID(*s)
	}
	return etc
}

// SetActive sets the "active" field.
func (etc *EmailTemplateCreate) SetActive(b bool) *EmailTemplateCreate {
	etc.mutation.SetActive(b)
	return etc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableActive(b *bool) *EmailTemplateCreate {
	if b != nil {
		etc.SetActive(*b)
	}
	return etc
}

// SetSystem sets the "system" field.
func (etc *EmailTemplateCreate) SetSystem(b bool) *EmailTemplateCreate {
	etc.mutation.SetSystem(b)
	return etc
}

// SetNillableSystem sets the "system" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableSystem(b *bool) *EmailTemplateCreate {
	if b != nil {
		etc.SetSystem(*b)
	}
	return etc
}

// SetLocale sets the "locale" field.
func (etc *EmailTemplateCreate) SetLocale(s string) *EmailTemplateCreate {
	etc.mutation.SetLocale(s)
	return etc
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableLocale(s *string) *EmailTemplateCreate {
	if s != nil {
		etc.SetLocale(*s)
	}
	return etc
}

// SetMetadata sets the "metadata" field.
func (etc *EmailTemplateCreate) SetMetadata(m map[string]interface{}) *EmailTemplateCreate {
	etc.mutation.SetMetadata(m)
	return etc
}

// SetID sets the "id" field.
func (etc *EmailTemplateCreate) SetID(s string) *EmailTemplateCreate {
	etc.mutation.SetID(s)
	return etc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableID(s *string) *EmailTemplateCreate {
	if s != nil {
		etc.SetID(*s)
	}
	return etc
}

// Mutation returns the EmailTemplateMutation object of the builder.
func (etc *EmailTemplateCreate) Mutation() *EmailTemplateMutation {
	return etc.mutation
}

// Save creates the EmailTemplate in the database.
func (etc *EmailTemplateCreate) Save(ctx context.Context) (*EmailTemplate, error) {
	etc.defaults()
	return withHooks(ctx, etc.sqlSave, etc.mutation, etc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (etc *EmailTemplateCreate) SaveX(ctx context.Context) *EmailTemplate {
	v, err := etc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etc *EmailTemplateCreate) Exec(ctx context.Context) error {
	_, err := etc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etc *EmailTemplateCreate) ExecX(ctx context.Context) {
	if err := etc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etc *EmailTemplateCreate) defaults() {
	if _, ok := etc.mutation.CreatedAt(); !ok {
		v := emailtemplate.DefaultCreatedAt()
		etc.mutation.SetCreatedAt(v)
	}
	if _, ok := etc.mutation.UpdatedAt(); !ok {
		v := emailtemplate.DefaultUpdatedAt()
		etc.mutation.SetUpdatedAt(v)
	}
	if _, ok := etc.mutation.Active(); !ok {
		v := emailtemplate.DefaultActive
		etc.mutation.SetActive(v)
	}
	if _, ok := etc.mutation.System(); !ok {
		v := emailtemplate.DefaultSystem
		etc.mutation.SetSystem(v)
	}
	if _, ok := etc.mutation.Locale(); !ok {
		v := emailtemplate.DefaultLocale
		etc.mutation.SetLocale(v)
	}
	if _, ok := etc.mutation.ID(); !ok {
		v := emailtemplate.DefaultID()
		etc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etc *EmailTemplateCreate) check() error {
	if _, ok := etc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EmailTemplate.created_at"`)}
	}
	if _, ok := etc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EmailTemplate.updated_at"`)}
	}
	if _, ok := etc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "EmailTemplate.name"`)}
	}
	if v, ok := etc.mutation.Name(); ok {
		if err := emailtemplate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EmailTemplate.name": %w`, err)}
		}
	}
	if _, ok := etc.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "EmailTemplate.subject"`)}
	}
	if v, ok := etc.mutation.Subject(); ok {
		if err := emailtemplate.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "subject", err: fmt.Errorf(`ent: validator failed for field "EmailTemplate.subject": %w`, err)}
		}
	}
	if _, ok := etc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "EmailTemplate.type"`)}
	}
	if v, ok := etc.mutation.GetType(); ok {
		if err := emailtemplate.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "EmailTemplate.type": %w`, err)}
		}
	}
	if _, ok := etc.mutation.HTMLContent(); !ok {
		return &ValidationError{Name: "html_content", err: errors.New(`ent: missing required field "EmailTemplate.html_content"`)}
	}
	if v, ok := etc.mutation.HTMLContent(); ok {
		if err := emailtemplate.HTMLContentValidator(v); err != nil {
			return &ValidationError{Name: "html_content", err: fmt.Errorf(`ent: validator failed for field "EmailTemplate.html_content": %w`, err)}
		}
	}
	if _, ok := etc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "EmailTemplate.active"`)}
	}
	if _, ok := etc.mutation.System(); !ok {
		return &ValidationError{Name: "system", err: errors.New(`ent: missing required field "EmailTemplate.system"`)}
	}
	if _, ok := etc.mutation.Locale(); !ok {
		return &ValidationError{Name: "locale", err: errors.New(`ent: missing required field "EmailTemplate.locale"`)}
	}
	return nil
}

func (etc *EmailTemplateCreate) sqlSave(ctx context.Context) (*EmailTemplate, error) {
	if err := etc.check(); err != nil {
		return nil, err
	}
	_node, _spec := etc.createSpec()
	if err := sqlgraph.CreateNode(ctx, etc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected EmailTemplate.ID type: %T", _spec.ID.Value)
		}
	}
	etc.mutation.id = &_node.ID
	etc.mutation.done = true
	return _node, nil
}

func (etc *EmailTemplateCreate) createSpec() (*EmailTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailTemplate{config: etc.config}
		_spec = sqlgraph.NewCreateSpec(emailtemplate.Table, sqlgraph.NewFieldSpec(emailtemplate.FieldID, field.TypeString))
	)
	if id, ok := etc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := etc.mutation.CreatedAt(); ok {
		_spec.SetField(emailtemplate.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := etc.mutation.UpdatedAt(); ok {
		_spec.SetField(emailtemplate.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := etc.mutation.Name(); ok {
		_spec.SetField(emailtemplate.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := etc.mutation.Subject(); ok {
		_spec.SetField(emailtemplate.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if value, ok := etc.mutation.GetType(); ok {
		_spec.SetField(emailtemplate.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := etc.mutation.HTMLContent(); ok {
		_spec.SetField(emailtemplate.FieldHTMLContent, field.TypeString, value)
		_node.HTMLContent = value
	}
	if value, ok := etc.mutation.TextContent(); ok {
		_spec.SetField(emailtemplate.FieldTextContent, field.TypeString, value)
		_node.TextContent = value
	}
	if value, ok := etc.mutation.OrganizationID(); ok {
		_spec.SetField(emailtemplate.FieldOrganizationID, field.TypeString, value)
		_node.OrganizationID = value
	}
	if value, ok := etc.mutation.Active(); ok {
		_spec.SetField(emailtemplate.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := etc.mutation.System(); ok {
		_spec.SetField(emailtemplate.FieldSystem, field.TypeBool, value)
		_node.System = value
	}
	if value, ok := etc.mutation.Locale(); ok {
		_spec.SetField(emailtemplate.FieldLocale, field.TypeString, value)
		_node.Locale = value
	}
	if value, ok := etc.mutation.Metadata(); ok {
		_spec.SetField(emailtemplate.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	return _node, _spec
}

// EmailTemplateCreateBulk is the builder for creating many EmailTemplate entities in bulk.
type EmailTemplateCreateBulk struct {
	config
	err      error
	builders []*EmailTemplateCreate
}

// Save creates the EmailTemplate entities in the database.
func (etcb *EmailTemplateCreateBulk) Save(ctx context.Context) ([]*EmailTemplate, error) {
	if etcb.err != nil {
		return nil, etcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(etcb.builders))
	nodes := make([]*EmailTemplate, len(etcb.builders))
	mutators := make([]Mutator, len(etcb.builders))
	for i := range etcb.builders {
		func(i int, root context.Context) {
			builder := etcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailTemplateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, etcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, etcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, etcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (etcb *EmailTemplateCreateBulk) SaveX(ctx context.Context) []*EmailTemplate {
	v, err := etcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etcb *EmailTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := etcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etcb *EmailTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := etcb.Exec(ctx); err != nil {
		panic(err)
	}
}
