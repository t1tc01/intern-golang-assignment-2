// Code generated by ent, DO NOT EDIT.

package featuretype

import (
	"golang-intern/assignment-2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLTE(FieldID, id))
}

// FeatType applies equality check predicate on the "feat_type" field. It's identical to FeatTypeEQ.
func FeatType(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldFeatType, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldDeletedAt, v))
}

// FeatTypeEQ applies the EQ predicate on the "feat_type" field.
func FeatTypeEQ(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldFeatType, v))
}

// FeatTypeNEQ applies the NEQ predicate on the "feat_type" field.
func FeatTypeNEQ(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNEQ(FieldFeatType, v))
}

// FeatTypeIn applies the In predicate on the "feat_type" field.
func FeatTypeIn(vs ...string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldIn(FieldFeatType, vs...))
}

// FeatTypeNotIn applies the NotIn predicate on the "feat_type" field.
func FeatTypeNotIn(vs ...string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNotIn(FieldFeatType, vs...))
}

// FeatTypeGT applies the GT predicate on the "feat_type" field.
func FeatTypeGT(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGT(FieldFeatType, v))
}

// FeatTypeGTE applies the GTE predicate on the "feat_type" field.
func FeatTypeGTE(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGTE(FieldFeatType, v))
}

// FeatTypeLT applies the LT predicate on the "feat_type" field.
func FeatTypeLT(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLT(FieldFeatType, v))
}

// FeatTypeLTE applies the LTE predicate on the "feat_type" field.
func FeatTypeLTE(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLTE(FieldFeatType, v))
}

// FeatTypeContains applies the Contains predicate on the "feat_type" field.
func FeatTypeContains(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldContains(FieldFeatType, v))
}

// FeatTypeHasPrefix applies the HasPrefix predicate on the "feat_type" field.
func FeatTypeHasPrefix(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldHasPrefix(FieldFeatType, v))
}

// FeatTypeHasSuffix applies the HasSuffix predicate on the "feat_type" field.
func FeatTypeHasSuffix(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldHasSuffix(FieldFeatType, v))
}

// FeatTypeIsNil applies the IsNil predicate on the "feat_type" field.
func FeatTypeIsNil() predicate.FeatureType {
	return predicate.FeatureType(sql.FieldIsNull(FieldFeatType))
}

// FeatTypeNotNil applies the NotNil predicate on the "feat_type" field.
func FeatTypeNotNil() predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNotNull(FieldFeatType))
}

// FeatTypeEqualFold applies the EqualFold predicate on the "feat_type" field.
func FeatTypeEqualFold(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEqualFold(FieldFeatType, v))
}

// FeatTypeContainsFold applies the ContainsFold predicate on the "feat_type" field.
func FeatTypeContainsFold(v string) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldContainsFold(FieldFeatType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.FeatureType {
	return predicate.FeatureType(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.FeatureType {
	return predicate.FeatureType(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.FeatureType {
	return predicate.FeatureType(sql.FieldNotNull(FieldDeletedAt))
}

// HasFtypeEarthquakes applies the HasEdge predicate on the "ftype_earthquakes" edge.
func HasFtypeEarthquakes() predicate.FeatureType {
	return predicate.FeatureType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FtypeEarthquakesTable, FtypeEarthquakesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFtypeEarthquakesWith applies the HasEdge predicate on the "ftype_earthquakes" edge with a given conditions (other predicates).
func HasFtypeEarthquakesWith(preds ...predicate.FtypeEarthquake) predicate.FeatureType {
	return predicate.FeatureType(func(s *sql.Selector) {
		step := newFtypeEarthquakesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FeatureType) predicate.FeatureType {
	return predicate.FeatureType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FeatureType) predicate.FeatureType {
	return predicate.FeatureType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FeatureType) predicate.FeatureType {
	return predicate.FeatureType(sql.NotPredicates(p))
}
