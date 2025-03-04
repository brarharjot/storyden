// Code generated by ent, DO NOT EDIT.

package category

import (
	"time"

	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldColour holds the string denoting the colour field in the database.
	FieldColour = "colour"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "category_id"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldColour,
	FieldSort,
	FieldAdmin,
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
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultColour holds the default value on creation for the "colour" field.
	DefaultColour string
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int
	// DefaultAdmin holds the default value on creation for the "admin" field.
	DefaultAdmin bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
