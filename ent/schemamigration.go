// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"golang-intern/assignment-2/ent/schemamigration"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SchemaMigration is the model entity for the SchemaMigration schema.
type SchemaMigration struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Dirty holds the value of the "dirty" field.
	Dirty        bool `json:"dirty,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SchemaMigration) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case schemamigration.FieldDirty:
			values[i] = new(sql.NullBool)
		case schemamigration.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SchemaMigration fields.
func (sm *SchemaMigration) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case schemamigration.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sm.ID = int(value.Int64)
		case schemamigration.FieldDirty:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field dirty", values[i])
			} else if value.Valid {
				sm.Dirty = value.Bool
			}
		default:
			sm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SchemaMigration.
// This includes values selected through modifiers, order, etc.
func (sm *SchemaMigration) Value(name string) (ent.Value, error) {
	return sm.selectValues.Get(name)
}

// Update returns a builder for updating this SchemaMigration.
// Note that you need to call SchemaMigration.Unwrap() before calling this method if this SchemaMigration
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *SchemaMigration) Update() *SchemaMigrationUpdateOne {
	return NewSchemaMigrationClient(sm.config).UpdateOne(sm)
}

// Unwrap unwraps the SchemaMigration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *SchemaMigration) Unwrap() *SchemaMigration {
	_tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("ent: SchemaMigration is not a transactional entity")
	}
	sm.config.driver = _tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *SchemaMigration) String() string {
	var builder strings.Builder
	builder.WriteString("SchemaMigration(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sm.ID))
	builder.WriteString("dirty=")
	builder.WriteString(fmt.Sprintf("%v", sm.Dirty))
	builder.WriteByte(')')
	return builder.String()
}

// SchemaMigrations is a parsable slice of SchemaMigration.
type SchemaMigrations []*SchemaMigration
