package app

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mangoDb/models"
	"mangoDb/services"
	"net/http"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo // TODO : gelen body struct'ın yapısına uygun olmayabilir kontrol et

	if err := c.BodyParser(&todo); err != nil { // Todo : bodyParser görevini yapmıyor
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if todo.Title == "" || todo.Content == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": "Bad Request"})
	}

	result, err := h.Service.TodoInsert(todo)

	if err != nil || result == false {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"Status": "Success"})
}

func (h TodoHandler) GetAllTodo(c *fiber.Ctx) error {
	result, err := h.Service.TodoGetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h TodoHandler) DeleteTodo(c *fiber.Ctx) error {

	query := c.Params("id")
	primId, _ := primitive.ObjectIDFromHex(query)
	result, err := h.Service.TodoDelete(primId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	if result == false { // burada result false döndüğünde id bulunamadı demektir
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"Status": "Not Found"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Status": "Success"})
}

func (h TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	primId, _ := primitive.ObjectIDFromHex(c.Params("id"))
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	todo.Id = primId
	result, err := h.Service.TodoUpdate(primId, todo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	if result == false {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"Status": "Not Found"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Status": "Success"})
}
