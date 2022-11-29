package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Water holds the schema definition for the Water entity.
type Water struct {
	ent.Schema
}

// Fields of the Water.
func (Water) Fields() []ent.Field {
	return []ent.Field{
		field.Float("liters"),
		field.String("topic"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Water.
func (Water) Edges() []ent.Edge {
	return nil
}
