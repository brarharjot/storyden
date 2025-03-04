// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldFirst holds the string denoting the first field in the database.
	FieldFirst = "first"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldPinned holds the string denoting the pinned field in the database.
	FieldPinned = "pinned"
	// FieldRootPostID holds the string denoting the root_post_id field in the database.
	FieldRootPostID = "root_post_id"
	// FieldReplyToPostID holds the string denoting the reply_to_post_id field in the database.
	FieldReplyToPostID = "reply_to_post_id"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldShort holds the string denoting the short field in the database.
	FieldShort = "short"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeRoot holds the string denoting the root edge name in mutations.
	EdgeRoot = "root"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeReplyTo holds the string denoting the replyto edge name in mutations.
	EdgeReplyTo = "replyTo"
	// EdgeReplies holds the string denoting the replies edge name in mutations.
	EdgeReplies = "replies"
	// EdgeReacts holds the string denoting the reacts edge name in mutations.
	EdgeReacts = "reacts"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "posts"
	// AuthorInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AuthorInverseTable = "accounts"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "account_posts"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "posts"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_id"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "tag_posts"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// RootTable is the table that holds the root relation/edge.
	RootTable = "posts"
	// RootColumn is the table column denoting the root relation/edge.
	RootColumn = "root_post_id"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "root_post_id"
	// ReplyToTable is the table that holds the replyTo relation/edge.
	ReplyToTable = "posts"
	// ReplyToColumn is the table column denoting the replyTo relation/edge.
	ReplyToColumn = "reply_to_post_id"
	// RepliesTable is the table that holds the replies relation/edge.
	RepliesTable = "posts"
	// RepliesColumn is the table column denoting the replies relation/edge.
	RepliesColumn = "reply_to_post_id"
	// ReactsTable is the table that holds the reacts relation/edge.
	ReactsTable = "reacts"
	// ReactsInverseTable is the table name for the React entity.
	// It exists in this package in order to avoid circular dependency with the "react" package.
	ReactsInverseTable = "reacts"
	// ReactsColumn is the table column denoting the reacts relation/edge.
	ReactsColumn = "post_reacts"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldFirst,
	FieldTitle,
	FieldSlug,
	FieldPinned,
	FieldRootPostID,
	FieldReplyToPostID,
	FieldBody,
	FieldShort,
	FieldMetadata,
	FieldCategoryID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_posts",
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"tag_id", "post_id"}
)

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
	// DefaultPinned holds the default value on creation for the "pinned" field.
	DefaultPinned bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
