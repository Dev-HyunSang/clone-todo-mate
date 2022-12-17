package main

import (
	"context"
	"log"

	"github.com/dev-hyunsang/clone-todo-mate/database"
	"github.com/dev-hyunsang/clone-todo-mate/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Middleware(app)

	client, err := database.ConnectionSQLite()
	if err != nil {
		log.Println("[ERROR] Failed to Connection DataBase")
		log.Println(err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Println("[ERROR] Failed to Schema Create")
		log.Println(err)
	}

	if err := app.Listen(":3000"); err != nil {
		log.Println("[ERROR] Failed to Server Starting...")
		panic(err)
	}
}
