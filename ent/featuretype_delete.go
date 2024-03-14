// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"golang-intern/assignment-2/ent/featuretype"
	"golang-intern/assignment-2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeatureTypeDelete is the builder for deleting a FeatureType entity.
type FeatureTypeDelete struct {
	config
	hooks    []Hook
	mutation *FeatureTypeMutation
}

// Where appends a list predicates to the FeatureTypeDelete builder.
func (ftd *FeatureTypeDelete) Where(ps ...predicate.FeatureType) *FeatureTypeDelete {
	ftd.mutation.Where(ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FeatureTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ftd.sqlExec, ftd.mutation, ftd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FeatureTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FeatureTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(featuretype.Table, sqlgraph.NewFieldSpec(featuretype.FieldID, field.TypeInt))
	if ps := ftd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ftd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ftd.mutation.done = true
	return affected, err
}

// FeatureTypeDeleteOne is the builder for deleting a single FeatureType entity.
type FeatureTypeDeleteOne struct {
	ftd *FeatureTypeDelete
}

// Where appends a list predicates to the FeatureTypeDelete builder.
func (ftdo *FeatureTypeDeleteOne) Where(ps ...predicate.FeatureType) *FeatureTypeDeleteOne {
	ftdo.ftd.mutation.Where(ps...)
	return ftdo
}

// Exec executes the deletion query.
func (ftdo *FeatureTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{featuretype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FeatureTypeDeleteOne) ExecX(ctx context.Context) {
	if err := ftdo.Exec(ctx); err != nil {
		panic(err)
	}
}
