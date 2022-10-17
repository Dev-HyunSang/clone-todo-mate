// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/ent/predicate"
	"github.com/dev-hyunsang/clone-todo-mate/ent/todo"
	"github.com/dev-hyunsang/clone-todo-mate/ent/user"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeToDo = "ToDo"
	TypeUser = "User"
)

// ToDoMutation represents an operation that mutates the ToDo nodes in the graph.
type ToDoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	todo_uuid     *uuid.UUID
	user_uuid     *uuid.UUID
	todo_context  *string
	completion    *bool
	completed_at  *time.Time
	created_at    *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ToDo, error)
	predicates    []predicate.ToDo
}

var _ ent.Mutation = (*ToDoMutation)(nil)

// todoOption allows management of the mutation configuration using functional options.
type todoOption func(*ToDoMutation)

// newToDoMutation creates new mutation for the ToDo entity.
func newToDoMutation(c config, op Op, opts ...todoOption) *ToDoMutation {
	m := &ToDoMutation{
		config:        c,
		op:            op,
		typ:           TypeToDo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withToDoID sets the ID field of the mutation.
func withToDoID(id int) todoOption {
	return func(m *ToDoMutation) {
		var (
			err   error
			once  sync.Once
			value *ToDo
		)
		m.oldValue = func(ctx context.Context) (*ToDo, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ToDo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withToDo sets the old ToDo of the mutation.
func withToDo(node *ToDo) todoOption {
	return func(m *ToDoMutation) {
		m.oldValue = func(context.Context) (*ToDo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ToDoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ToDoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ToDoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ToDoMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ToDo.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTodoUUID sets the "todo_uuid" field.
func (m *ToDoMutation) SetTodoUUID(u uuid.UUID) {
	m.todo_uuid = &u
}

// TodoUUID returns the value of the "todo_uuid" field in the mutation.
func (m *ToDoMutation) TodoUUID() (r uuid.UUID, exists bool) {
	v := m.todo_uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldTodoUUID returns the old "todo_uuid" field's value of the ToDo entity.
// If the ToDo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ToDoMutation) OldTodoUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTodoUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTodoUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTodoUUID: %w", err)
	}
	return oldValue.TodoUUID, nil
}

// ResetTodoUUID resets all changes to the "todo_uuid" field.
func (m *ToDoMutation) ResetTodoUUID() {
	m.todo_uuid = nil
}

// SetUserUUID sets the "user_uuid" field.
func (m *ToDoMutation) SetUserUUID(u uuid.UUID) {
	m.user_uuid = &u
}

// UserUUID returns the value of the "user_uuid" field in the mutation.
func (m *ToDoMutation) UserUUID() (r uuid.UUID, exists bool) {
	v := m.user_uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUserUUID returns the old "user_uuid" field's value of the ToDo entity.
// If the ToDo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ToDoMutation) OldUserUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserUUID: %w", err)
	}
	return oldValue.UserUUID, nil
}

// ResetUserUUID resets all changes to the "user_uuid" field.
func (m *ToDoMutation) ResetUserUUID() {
	m.user_uuid = nil
}

// SetTodoContext sets the "todo_context" field.
func (m *ToDoMutation) SetTodoContext(s string) {
	m.todo_context = &s
}

// TodoContext returns the value of the "todo_context" field in the mutation.
func (m *ToDoMutation) TodoContext() (r string, exists bool) {
	v := m.todo_context
	if v == nil {
		return
	}
	return *v, true
}

// OldTodoContext returns the old "todo_context" field's value of the ToDo entity.
// If the ToDo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ToDoMutation) OldTodoContext(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTodoContext is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTodoContext requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTodoContext: %w", err)
	}
	return oldValue.TodoContext, nil
}

// ResetTodoContext resets all changes to the "todo_context" field.
func (m *ToDoMutation) ResetTodoContext() {
	m.todo_context = nil
}

// SetCompletion sets the "completion" field.
func (m *ToDoMutation) SetCompletion(b bool) {
	m.completion = &b
}

// Completion returns the value of the "completion" field in the mutation.
func (m *ToDoMutation) Completion() (r bool, exists bool) {
	v := m.completion
	if v == nil {
		return
	}
	return *v, true
}

// OldCompletion returns the old "completion" field's value of the ToDo entity.
// If the ToDo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ToDoMutation) OldCompletion(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCompletion is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCompletion requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCompletion: %w", err)
	}
	return oldValue.Completion, nil
}

// ResetCompletion resets all changes to the "completion" field.
func (m *ToDoMutation) ResetCompletion() {
	m.completion = nil
}

// SetCompletedAt sets the "completed_at" field.
func (m *ToDoMutation) SetCompletedAt(t time.Time) {
	m.completed_at = &t
}

// CompletedAt returns the value of the "completed_at" field in the mutation.
func (m *ToDoMutation) CompletedAt() (r time.Time, exists bool) {
	v := m.completed_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCompletedAt returns the old "completed_at" field's value of the ToDo entity.
// If the ToDo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ToDoMutation) OldCompletedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCompletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCompletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCompletedAt: %w", err)
	}
	return oldValue.CompletedAt, nil
}

// ResetCompletedAt resets all changes to the "completed_at" field.
func (m *ToDoMutation) ResetCompletedAt() {
	m.completed_at = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ToDoMutation) SetCreatedAt(s string) {
	m.created_at = &s
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ToDoMutation) CreatedAt() (r string, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the ToDo entity.
// If the ToDo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ToDoMutation) OldCreatedAt(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ToDoMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the ToDoMutation builder.
func (m *ToDoMutation) Where(ps ...predicate.ToDo) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ToDoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (ToDo).
func (m *ToDoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ToDoMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.todo_uuid != nil {
		fields = append(fields, todo.FieldTodoUUID)
	}
	if m.user_uuid != nil {
		fields = append(fields, todo.FieldUserUUID)
	}
	if m.todo_context != nil {
		fields = append(fields, todo.FieldTodoContext)
	}
	if m.completion != nil {
		fields = append(fields, todo.FieldCompletion)
	}
	if m.completed_at != nil {
		fields = append(fields, todo.FieldCompletedAt)
	}
	if m.created_at != nil {
		fields = append(fields, todo.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ToDoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldTodoUUID:
		return m.TodoUUID()
	case todo.FieldUserUUID:
		return m.UserUUID()
	case todo.FieldTodoContext:
		return m.TodoContext()
	case todo.FieldCompletion:
		return m.Completion()
	case todo.FieldCompletedAt:
		return m.CompletedAt()
	case todo.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ToDoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case todo.FieldTodoUUID:
		return m.OldTodoUUID(ctx)
	case todo.FieldUserUUID:
		return m.OldUserUUID(ctx)
	case todo.FieldTodoContext:
		return m.OldTodoContext(ctx)
	case todo.FieldCompletion:
		return m.OldCompletion(ctx)
	case todo.FieldCompletedAt:
		return m.OldCompletedAt(ctx)
	case todo.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown ToDo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ToDoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case todo.FieldTodoUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTodoUUID(v)
		return nil
	case todo.FieldUserUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserUUID(v)
		return nil
	case todo.FieldTodoContext:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTodoContext(v)
		return nil
	case todo.FieldCompletion:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCompletion(v)
		return nil
	case todo.FieldCompletedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCompletedAt(v)
		return nil
	case todo.FieldCreatedAt:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown ToDo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ToDoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ToDoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ToDoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ToDo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ToDoMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ToDoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ToDoMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ToDo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ToDoMutation) ResetField(name string) error {
	switch name {
	case todo.FieldTodoUUID:
		m.ResetTodoUUID()
		return nil
	case todo.FieldUserUUID:
		m.ResetUserUUID()
		return nil
	case todo.FieldTodoContext:
		m.ResetTodoContext()
		return nil
	case todo.FieldCompletion:
		m.ResetCompletion()
		return nil
	case todo.FieldCompletedAt:
		m.ResetCompletedAt()
		return nil
	case todo.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown ToDo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ToDoMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ToDoMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ToDoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ToDoMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ToDoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ToDoMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ToDoMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ToDo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ToDoMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ToDo edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	user_uuid     *uuid.UUID
	user_email    *string
	user_password *string
	user_nickname *string
	updated_at    *time.Time
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserUUID sets the "user_uuid" field.
func (m *UserMutation) SetUserUUID(u uuid.UUID) {
	m.user_uuid = &u
}

// UserUUID returns the value of the "user_uuid" field in the mutation.
func (m *UserMutation) UserUUID() (r uuid.UUID, exists bool) {
	v := m.user_uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUserUUID returns the old "user_uuid" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUserUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserUUID: %w", err)
	}
	return oldValue.UserUUID, nil
}

// ResetUserUUID resets all changes to the "user_uuid" field.
func (m *UserMutation) ResetUserUUID() {
	m.user_uuid = nil
}

// SetUserEmail sets the "user_email" field.
func (m *UserMutation) SetUserEmail(s string) {
	m.user_email = &s
}

// UserEmail returns the value of the "user_email" field in the mutation.
func (m *UserMutation) UserEmail() (r string, exists bool) {
	v := m.user_email
	if v == nil {
		return
	}
	return *v, true
}

// OldUserEmail returns the old "user_email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUserEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserEmail: %w", err)
	}
	return oldValue.UserEmail, nil
}

// ResetUserEmail resets all changes to the "user_email" field.
func (m *UserMutation) ResetUserEmail() {
	m.user_email = nil
}

// SetUserPassword sets the "user_password" field.
func (m *UserMutation) SetUserPassword(s string) {
	m.user_password = &s
}

// UserPassword returns the value of the "user_password" field in the mutation.
func (m *UserMutation) UserPassword() (r string, exists bool) {
	v := m.user_password
	if v == nil {
		return
	}
	return *v, true
}

// OldUserPassword returns the old "user_password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUserPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserPassword: %w", err)
	}
	return oldValue.UserPassword, nil
}

// ResetUserPassword resets all changes to the "user_password" field.
func (m *UserMutation) ResetUserPassword() {
	m.user_password = nil
}

// SetUserNickname sets the "user_nickname" field.
func (m *UserMutation) SetUserNickname(s string) {
	m.user_nickname = &s
}

// UserNickname returns the value of the "user_nickname" field in the mutation.
func (m *UserMutation) UserNickname() (r string, exists bool) {
	v := m.user_nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldUserNickname returns the old "user_nickname" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUserNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserNickname: %w", err)
	}
	return oldValue.UserNickname, nil
}

// ResetUserNickname resets all changes to the "user_nickname" field.
func (m *UserMutation) ResetUserNickname() {
	m.user_nickname = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.user_uuid != nil {
		fields = append(fields, user.FieldUserUUID)
	}
	if m.user_email != nil {
		fields = append(fields, user.FieldUserEmail)
	}
	if m.user_password != nil {
		fields = append(fields, user.FieldUserPassword)
	}
	if m.user_nickname != nil {
		fields = append(fields, user.FieldUserNickname)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUserUUID:
		return m.UserUUID()
	case user.FieldUserEmail:
		return m.UserEmail()
	case user.FieldUserPassword:
		return m.UserPassword()
	case user.FieldUserNickname:
		return m.UserNickname()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUserUUID:
		return m.OldUserUUID(ctx)
	case user.FieldUserEmail:
		return m.OldUserEmail(ctx)
	case user.FieldUserPassword:
		return m.OldUserPassword(ctx)
	case user.FieldUserNickname:
		return m.OldUserNickname(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUserUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserUUID(v)
		return nil
	case user.FieldUserEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserEmail(v)
		return nil
	case user.FieldUserPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserPassword(v)
		return nil
	case user.FieldUserNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserNickname(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUserUUID:
		m.ResetUserUUID()
		return nil
	case user.FieldUserEmail:
		m.ResetUserEmail()
		return nil
	case user.FieldUserPassword:
		m.ResetUserPassword()
		return nil
	case user.FieldUserNickname:
		m.ResetUserNickname()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
