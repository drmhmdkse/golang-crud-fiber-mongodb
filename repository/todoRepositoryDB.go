package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mangoDb/models"
	"time"
)

//go:generate mockgen -destination=../mocks/repository/mockTodoRepository.go -package=repository mangoDb/repository TodoRepository
type TodoRepositoryDB struct {
	todoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
	Update(id primitive.ObjectID, todo models.Todo) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todo.Id = primitive.NewObjectID()
	result, err := t.todoCollection.InsertOne(ctx, todo)
	if result.InsertedID != nil || err != nil {
		return false, err
	}
	return true, nil
}

func (t TodoRepositoryDB) GetAll() ([]models.Todo, error) {
	var todo models.Todo
	var todos []models.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.todoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.todoCollection.DeleteOne(ctx, bson.M{"id": id})

	if result.DeletedCount == 0 {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (t TodoRepositoryDB) Update(id primitive.ObjectID, todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.todoCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": todo})
	if result.ModifiedCount != 1 {
		return false, errors.New("todo not found")
	}
	if err != nil {
		return false, err
	}

	return true, nil

}

func NewTodoRepositoryDB(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{todoCollection: dbClient}
}
