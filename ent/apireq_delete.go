// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"golang-intern/assignment-2/ent/apireq"
	"golang-intern/assignment-2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApireqDelete is the builder for deleting a Apireq entity.
type ApireqDelete struct {
	config
	hooks    []Hook
	mutation *ApireqMutation
}

// Where appends a list predicates to the ApireqDelete builder.
func (ad *ApireqDelete) Where(ps ...predicate.Apireq) *ApireqDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *ApireqDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *ApireqDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *ApireqDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(apireq.Table, sqlgraph.NewFieldSpec(apireq.FieldID, field.TypeInt))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// ApireqDeleteOne is the builder for deleting a single Apireq entity.
type ApireqDeleteOne struct {
	ad *ApireqDelete
}

// Where appends a list predicates to the ApireqDelete builder.
func (ado *ApireqDeleteOne) Where(ps ...predicate.Apireq) *ApireqDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *ApireqDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{apireq.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *ApireqDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}