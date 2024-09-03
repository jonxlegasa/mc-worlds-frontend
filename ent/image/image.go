// Code generated by ent, DO NOT EDIT.

package image

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeSizes holds the string denoting the sizes edge name in mutations.
	EdgeSizes = "sizes"
	// Table holds the table name of the image in the database.
	Table = "images"
	// SizesTable is the table that holds the sizes relation/edge.
	SizesTable = "image_sizes"
	// SizesInverseTable is the table name for the ImageSize entity.
	// It exists in this package in order to avoid circular dependency with the "imagesize" package.
	SizesInverseTable = "image_sizes"
	// SizesColumn is the table column denoting the sizes relation/edge.
	SizesColumn = "image_sizes"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profile_photos",
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
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeProfilePhoto   Type = "profile_photo"
	TypeProfileGallery Type = "profile_gallery"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeProfilePhoto, TypeProfileGallery:
		return nil
	default:
		return fmt.Errorf("image: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Image queries.
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

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// BySizesCount orders the results by sizes count.
func BySizesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSizesStep(), opts...)
	}
}

// BySizes orders the results by sizes terms.
func BySizes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSizesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSizesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SizesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SizesTable, SizesColumn),
	)
}
