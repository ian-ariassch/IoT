// Code generated by ent, DO NOT EDIT.

package ent

import (
	"minimal/ent/schema"
	"minimal/ent/water"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	waterFields := schema.Water{}.Fields()
	_ = waterFields
	// waterDescCreatedAt is the schema descriptor for created_at field.
	waterDescCreatedAt := waterFields[2].Descriptor()
	// water.DefaultCreatedAt holds the default value on creation for the created_at field.
	water.DefaultCreatedAt = waterDescCreatedAt.Default.(func() time.Time)
}
