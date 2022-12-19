package database

import (
	"context"
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/ent"
	"github.com/dev-hyunsang/clone-todo-mate/ent/user"
	"github.com/dev-hyunsang/clone-todo-mate/models"
	"github.com/google/uuid"
)

func CreateUser(createUser models.User) (*ent.User, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, err
	}

	data, err := client.User.
		Create().
		SetUserEmail(createUser.UserEmail).
		SetUserPassword(string(createUser.UserPassword)).
		SetUserNickname(createUser.UserNickname).
		SetUpdatedAt(time.Now()).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return data, nil
}

// 사용자가 입력한 정보를 토대로 기존의 사용자 정보를 업데이트 하는 기능
func UpdateUser(inputData models.User) error {
	client, err := ConnectionSQLite()
	if err != nil {
		return err
	}

	_, err = client.User.Update().
		SetUserEmail(inputData.UserEmail).
		SetUserPassword(inputData.UserPassword).
		SetUserNickname(inputData.UserNickname).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func SearchingUserInfo(email string) (*ent.User, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, err
	}

	result, err := client.User.Query().Where(user.UserEmail(email)).Only(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 사용자가 탈퇴는 원하는 경우 탈퇴할 수 있도록 하는 기능
// 사용자의 UUID를 토대로 사용자 정보를 삭제합니다.
func DeleteUser(userUUID uuid.UUID) error {
	client, err := ConnectionSQLite()
	if err != nil {
		return err
	}

	user := client.User.Query().
		Where(user.UserUUID(userUUID)).
		OnlyX(context.Background())

	client.User.DeleteOne(user)

	return nil
}
