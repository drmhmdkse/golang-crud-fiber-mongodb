package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mangoDb/mocks/repository"
	"mangoDb/models"
	"testing"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{primitive.NewObjectID(), "title 1", "content 1"},
	{primitive.NewObjectID(), "title 2", "content 2"},
	{primitive.NewObjectID(), "title 3", "content 3"},
	{primitive.NewObjectID(), "title 4", "content 4"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)

	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultTodoService_TodoGetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil) // Mock'a fake veri gönderiyoruz // TodoRepository interface'indeki GetAll() fonksiyonu çağrılıyor
	result, err := service.TodoGetAll()              // burada fake database'den dönen datayı alıyoruz
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}

func TestDefaultTodoService_TodoDelete(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().Delete(gomock.Any()).Return(true, nil)
	result, err := service.TodoDelete(primitive.NewObjectID())
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}

func TestDefaultTodoService_TodoInsert(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().Insert(FakeData[0]).Return(true, nil)
	result, err := service.TodoInsert(FakeData[0])
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}

func TestDefaultTodoService_TodoUpdate(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().Update(FakeData[0].Id, FakeData[0]).Return(true, nil)
	result, err := service.TodoUpdate(FakeData[0].Id, FakeData[0])
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}
