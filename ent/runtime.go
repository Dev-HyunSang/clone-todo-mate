// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/ent/schema"
	"github.com/dev-hyunsang/clone-todo-mate/ent/todo"
	"github.com/dev-hyunsang/clone-todo-mate/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.ToDo{}.Fields()
	_ = todoFields
	// todoDescTodoUUID is the schema descriptor for todo_uuid field.
	todoDescTodoUUID := todoFields[0].Descriptor()
	// todo.DefaultTodoUUID holds the default value on creation for the todo_uuid field.
	todo.DefaultTodoUUID = todoDescTodoUUID.Default.(func() uuid.UUID)
	// todoDescUserUUID is the schema descriptor for user_uuid field.
	todoDescUserUUID := todoFields[1].Descriptor()
	// todo.DefaultUserUUID holds the default value on creation for the user_uuid field.
	todo.DefaultUserUUID = todoDescUserUUID.Default.(func() uuid.UUID)
	// todoDescTodoContext is the schema descriptor for todo_context field.
	todoDescTodoContext := todoFields[2].Descriptor()
	// todo.DefaultTodoContext holds the default value on creation for the todo_context field.
	todo.DefaultTodoContext = todoDescTodoContext.Default.(string)
	// todoDescCompletion is the schema descriptor for completion field.
	todoDescCompletion := todoFields[3].Descriptor()
	// todo.DefaultCompletion holds the default value on creation for the completion field.
	todo.DefaultCompletion = todoDescCompletion.Default.(bool)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[5].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(time.Time)
	// todoDescEditedAt is the schema descriptor for edited_at field.
	todoDescEditedAt := todoFields[6].Descriptor()
	// todo.DefaultEditedAt holds the default value on creation for the edited_at field.
	todo.DefaultEditedAt = todoDescEditedAt.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserUUID is the schema descriptor for user_uuid field.
	userDescUserUUID := userFields[0].Descriptor()
	// user.DefaultUserUUID holds the default value on creation for the user_uuid field.
	user.DefaultUserUUID = userDescUserUUID.Default.(func() uuid.UUID)
	// userDescUserEmail is the schema descriptor for user_email field.
	userDescUserEmail := userFields[1].Descriptor()
	// user.DefaultUserEmail holds the default value on creation for the user_email field.
	user.DefaultUserEmail = userDescUserEmail.Default.(string)
	// userDescUserPassword is the schema descriptor for user_password field.
	userDescUserPassword := userFields[2].Descriptor()
	// user.DefaultUserPassword holds the default value on creation for the user_password field.
	user.DefaultUserPassword = userDescUserPassword.Default.(string)
	// userDescUserNickname is the schema descriptor for user_nickname field.
	userDescUserNickname := userFields[3].Descriptor()
	// user.DefaultUserNickname holds the default value on creation for the user_nickname field.
	user.DefaultUserNickname = userDescUserNickname.Default.(string)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
