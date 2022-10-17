package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ToDo holds the schema definition for the ToDo entity.
type ToDo struct {
	ent.Schema
}

// Fields of the ToDo.
func (ToDo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("todo_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("todo_context").
			Default("null"),
		field.Bool("completion").
			Default(false),
		field.Time("completed_at").
			Default("null"), // 완료하지 않으면 표시하지 않음.
		field.String("created_at").
			Default("null"), // Year-Mouth-Day
	}
}

// Edges of the ToDo.
func (ToDo) Edges() []ent.Edge {
	return nil
}