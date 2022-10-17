package middleware

import "github.com/gofiber/fiber/v2"

func Middleware(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/join")
	auth.Post("/login")
	auth.Post("/logout")

	todo := api.Group("/todo")
	todo.Post("/create")
	todo.Post("/read")
	todo.Post("/read/:uuid")
	todo.Post("/update")
	todo.Post("/completion")
	todo.Delete("/delete/:uuid")
}
