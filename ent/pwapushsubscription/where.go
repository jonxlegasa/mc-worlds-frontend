// Code generated by ent, DO NOT EDIT.

package pwapushsubscription

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldUpdatedAt, v))
}

// Endpoint applies equality check predicate on the "endpoint" field. It's identical to EndpointEQ.
func Endpoint(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldEndpoint, v))
}

// P256dh applies equality check predicate on the "p256dh" field. It's identical to P256dhEQ.
func P256dh(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldP256dh, v))
}

// Auth applies equality check predicate on the "auth" field. It's identical to AuthEQ.
func Auth(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldAuth, v))
}

// ProfileID applies equality check predicate on the "profile_id" field. It's identical to ProfileIDEQ.
func ProfileID(v int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldProfileID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLTE(FieldUpdatedAt, v))
}

// EndpointEQ applies the EQ predicate on the "endpoint" field.
func EndpointEQ(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldEndpoint, v))
}

// EndpointNEQ applies the NEQ predicate on the "endpoint" field.
func EndpointNEQ(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNEQ(FieldEndpoint, v))
}

// EndpointIn applies the In predicate on the "endpoint" field.
func EndpointIn(vs ...string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldIn(FieldEndpoint, vs...))
}

// EndpointNotIn applies the NotIn predicate on the "endpoint" field.
func EndpointNotIn(vs ...string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNotIn(FieldEndpoint, vs...))
}

// EndpointGT applies the GT predicate on the "endpoint" field.
func EndpointGT(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGT(FieldEndpoint, v))
}

// EndpointGTE applies the GTE predicate on the "endpoint" field.
func EndpointGTE(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGTE(FieldEndpoint, v))
}

// EndpointLT applies the LT predicate on the "endpoint" field.
func EndpointLT(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLT(FieldEndpoint, v))
}

// EndpointLTE applies the LTE predicate on the "endpoint" field.
func EndpointLTE(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLTE(FieldEndpoint, v))
}

// EndpointContains applies the Contains predicate on the "endpoint" field.
func EndpointContains(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldContains(FieldEndpoint, v))
}

// EndpointHasPrefix applies the HasPrefix predicate on the "endpoint" field.
func EndpointHasPrefix(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldHasPrefix(FieldEndpoint, v))
}

// EndpointHasSuffix applies the HasSuffix predicate on the "endpoint" field.
func EndpointHasSuffix(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldHasSuffix(FieldEndpoint, v))
}

// EndpointEqualFold applies the EqualFold predicate on the "endpoint" field.
func EndpointEqualFold(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEqualFold(FieldEndpoint, v))
}

// EndpointContainsFold applies the ContainsFold predicate on the "endpoint" field.
func EndpointContainsFold(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldContainsFold(FieldEndpoint, v))
}

// P256dhEQ applies the EQ predicate on the "p256dh" field.
func P256dhEQ(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldP256dh, v))
}

// P256dhNEQ applies the NEQ predicate on the "p256dh" field.
func P256dhNEQ(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNEQ(FieldP256dh, v))
}

// P256dhIn applies the In predicate on the "p256dh" field.
func P256dhIn(vs ...string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldIn(FieldP256dh, vs...))
}

// P256dhNotIn applies the NotIn predicate on the "p256dh" field.
func P256dhNotIn(vs ...string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNotIn(FieldP256dh, vs...))
}

// P256dhGT applies the GT predicate on the "p256dh" field.
func P256dhGT(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGT(FieldP256dh, v))
}

// P256dhGTE applies the GTE predicate on the "p256dh" field.
func P256dhGTE(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGTE(FieldP256dh, v))
}

// P256dhLT applies the LT predicate on the "p256dh" field.
func P256dhLT(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLT(FieldP256dh, v))
}

// P256dhLTE applies the LTE predicate on the "p256dh" field.
func P256dhLTE(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLTE(FieldP256dh, v))
}

// P256dhContains applies the Contains predicate on the "p256dh" field.
func P256dhContains(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldContains(FieldP256dh, v))
}

// P256dhHasPrefix applies the HasPrefix predicate on the "p256dh" field.
func P256dhHasPrefix(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldHasPrefix(FieldP256dh, v))
}

// P256dhHasSuffix applies the HasSuffix predicate on the "p256dh" field.
func P256dhHasSuffix(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldHasSuffix(FieldP256dh, v))
}

// P256dhEqualFold applies the EqualFold predicate on the "p256dh" field.
func P256dhEqualFold(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEqualFold(FieldP256dh, v))
}

// P256dhContainsFold applies the ContainsFold predicate on the "p256dh" field.
func P256dhContainsFold(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldContainsFold(FieldP256dh, v))
}

// AuthEQ applies the EQ predicate on the "auth" field.
func AuthEQ(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldAuth, v))
}

// AuthNEQ applies the NEQ predicate on the "auth" field.
func AuthNEQ(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNEQ(FieldAuth, v))
}

// AuthIn applies the In predicate on the "auth" field.
func AuthIn(vs ...string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldIn(FieldAuth, vs...))
}

// AuthNotIn applies the NotIn predicate on the "auth" field.
func AuthNotIn(vs ...string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNotIn(FieldAuth, vs...))
}

// AuthGT applies the GT predicate on the "auth" field.
func AuthGT(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGT(FieldAuth, v))
}

// AuthGTE applies the GTE predicate on the "auth" field.
func AuthGTE(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldGTE(FieldAuth, v))
}

// AuthLT applies the LT predicate on the "auth" field.
func AuthLT(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLT(FieldAuth, v))
}

// AuthLTE applies the LTE predicate on the "auth" field.
func AuthLTE(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldLTE(FieldAuth, v))
}

// AuthContains applies the Contains predicate on the "auth" field.
func AuthContains(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldContains(FieldAuth, v))
}

// AuthHasPrefix applies the HasPrefix predicate on the "auth" field.
func AuthHasPrefix(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldHasPrefix(FieldAuth, v))
}

// AuthHasSuffix applies the HasSuffix predicate on the "auth" field.
func AuthHasSuffix(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldHasSuffix(FieldAuth, v))
}

// AuthEqualFold applies the EqualFold predicate on the "auth" field.
func AuthEqualFold(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEqualFold(FieldAuth, v))
}

// AuthContainsFold applies the ContainsFold predicate on the "auth" field.
func AuthContainsFold(v string) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldContainsFold(FieldAuth, v))
}

// ProfileIDEQ applies the EQ predicate on the "profile_id" field.
func ProfileIDEQ(v int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldEQ(FieldProfileID, v))
}

// ProfileIDNEQ applies the NEQ predicate on the "profile_id" field.
func ProfileIDNEQ(v int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNEQ(FieldProfileID, v))
}

// ProfileIDIn applies the In predicate on the "profile_id" field.
func ProfileIDIn(vs ...int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldIn(FieldProfileID, vs...))
}

// ProfileIDNotIn applies the NotIn predicate on the "profile_id" field.
func ProfileIDNotIn(vs ...int) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.FieldNotIn(FieldProfileID, vs...))
}

// HasProfile applies the HasEdge predicate on the "profile" edge.
func HasProfile() predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfileWith applies the HasEdge predicate on the "profile" edge with a given conditions (other predicates).
func HasProfileWith(preds ...predicate.Profile) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(func(s *sql.Selector) {
		step := newProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PwaPushSubscription) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PwaPushSubscription) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PwaPushSubscription) predicate.PwaPushSubscription {
	return predicate.PwaPushSubscription(sql.NotPredicates(p))
}
