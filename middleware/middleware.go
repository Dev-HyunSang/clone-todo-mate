package middleware

import (
	"github.com/dev-hyunsang/clone-todo-mate/cmd"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/join", cmd.JoinUser)
	auth.Post("/login", cmd.LoginUser)

	todo := api.Group("/todo")
	todo.Post("/create", cmd.CreaeteToDo)
	todo.Get("/read", cmd.ReadToDo)
	todo.Post("/edit", cmd.EditToDo)
	todo.Post("/completion", cmd.CompletionToDo)
	todo.Delete("/delete", cmd.DeleteToDo)
	todo.Delete("/delete/:uuid", cmd.DeleteToDoParameter)
}
