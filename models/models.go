package models

import (
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/ent"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// === Data Struct ===
type User struct {
	UserUUID     uuid.UUID `json:"user_uuid"`
	UserEmail    string    `json:"user_email"`
	UserPassword string    `json:"user_passowrd"`
	UserNickname string    `json:"user_nickname"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type ToDo struct {
	ToDoUUID       uuid.UUID `json:"todo_uuid"`
	UserUUID       uuid.UUID `json:"user_uuid"`
	ToDoContext    string    `json:"todo_context"`
	ToDoCompletion bool      `json:"todo_completion"`
	CreatedAt      time.Time `json:"created_at"`
	CompletedAt    time.Time `json:"completion_at"`
	EditedAt       time.Time `json:"edited_at"`
}

// === API Struct ===
type RequestLogin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type SuccessRespUserData struct {
	Code        int       `json:"code"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	Data        *ent.User `json:"data"`
	RespondedAt time.Time `json:"responded_at"`
}

type SuccessResp struct {
	Code        int       `json:"code"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	RespondedAt time.Time `json:"responded_at"`
}

type SuccessLoginResp struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Token jwt.Token `json:"token"`
	}
	ResopondedAt time.Time `json:"resoponded_at"`
}

type ErrResp struct {
	Code        int       `json:"code"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	ErrMessage  error     `json:"err_message"`
	RespondedAt time.Time `json:"responded_at"`
}
