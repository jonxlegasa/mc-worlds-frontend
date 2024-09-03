// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/sentemail"
)

// SentEmailDelete is the builder for deleting a SentEmail entity.
type SentEmailDelete struct {
	config
	hooks    []Hook
	mutation *SentEmailMutation
}

// Where appends a list predicates to the SentEmailDelete builder.
func (sed *SentEmailDelete) Where(ps ...predicate.SentEmail) *SentEmailDelete {
	sed.mutation.Where(ps...)
	return sed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sed *SentEmailDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sed.sqlExec, sed.mutation, sed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sed *SentEmailDelete) ExecX(ctx context.Context) int {
	n, err := sed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sed *SentEmailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sentemail.Table, sqlgraph.NewFieldSpec(sentemail.FieldID, field.TypeInt))
	if ps := sed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sed.mutation.done = true
	return affected, err
}

// SentEmailDeleteOne is the builder for deleting a single SentEmail entity.
type SentEmailDeleteOne struct {
	sed *SentEmailDelete
}

// Where appends a list predicates to the SentEmailDelete builder.
func (sedo *SentEmailDeleteOne) Where(ps ...predicate.SentEmail) *SentEmailDeleteOne {
	sedo.sed.mutation.Where(ps...)
	return sedo
}

// Exec executes the deletion query.
func (sedo *SentEmailDeleteOne) Exec(ctx context.Context) error {
	n, err := sedo.sed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sentemail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sedo *SentEmailDeleteOne) ExecX(ctx context.Context) {
	if err := sedo.Exec(ctx); err != nil {
		panic(err)
	}
}
