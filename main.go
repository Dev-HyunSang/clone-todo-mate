package main

import (
	"github.com/dev-hyunsang/clone-todo-mate/middleware"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	middleware.Middleware(app)

	if err := app.Listen(":3000"); err != nil {
		log.Println("[ERROR] Failed to Server Starting...")
		panic(err)
	}
}
