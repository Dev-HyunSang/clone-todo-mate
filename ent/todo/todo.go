// Code generated by ent, DO NOT EDIT.

package todo

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "to_do"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTodoUUID holds the string denoting the todo_uuid field in the database.
	FieldTodoUUID = "todo_uuid"
	// FieldUserUUID holds the string denoting the user_uuid field in the database.
	FieldUserUUID = "user_uuid"
	// FieldTodoContext holds the string denoting the todo_context field in the database.
	FieldTodoContext = "todo_context"
	// FieldCompletion holds the string denoting the completion field in the database.
	FieldCompletion = "completion"
	// FieldCompletedAt holds the string denoting the completed_at field in the database.
	FieldCompletedAt = "completed_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldEditedAt holds the string denoting the edited_at field in the database.
	FieldEditedAt = "edited_at"
	// Table holds the table name of the todo in the database.
	Table = "to_dos"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldTodoUUID,
	FieldUserUUID,
	FieldTodoContext,
	FieldCompletion,
	FieldCompletedAt,
	FieldCreatedAt,
	FieldEditedAt,
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
	// DefaultTodoUUID holds the default value on creation for the "todo_uuid" field.
	DefaultTodoUUID func() uuid.UUID
	// DefaultUserUUID holds the default value on creation for the "user_uuid" field.
	DefaultUserUUID func() uuid.UUID
	// DefaultTodoContext holds the default value on creation for the "todo_context" field.
	DefaultTodoContext string
	// DefaultCompletion holds the default value on creation for the "completion" field.
	DefaultCompletion bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultEditedAt holds the default value on creation for the "edited_at" field.
	DefaultEditedAt time.Time
)
