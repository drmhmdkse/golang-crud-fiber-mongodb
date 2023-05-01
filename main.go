package main

import (
	"mangoDb/app"
	"mangoDb/configs"
	"mangoDb/repository"
	"mangoDb/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")
	TodoRepositoryDb := repository.NewTodoRepositoryDB(dbClient)
	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Get("/api/todo/:id", td.DeleteTodo)
	appRoute.Post("/api/todo/update/:id", td.UpdateTodo)

	if err := appRoute.Listen(":8080"); err != nil {
		return
	}

}
