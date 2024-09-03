// Code generated by ent, DO NOT EDIT.

package invitation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the invitation type in the database.
	Label = "invitation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldInviteeName holds the string denoting the invitee_name field in the database.
	FieldInviteeName = "invitee_name"
	// FieldConfirmationCode holds the string denoting the confirmation_code field in the database.
	FieldConfirmationCode = "confirmation_code"
	// EdgeInviter holds the string denoting the inviter edge name in mutations.
	EdgeInviter = "inviter"
	// Table holds the table name of the invitation in the database.
	Table = "invitations"
	// InviterTable is the table that holds the inviter relation/edge.
	InviterTable = "invitations"
	// InviterInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	InviterInverseTable = "profiles"
	// InviterColumn is the table column denoting the inviter relation/edge.
	InviterColumn = "profile_invitations"
)

// Columns holds all SQL columns for invitation fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldInviteeName,
	FieldConfirmationCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "invitations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profile_invitations",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// InviteeNameValidator is a validator for the "invitee_name" field. It is called by the builders before save.
	InviteeNameValidator func(string) error
	// ConfirmationCodeValidator is a validator for the "confirmation_code" field. It is called by the builders before save.
	ConfirmationCodeValidator func(string) error
)

// OrderOption defines the ordering options for the Invitation queries.
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

// ByInviteeName orders the results by the invitee_name field.
func ByInviteeName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInviteeName, opts...).ToFunc()
}

// ByConfirmationCode orders the results by the confirmation_code field.
func ByConfirmationCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfirmationCode, opts...).ToFunc()
}

// ByInviterField orders the results by inviter field.
func ByInviterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInviterStep(), sql.OrderByField(field, opts...))
	}
}
func newInviterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InviterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InviterTable, InviterColumn),
	)
}
