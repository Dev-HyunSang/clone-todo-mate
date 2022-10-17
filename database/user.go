package database

import (
	"context"
	"github.com/dev-hyunsang/clone-todo-mate/cmd"
	"github.com/dev-hyunsang/clone-todo-mate/ent"
	"github.com/dev-hyunsang/clone-todo-mate/ent/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CreateUser(user cmd.User) (*ent.User, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, err
	}

	data, err := client.User.
		Create().
		SetUserEmail(user.UserEmail).
		SetUserPassword(user.UserPassword).
		SetUserNickname(user.UserNickname).
		SetUpdatedAt(time.Now()).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return data, nil
}

func LoginUser(inputData *cmd.RequestLogin) (*ent.User, bool, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, false, err
	}

	data, err := client.User.Query().
		Where(user.UserEmail(inputData.UserEmail)).
		Only(context.Background())
	if err != nil {
		return nil, false, err
	}

	// research Email and Password
	// 사용자가 입력한 메일을 첫번째로 검증하며, 두번째로는 패스워드를 검증함.
	// TODO: 로직 검증 필수
	err = bcrypt.CompareHashAndPassword([]byte(data.UserPassword), []byte(inputData.UserPassword))
	if err != nil {
		return nil, false, err
	}

	return data, true, nil
}

func UpdateUser(inputData cmd.User) error {
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

func DeleteUser(inputUUID string) error {
	UserUUID, err := uuid.Parse(inputUUID)
	if err != nil {
		return err
	}

	client, err := ConnectionSQLite()
	if err != nil {
		return err
	}

	user := client.User.Query().
		Where(user.UserUUID(UserUUID)).
		OnlyX(context.Background())

	client.User.DeleteOne(user)

	return nil
}
