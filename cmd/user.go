package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/database"
	"github.com/dev-hyunsang/clone-todo-mate/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type SearchingEmail struct {
	Email       string    `json:"user_email"`
	RequestDate time.Time `json:"request_date"`
}

func SeachingUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userUUID := claims["uuid"].(uuid.UUID)

	if userUUID == uuid.Nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrResp{
			Code:        "error",
			StatusCode:  fiber.StatusUnauthorized,
			Message:     "로그인이 되지 않았어요. 로그인 이후 다시 시도해 주세요.",
			ErrMessage:  nil,
			RespondedAt: time.Now(),
		})
	}

	req := new(SearchingEmail)
	if err := c.BodyParser(req); err != nil {
		log.Println("[ERROR] Failed to Context BodyParser")
		log.Println(err)
	}

	result, err := database.SeachingUserinToDo(req.Email, req.RequestDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResp{
			StatusCode:  fiber.StatusBadRequest,
			Code:        "error",
			Success:     false,
			Message:     "서버에서 처리하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
			ErrMessage:  err,
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessRespSeachingUserToDo{
		Code:       "error",
		StatusCode: fiber.StatusOK,
		Message:    "성공적으로 입력해 주신 계정을 찾았어요! 팔로우 해 보세요.",
		Data:       result,
		ResponedAt: time.Now(),
	})
}
