package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/clone-todo-mate/auth"
	"github.com/dev-hyunsang/clone-todo-mate/database"
	"github.com/dev-hyunsang/clone-todo-mate/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func JoinUser(c *fiber.Ctx) error {
	req := new(models.RequestJoin)
	if err := c.BodyParser(req); err != nil {
		panic(err)
	}

	userUUID := uuid.New()
	pwHash, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassword), 10)

	userData := models.User{
		UserUUID:     userUUID,
		UserEmail:    req.UserEmail,
		UserPassword: string(pwHash),
		UserNickname: req.UserNickname,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
	}

	// 추후 중복성 체크 기능 추가 예정.

	data, err := database.CreateUser(userData)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResp{
			Code:        "error",
			StatusCode:  fiber.StatusInternalServerError,
			Success:     false,
			Message:     "사용자에 대한 요청을 처리하던 도중 데이터베이스에 대한 오류가 발생 했어요. 잠시후 다시 시도해 주세요.",
			ErrMessage:  err,
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessRespUserData{
		Code:        "success",
		StatusCode:  fiber.StatusOK,
		Success:     true,
		Message:     "어서오세요. 하루를 더 체계적으로 계획적인 하루를 보내 봐요.",
		Data:        data,
		RespondedAt: time.Now(),
	})
}

func LoginUser(c *fiber.Ctx) error {
	req := new(models.RequestLogin)
	if err := c.BodyParser(req); err != nil {
		panic(err)
	}

	userData, err := database.SearchingUserInfo(req.UserEmail)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResp{
			Code:        "error",
			StatusCode:  fiber.StatusBadRequest,
			Success:     false,
			Message:     "입력해 주신 메일 또는 패스워드를 찾을 수 없어요. 확인 후 다시 시도해 주세요.",
			ErrMessage:  err,
			RespondedAt: time.Now(),
		})
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.UserPassword), []byte(req.UserPassword)); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrResp{
			Code:        "error",
			StatusCode:  fiber.StatusBadRequest,
			Success:     false,
			Message:     "입력해 주신 메일 또는 패스워드를 찾을 수 없어요. 확인 후 다시 시도해 주세요.",
			ErrMessage:  err,
			RespondedAt: time.Now(),
		})
	}

	token, err := auth.NewAuthJWT(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResp{
			Code:        "error",
			StatusCode:  fiber.StatusInternalServerError,
			Success:     false,
			Message:     "JWT 토큰 발행 중 알 수 없는 오류가 발생 했어요. 잠시후 다시 시도해 주세요.",
			ErrMessage:  err,
			RespondedAt: time.Now(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: token,
	})

	return c.Status(fiber.StatusOK).JSON(models.SuccessLoginResp{
		Code:       "success",
		StatusCode: fiber.StatusOK,
		Success:    true,
		Data: struct {
			Token string "json:\"token\""
		}{
			Token: token,
		},
		ResopondedAt: time.Now(),
	})
}

func DeleteUser(c *fiber.Ctx) error {
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

	if err = database.DeleteUser(user.UserUUID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrResp{
			Code:        "error",
			StatusCode:  fiber.StatusInternalServerError,
			Success:     false,
			Message:     "사용자를 삭제하던 도중 오류가 발생했어요. 잠시후 다시 시도 주세요.",
			ErrMessage:  err,
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResp{
		Code:       "success",
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "성공적으로 사용자를 삭제했어요. 사용자",
	})
}
