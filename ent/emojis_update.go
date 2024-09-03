// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/emojis"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// EmojisUpdate is the builder for updating Emojis entities.
type EmojisUpdate struct {
	config
	hooks    []Hook
	mutation *EmojisMutation
}

// Where appends a list predicates to the EmojisUpdate builder.
func (eu *EmojisUpdate) Where(ps ...predicate.Emojis) *EmojisUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUnifiedCode sets the "unified_code" field.
func (eu *EmojisUpdate) SetUnifiedCode(s string) *EmojisUpdate {
	eu.mutation.SetUnifiedCode(s)
	return eu
}

// SetNillableUnifiedCode sets the "unified_code" field if the given value is not nil.
func (eu *EmojisUpdate) SetNillableUnifiedCode(s *string) *EmojisUpdate {
	if s != nil {
		eu.SetUnifiedCode(*s)
	}
	return eu
}

// SetShortcode sets the "shortcode" field.
func (eu *EmojisUpdate) SetShortcode(s string) *EmojisUpdate {
	eu.mutation.SetShortcode(s)
	return eu
}

// SetNillableShortcode sets the "shortcode" field if the given value is not nil.
func (eu *EmojisUpdate) SetNillableShortcode(s *string) *EmojisUpdate {
	if s != nil {
		eu.SetShortcode(*s)
	}
	return eu
}

// Mutation returns the EmojisMutation object of the builder.
func (eu *EmojisUpdate) Mutation() *EmojisMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmojisUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmojisUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmojisUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmojisUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EmojisUpdate) check() error {
	if v, ok := eu.mutation.UnifiedCode(); ok {
		if err := emojis.UnifiedCodeValidator(v); err != nil {
			return &ValidationError{Name: "unified_code", err: fmt.Errorf(`ent: validator failed for field "Emojis.unified_code": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Shortcode(); ok {
		if err := emojis.ShortcodeValidator(v); err != nil {
			return &ValidationError{Name: "shortcode", err: fmt.Errorf(`ent: validator failed for field "Emojis.shortcode": %w`, err)}
		}
	}
	return nil
}

func (eu *EmojisUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(emojis.Table, emojis.Columns, sqlgraph.NewFieldSpec(emojis.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UnifiedCode(); ok {
		_spec.SetField(emojis.FieldUnifiedCode, field.TypeString, value)
	}
	if value, ok := eu.mutation.Shortcode(); ok {
		_spec.SetField(emojis.FieldShortcode, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emojis.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmojisUpdateOne is the builder for updating a single Emojis entity.
type EmojisUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmojisMutation
}

// SetUnifiedCode sets the "unified_code" field.
func (euo *EmojisUpdateOne) SetUnifiedCode(s string) *EmojisUpdateOne {
	euo.mutation.SetUnifiedCode(s)
	return euo
}

// SetNillableUnifiedCode sets the "unified_code" field if the given value is not nil.
func (euo *EmojisUpdateOne) SetNillableUnifiedCode(s *string) *EmojisUpdateOne {
	if s != nil {
		euo.SetUnifiedCode(*s)
	}
	return euo
}

// SetShortcode sets the "shortcode" field.
func (euo *EmojisUpdateOne) SetShortcode(s string) *EmojisUpdateOne {
	euo.mutation.SetShortcode(s)
	return euo
}

// SetNillableShortcode sets the "shortcode" field if the given value is not nil.
func (euo *EmojisUpdateOne) SetNillableShortcode(s *string) *EmojisUpdateOne {
	if s != nil {
		euo.SetShortcode(*s)
	}
	return euo
}

// Mutation returns the EmojisMutation object of the builder.
func (euo *EmojisUpdateOne) Mutation() *EmojisMutation {
	return euo.mutation
}

// Where appends a list predicates to the EmojisUpdate builder.
func (euo *EmojisUpdateOne) Where(ps ...predicate.Emojis) *EmojisUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmojisUpdateOne) Select(field string, fields ...string) *EmojisUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Emojis entity.
func (euo *EmojisUpdateOne) Save(ctx context.Context) (*Emojis, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmojisUpdateOne) SaveX(ctx context.Context) *Emojis {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmojisUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmojisUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EmojisUpdateOne) check() error {
	if v, ok := euo.mutation.UnifiedCode(); ok {
		if err := emojis.UnifiedCodeValidator(v); err != nil {
			return &ValidationError{Name: "unified_code", err: fmt.Errorf(`ent: validator failed for field "Emojis.unified_code": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Shortcode(); ok {
		if err := emojis.ShortcodeValidator(v); err != nil {
			return &ValidationError{Name: "shortcode", err: fmt.Errorf(`ent: validator failed for field "Emojis.shortcode": %w`, err)}
		}
	}
	return nil
}

func (euo *EmojisUpdateOne) sqlSave(ctx context.Context) (_node *Emojis, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(emojis.Table, emojis.Columns, sqlgraph.NewFieldSpec(emojis.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Emojis.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emojis.FieldID)
		for _, f := range fields {
			if !emojis.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emojis.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UnifiedCode(); ok {
		_spec.SetField(emojis.FieldUnifiedCode, field.TypeString, value)
	}
	if value, ok := euo.mutation.Shortcode(); ok {
		_spec.SetField(emojis.FieldShortcode, field.TypeString, value)
	}
	_node = &Emojis{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emojis.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
