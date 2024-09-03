// Code generated by ent, DO NOT EDIT.

package pwapushsubscription

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the pwapushsubscription type in the database.
	Label = "pwa_push_subscription"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEndpoint holds the string denoting the endpoint field in the database.
	FieldEndpoint = "endpoint"
	// FieldP256dh holds the string denoting the p256dh field in the database.
	FieldP256dh = "p256dh"
	// FieldAuth holds the string denoting the auth field in the database.
	FieldAuth = "auth"
	// FieldProfileID holds the string denoting the profile_id field in the database.
	FieldProfileID = "profile_id"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// Table holds the table name of the pwapushsubscription in the database.
	Table = "pwa_push_subscriptions"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "pwa_push_subscriptions"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "profile_id"
)

// Columns holds all SQL columns for pwapushsubscription fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEndpoint,
	FieldP256dh,
	FieldAuth,
	FieldProfileID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// EndpointValidator is a validator for the "endpoint" field. It is called by the builders before save.
	EndpointValidator func(string) error
	// P256dhValidator is a validator for the "p256dh" field. It is called by the builders before save.
	P256dhValidator func(string) error
	// AuthValidator is a validator for the "auth" field. It is called by the builders before save.
	AuthValidator func(string) error
)

// OrderOption defines the ordering options for the PwaPushSubscription queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByEndpoint orders the results by the endpoint field.
func ByEndpoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndpoint, opts...).ToFunc()
}

// ByP256dh orders the results by the p256dh field.
func ByP256dh(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldP256dh, opts...).ToFunc()
}

// ByAuth orders the results by the auth field.
func ByAuth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuth, opts...).ToFunc()
}

// ByProfileID orders the results by the profile_id field.
func ByProfileID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfileID, opts...).ToFunc()
}

// ByProfileField orders the results by profile field.
func ByProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), sql.OrderByField(field, opts...))
	}
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
	)
}
