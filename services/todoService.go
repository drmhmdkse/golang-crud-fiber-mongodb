package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mangoDb/models"
	"mangoDb/repository"
)

//go:generate mockgen -destination=../mocks/service/mockTodoservice.go -package=services mangoDb/services TodoService
type DefaultTodoService struct {
	Repo repository.TodoRepository // interfaceye
}

type TodoService interface {
	TodoInsert(todo models.Todo) (bool, error)
	TodoGetAll() ([]models.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
	TodoUpdate(id primitive.ObjectID, todo models.Todo) (bool, error)
}

func (t DefaultTodoService) TodoInsert(todo models.Todo) (bool, error) {
	if len(todo.Title) < 3 {
		return false, nil
	}
	result, err := t.Repo.Insert(todo)
	if err != nil || result == false {
		return false, err
	}
	return result, nil
}

func (t DefaultTodoService) TodoGetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultTodoService) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil {
		return false, err
	}
	return result, nil
}

func (t DefaultTodoService) TodoUpdate(id primitive.ObjectID, todo models.Todo) (bool, error) {
	result, err := t.Repo.Update(id, todo)
	if err != nil {
		return false, err
	}
	return result, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
