package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/auth"
	"github.com/dev-hyunsang/clone-todo-mate/database"
	"github.com/dev-hyunsang/clone-todo-mate/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreaeteToDo(c *fiber.Ctx) error {
	user := new(models.User)
	cookie := c.Cookies("jwt")

	payload, err := auth.VerifyJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrResp{
			Code:        "Unauthorized",
			StatusCode:  fiber.StatusUnauthorized,
			Success:     false,
			Message:     "로그인 이후 다시 시도해 주세요.",
			RespondedAt: time.Now(),
		})
	}

	stringUUID := payload["user_uuid"].(string)
	user.UserUUID = uuid.MustParse(stringUUID)

	req := new(models.RequestCreateToDo)
	if err := c.BodyParser(req); err != nil {
		log.Println(err)
	}

	todo := models.ToDo{
		ToDoUUID:       uuid.New(),
		UserUUID:       user.UserUUID,
		ToDoContext:    req.ToDoContext,
		ToDoCompletion: false,
		CreatedAt:      time.Now(),
		EditedAt:       time.Now(),
	}

	// 새로운 할일을 추가합니다.
	if err := database.CreateToDo(todo); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResp{
			Code:        "error",
			StatusCode:  fiber.StatusBadRequest,
			Message:     "새로운 할일을 추가하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
			ErrMessage:  err,
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResp{
		Code:        "success",
		StatusCode:  fiber.StatusOK,
		Success:     true,
		Message:     "성공적으로 할일을 생성했어요!",
		RespondedAt: time.Now(),
	})
}

func DeleteToDo(c *fiber.Ctx) error {
	user := new(models.User)
	cookie := c.Cookies("jwt")

	payload, err := auth.VerifyJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrResp{
			Code:        "Unauthorized",
			StatusCode:  fiber.StatusUnauthorized,
			Success:     false,
			Message:     "로그인 이후 다시 시도해 주세요.",
			RespondedAt: time.Now(),
		})
	}

	stringUUID := payload["user_uuid"].(string)
	user.UserUUID = uuid.MustParse(stringUUID)

	req := new(models.RequestReadToDo)
	if err := c.BodyParser(req); err != nil {
		log.Println(err)
	}

	todo, err := database.AllReadToDo(user.UserUUID, req.RequestDate)
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(models.ErrResp{
			Code:        "error",
			StatusCode:  fiber.StatusNoContent,
			Success:     false,
			Message:     "해당 일자에 할일을 찾을 수 없네요. 추가 후 다시 시도해 주세요.",
			ErrMessage:  nil,
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessReadToDo{
		Code:       "success",
		StatusCode: fiber.StatusOK,
		Message:    "성공적으로 할 일을 불러왔어요!",
		Data:       todo,
		ResponedAt: time.Now(),
	})
}
