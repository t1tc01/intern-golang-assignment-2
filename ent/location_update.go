// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golang-intern/assignment-2/ent/geometry"
	"golang-intern/assignment-2/ent/location"
	"golang-intern/assignment-2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LocationUpdate is the builder for updating Location entities.
type LocationUpdate struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// Where appends a list predicates to the LocationUpdate builder.
func (lu *LocationUpdate) Where(ps ...predicate.Location) *LocationUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetPlace sets the "place" field.
func (lu *LocationUpdate) SetPlace(s string) *LocationUpdate {
	lu.mutation.SetPlace(s)
	return lu
}

// SetNillablePlace sets the "place" field if the given value is not nil.
func (lu *LocationUpdate) SetNillablePlace(s *string) *LocationUpdate {
	if s != nil {
		lu.SetPlace(*s)
	}
	return lu
}

// ClearPlace clears the value of the "place" field.
func (lu *LocationUpdate) ClearPlace() *LocationUpdate {
	lu.mutation.ClearPlace()
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LocationUpdate) SetCreatedAt(t time.Time) *LocationUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LocationUpdate) SetNillableCreatedAt(t *time.Time) *LocationUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LocationUpdate) SetUpdatedAt(t time.Time) *LocationUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lu *LocationUpdate) SetNillableUpdatedAt(t *time.Time) *LocationUpdate {
	if t != nil {
		lu.SetUpdatedAt(*t)
	}
	return lu
}

// SetDeletedAt sets the "deleted_at" field.
func (lu *LocationUpdate) SetDeletedAt(t time.Time) *LocationUpdate {
	lu.mutation.SetDeletedAt(t)
	return lu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lu *LocationUpdate) SetNillableDeletedAt(t *time.Time) *LocationUpdate {
	if t != nil {
		lu.SetDeletedAt(*t)
	}
	return lu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (lu *LocationUpdate) ClearDeletedAt() *LocationUpdate {
	lu.mutation.ClearDeletedAt()
	return lu
}

// AddGeometryIDs adds the "geometries" edge to the Geometry entity by IDs.
func (lu *LocationUpdate) AddGeometryIDs(ids ...int) *LocationUpdate {
	lu.mutation.AddGeometryIDs(ids...)
	return lu
}

// AddGeometries adds the "geometries" edges to the Geometry entity.
func (lu *LocationUpdate) AddGeometries(g ...*Geometry) *LocationUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return lu.AddGeometryIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lu *LocationUpdate) Mutation() *LocationMutation {
	return lu.mutation
}

// ClearGeometries clears all "geometries" edges to the Geometry entity.
func (lu *LocationUpdate) ClearGeometries() *LocationUpdate {
	lu.mutation.ClearGeometries()
	return lu
}

// RemoveGeometryIDs removes the "geometries" edge to Geometry entities by IDs.
func (lu *LocationUpdate) RemoveGeometryIDs(ids ...int) *LocationUpdate {
	lu.mutation.RemoveGeometryIDs(ids...)
	return lu
}

// RemoveGeometries removes "geometries" edges to Geometry entities.
func (lu *LocationUpdate) RemoveGeometries(g ...*Geometry) *LocationUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return lu.RemoveGeometryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LocationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LocationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LocationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LocationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Place(); ok {
		_spec.SetField(location.FieldPlace, field.TypeString, value)
	}
	if lu.mutation.PlaceCleared() {
		_spec.ClearField(location.FieldPlace, field.TypeString)
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(location.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(location.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.DeletedAt(); ok {
		_spec.SetField(location.FieldDeletedAt, field.TypeTime, value)
	}
	if lu.mutation.DeletedAtCleared() {
		_spec.ClearField(location.FieldDeletedAt, field.TypeTime)
	}
	if lu.mutation.GeometriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.GeometriesTable,
			Columns: []string{location.GeometriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedGeometriesIDs(); len(nodes) > 0 && !lu.mutation.GeometriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.GeometriesTable,
			Columns: []string{location.GeometriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.GeometriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.GeometriesTable,
			Columns: []string{location.GeometriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LocationUpdateOne is the builder for updating a single Location entity.
type LocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LocationMutation
}

// SetPlace sets the "place" field.
func (luo *LocationUpdateOne) SetPlace(s string) *LocationUpdateOne {
	luo.mutation.SetPlace(s)
	return luo
}

// SetNillablePlace sets the "place" field if the given value is not nil.
func (luo *LocationUpdateOne) SetNillablePlace(s *string) *LocationUpdateOne {
	if s != nil {
		luo.SetPlace(*s)
	}
	return luo
}

// ClearPlace clears the value of the "place" field.
func (luo *LocationUpdateOne) ClearPlace() *LocationUpdateOne {
	luo.mutation.ClearPlace()
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *LocationUpdateOne) SetCreatedAt(t time.Time) *LocationUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableCreatedAt(t *time.Time) *LocationUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LocationUpdateOne) SetUpdatedAt(t time.Time) *LocationUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableUpdatedAt(t *time.Time) *LocationUpdateOne {
	if t != nil {
		luo.SetUpdatedAt(*t)
	}
	return luo
}

// SetDeletedAt sets the "deleted_at" field.
func (luo *LocationUpdateOne) SetDeletedAt(t time.Time) *LocationUpdateOne {
	luo.mutation.SetDeletedAt(t)
	return luo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableDeletedAt(t *time.Time) *LocationUpdateOne {
	if t != nil {
		luo.SetDeletedAt(*t)
	}
	return luo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (luo *LocationUpdateOne) ClearDeletedAt() *LocationUpdateOne {
	luo.mutation.ClearDeletedAt()
	return luo
}

// AddGeometryIDs adds the "geometries" edge to the Geometry entity by IDs.
func (luo *LocationUpdateOne) AddGeometryIDs(ids ...int) *LocationUpdateOne {
	luo.mutation.AddGeometryIDs(ids...)
	return luo
}

// AddGeometries adds the "geometries" edges to the Geometry entity.
func (luo *LocationUpdateOne) AddGeometries(g ...*Geometry) *LocationUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return luo.AddGeometryIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (luo *LocationUpdateOne) Mutation() *LocationMutation {
	return luo.mutation
}

// ClearGeometries clears all "geometries" edges to the Geometry entity.
func (luo *LocationUpdateOne) ClearGeometries() *LocationUpdateOne {
	luo.mutation.ClearGeometries()
	return luo
}

// RemoveGeometryIDs removes the "geometries" edge to Geometry entities by IDs.
func (luo *LocationUpdateOne) RemoveGeometryIDs(ids ...int) *LocationUpdateOne {
	luo.mutation.RemoveGeometryIDs(ids...)
	return luo
}

// RemoveGeometries removes "geometries" edges to Geometry entities.
func (luo *LocationUpdateOne) RemoveGeometries(g ...*Geometry) *LocationUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return luo.RemoveGeometryIDs(ids...)
}

// Where appends a list predicates to the LocationUpdate builder.
func (luo *LocationUpdateOne) Where(ps ...predicate.Location) *LocationUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LocationUpdateOne) Select(field string, fields ...string) *LocationUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Location entity.
func (luo *LocationUpdateOne) Save(ctx context.Context) (*Location, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LocationUpdateOne) SaveX(ctx context.Context) *Location {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LocationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LocationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LocationUpdateOne) sqlSave(ctx context.Context) (_node *Location, err error) {
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Location.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for _, f := range fields {
			if !location.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Place(); ok {
		_spec.SetField(location.FieldPlace, field.TypeString, value)
	}
	if luo.mutation.PlaceCleared() {
		_spec.ClearField(location.FieldPlace, field.TypeString)
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(location.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(location.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.DeletedAt(); ok {
		_spec.SetField(location.FieldDeletedAt, field.TypeTime, value)
	}
	if luo.mutation.DeletedAtCleared() {
		_spec.ClearField(location.FieldDeletedAt, field.TypeTime)
	}
	if luo.mutation.GeometriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.GeometriesTable,
			Columns: []string{location.GeometriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedGeometriesIDs(); len(nodes) > 0 && !luo.mutation.GeometriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.GeometriesTable,
			Columns: []string{location.GeometriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.GeometriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.GeometriesTable,
			Columns: []string{location.GeometriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Location{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
