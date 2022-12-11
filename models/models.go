package models

import (
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/ent"
	"github.com/google/uuid"
)

// === Data Struct ===
type User struct {
	UserUUID     uuid.UUID `json:"user_uuid"`
	UserEmail    string    `json:"user_email"`
	UserPassword string    `json:"user_passoword"`
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
// API 구조체 생성 시 필수 사항:
/* type ApiStruct struct{
	Code string `json:"code"`
	StatusCode int `json:"status_code"`
	Success bool `json:"success"`
	Message string `json:"messgae"`
	ErrMessage string `json:"err_message"` 오류 응답 구조체에만 사용함.
	RespondedAt time.Time `json:"resoponded_at"`
}*/

type RequestJoin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserNickname string `json:"user_nickname"`
}

type RequestLogin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type RequestCreateToDo struct {
	ToDoContext    string `json:"todo_context"`
	ToDoCompletion bool   `json:"todo_completion"`
}

type RequestReadToDo struct {
	RequestDate time.Time `json:"request_date"`
}

// === Response ===
type SuccessRespUserData struct {
	Code        string    `json:"code"`
	StatusCode  int       `json:"status_code"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	Data        *ent.User `json:"data"`
	RespondedAt time.Time `json:"responded_at"`
}

// 구조체 이름이 길어지는 건 기분탓...
type SuccessRespSeachingUserToDo struct {
	Code       string `json:"code"`
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       []*ent.ToDo
	ResponedAt time.Time `json:"responed_at"`
}

type SuccessReadToDo struct {
	Code       string `json:"code"`
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       []*ent.ToDo
	ResponedAt time.Time `json:"responed_at"`
}

type SuccessResp struct {
	Code        string    `json:"code"`
	StatusCode  int       `json:"status_code"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	RespondedAt time.Time `json:"responded_at"`
}

type SuccessLoginResp struct {
	Code       string `json:"code"`
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       struct {
		Token string `json:"token"`
	}
	ResopondedAt time.Time `json:"resoponded_at"`
}

type ErrResp struct {
	Code        string    `json:"code"`
	StatusCode  int       `json:"status_code"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	ErrMessage  error     `json:"err_message"`
	RespondedAt time.Time `json:"responded_at"`
}
