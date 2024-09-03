// Code generated by ent, DO NOT EDIT.

package imagesize

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the imagesize type in the database.
	Label = "image_size"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// EdgeFile holds the string denoting the file edge name in mutations.
	EdgeFile = "file"
	// EdgeImage holds the string denoting the image edge name in mutations.
	EdgeImage = "image"
	// Table holds the table name of the imagesize in the database.
	Table = "image_sizes"
	// FileTable is the table that holds the file relation/edge.
	FileTable = "image_sizes"
	// FileInverseTable is the table name for the FileStorage entity.
	// It exists in this package in order to avoid circular dependency with the "filestorage" package.
	FileInverseTable = "file_storages"
	// FileColumn is the table column denoting the file relation/edge.
	FileColumn = "image_size_file"
	// ImageTable is the table that holds the image relation/edge.
	ImageTable = "image_sizes"
	// ImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImageInverseTable = "images"
	// ImageColumn is the table column denoting the image relation/edge.
	ImageColumn = "image_sizes"
)

// Columns holds all SQL columns for imagesize fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSize,
	FieldWidth,
	FieldHeight,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "image_sizes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"image_sizes",
	"image_size_file",
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
	// WidthValidator is a validator for the "width" field. It is called by the builders before save.
	WidthValidator func(int) error
	// HeightValidator is a validator for the "height" field. It is called by the builders before save.
	HeightValidator func(int) error
)

// Size defines the type for the "size" enum field.
type Size string

// Size values.
const (
	SizeThumbnail Size = "thumbnail"
	SizePreview   Size = "preview"
	SizeFull      Size = "full"
)

func (s Size) String() string {
	return string(s)
}

// SizeValidator is a validator for the "size" field enum values. It is called by the builders before save.
func SizeValidator(s Size) error {
	switch s {
	case SizeThumbnail, SizePreview, SizeFull:
		return nil
	default:
		return fmt.Errorf("imagesize: invalid enum value for size field: %q", s)
	}
}

// OrderOption defines the ordering options for the ImageSize queries.
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

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByWidth orders the results by the width field.
func ByWidth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWidth, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByFileField orders the results by file field.
func ByFileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFileStep(), sql.OrderByField(field, opts...))
	}
}

// ByImageField orders the results by image field.
func ByImageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImageStep(), sql.OrderByField(field, opts...))
	}
}
func newFileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, FileTable, FileColumn),
	)
}
func newImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ImageTable, ImageColumn),
	)
}
