package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	services "mangoDb/mocks/service"
	"mangoDb/models"
	"net/http/httptest"
	"testing"
)

var td TodoHandler
var mocService *services.MockTodoService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mocService = services.NewMockTodoService(ctrl)

	td = TodoHandler{Service: mocService}
	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_GetAllTodo(t *testing.T) {

	trd := setup(t)
	defer trd()

	router := fiber.New()
	router.Get("api/todos", td.GetAllTodo)

	var FakeDataForHandler = []models.Todo{
		{primitive.NewObjectID(), "title 1", "content 1"},
		{primitive.NewObjectID(), "title 2", "content 2"},
		{primitive.NewObjectID(), "title 3", "content 3"},
		{primitive.NewObjectID(), "title 4", "content 4"},
	}
	mocService.EXPECT().TodoGetAll().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)
	resp, _ := router.Test(req, 1)
	assert.Equal(t, 200, resp.StatusCode)

}

func TestTodoHandler_DeleteTodo(t *testing.T) {

	trd := setup(t)
	defer trd()

	router := fiber.New()
	router.Get("api/todo/:id", td.DeleteTodo)
	id := primitive.NewObjectID()
	mocService.EXPECT().TodoDelete(id).Return(true, nil)
	target := fmt.Sprintf("/api/todo/%s", id.Hex())
	req := httptest.NewRequest("GET", target, nil)
	resp, _ := router.Test(req, 1)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTodoHandler_CreateTodo(t *testing.T) {

	trd := setup(t)
	defer trd()

	router := fiber.New()
	router.Post("api/todo", td.CreateTodo)
	todo := models.Todo{
		Id:      primitive.NewObjectID(),
		Title:   "title1",
		Content: "content1",
	}
	mocService.EXPECT().TodoInsert(todo).Return(true, nil)
	jsonVeriTwo, _ := json.Marshal(todo)
	req := httptest.NewRequest("POST", "/api/todo", bytes.NewBuffer(jsonVeriTwo))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := router.Test(req, 1)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestTodoHandler_UpdateTodo(t *testing.T) {

	trd := setup(t)
	defer trd()

	router := fiber.New()
	router.Post("api/todo/update/:id", td.UpdateTodo)
	todo := models.Todo{
		Id:      primitive.NewObjectID(),
		Title:   "title1",
		Content: "content1",
	}
	mocService.EXPECT().TodoUpdate(todo.Id, todo).Return(true, nil)
	jsonTodo, _ := json.Marshal(todo)
	target := fmt.Sprintf("/api/todo/update/%s", todo.Id.Hex())
	req := httptest.NewRequest("POST", target, bytes.NewBuffer(jsonTodo))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := router.Test(req, 1)
	assert.Equal(t, 200, resp.StatusCode)
}
