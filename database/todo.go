package database

import (
	"context"
	"github.com/dev-hyunsang/clone-todo-mate/cmd"
	"github.com/dev-hyunsang/clone-todo-mate/ent"
	"github.com/dev-hyunsang/clone-todo-mate/ent/todo"
	"github.com/google/uuid"
)

// 새로운 할일을 생성합니다.
// 구조체를 입력 받으며, 오류를
func CreateToDo(inputData cmd.ToDo) error {
	client, err := ConnectionSQLite()
	if err != nil {
		return err
	}

	_, err = client.ToDo.Create().
		SetTodoUUID(inputData.ToDoUUID).
		SetUserUUID(inputData.UserUUID).
		SetTodoContext(inputData.ToDoContext).
		SetCompletion(inputData.Completion).
		SetCompletedAt(inputData.CompletedAt).
		SetCreatedAt(inputData.CreatedAt).
		Save(context.Background())

	return err
}

func AllReadToDo(userUUID uuid.UUID, date string) ([]*ent.ToDo, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, err
	}

	data, err := client.ToDo.Query().
		Where(todo.CreatedAt(date)).
		Where(todo.UserUUID(userUUID)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CompletionToDo(userUUID, todoUUID uuid.UUID) error {
	client, err := ConnectionSQLite()
	if err != nil {
		return err
	}

	_, err = client.ToDo.Update().
		Where(todo.UserUUID(userUUID)).
		Where(todo.TodoUUID(todoUUID)).
		SetCompletion(true).
		Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}
