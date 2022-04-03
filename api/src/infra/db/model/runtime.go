// Code generated by entc, DO NOT EDIT.

package model

import (
	"time"

	"github.com/Southclaws/storyden/api/src/infra/db/model/category"
	"github.com/Southclaws/storyden/api/src/infra/db/model/post"
	"github.com/Southclaws/storyden/api/src/infra/db/model/user"
	"github.com/Southclaws/storyden/api/src/infra/db/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescDescription is the schema descriptor for description field.
	categoryDescDescription := categoryFields[2].Descriptor()
	// category.DefaultDescription holds the default value on creation for the description field.
	category.DefaultDescription = categoryDescDescription.Default.(string)
	// categoryDescColour is the schema descriptor for colour field.
	categoryDescColour := categoryFields[3].Descriptor()
	// category.DefaultColour holds the default value on creation for the colour field.
	category.DefaultColour = categoryDescColour.Default.(string)
	// categoryDescSort is the schema descriptor for sort field.
	categoryDescSort := categoryFields[4].Descriptor()
	// category.DefaultSort holds the default value on creation for the sort field.
	category.DefaultSort = categoryDescSort.Default.(int)
	// categoryDescAdmin is the schema descriptor for admin field.
	categoryDescAdmin := categoryFields[5].Descriptor()
	// category.DefaultAdmin holds the default value on creation for the admin field.
	category.DefaultAdmin = categoryDescAdmin.Default.(bool)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescPinned is the schema descriptor for pinned field.
	postDescPinned := postFields[6].Descriptor()
	// post.DefaultPinned holds the default value on creation for the pinned field.
	post.DefaultPinned = postDescPinned.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[4].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescUpdatedAt is the schema descriptor for updatedAt field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
