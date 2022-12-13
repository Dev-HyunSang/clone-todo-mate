package database

import (
	"context"
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/ent"
	"github.com/dev-hyunsang/clone-todo-mate/ent/todo"
	"github.com/dev-hyunsang/clone-todo-mate/ent/user"
	"github.com/dev-hyunsang/clone-todo-mate/models"
	"github.com/google/uuid"
)

// 새로운 할일을 생성합니다.
// 구조체를 입력 받으며, 오류를
func CreateToDo(inputData models.ToDo) error {
	client, err := ConnectionSQLite()
	if err != nil {
		return err
	}

	_, err = client.ToDo.Create().
		SetTodoUUID(inputData.ToDoUUID).
		SetUserUUID(inputData.UserUUID).
		SetTodoContext(inputData.ToDoContext).
		SetCompletion(inputData.ToDoCompletion).
		SetCompletedAt(inputData.CompletedAt).
		SetCreatedAt(inputData.CreatedAt).
		Save(context.Background())

	return err
}

// 모든 ToDo 항목을 절때로 불러오지 않습니다. 프론트 상에서 처리하기 어려움.
// 요청하는 날짜에만, 사용자가 선택한 날짜를 선택하여 요청하는 경우 해당 날짜의 할일을 불러옴.
func AllReadToDo(userUUID uuid.UUID, date time.Time) ([]*ent.ToDo, error) {
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

func SeachingUserinToDo(email string, date time.Time) ([]*ent.ToDo, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, err
	}

	searchUser, err := client.User.Query().
		Where(user.UserEmail(email)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	resultToDo, err := client.ToDo.Query().
		Where(todo.UserUUID(searchUser.UserUUID)).
		Where(todo.CreatedAt(date)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return resultToDo, nil
}

func DeleteToDo(todoUUID, userUUID uuid.UUID) error {
	client, err := ConnectionSQLite()
	if err != nil {
		return err
	}

	deleteToDo := client.ToDo.Query().
		Where(todo.UserUUID(userUUID)). // 생성한 사용자의 UUID 확인함.
		Where(todo.TodoUUID(todoUUID)).
		OnlyX(context.Background())

	err = client.ToDo.DeleteOne(deleteToDo).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}
