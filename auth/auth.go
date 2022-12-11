package auth

import (
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/ent"
	"github.com/golang-jwt/jwt/v4"
)

// 해당하는 유저에 대한 새로운 JWT를 발행합니다.
func NewAuthJWT(user *ent.User) (string, error) {
	claims := jwt.MapClaims{
		"user_uuid":     user.UserUUID,
		"user_nickname": user.UserNickname,
		"exp":           time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("clone_todo_mate"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyJWT(cookie string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		return []byte("clone_todo_mate"), nil
	})
	if err == nil {
		return nil, err
	}

	payload := token.Claims.(jwt.MapClaims)

	return payload, nil
}
