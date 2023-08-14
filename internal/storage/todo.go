package storage

import (
	"context"
	"fmt"
	"github.com/muratovdias/todo-list-tt/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const path = "storage.todo."

type ToDoStore struct {
	collection *mongo.Collection
}

func NewTODOStore(collection *mongo.Collection) *ToDoStore {
	return &ToDoStore{
		collection: collection,
	}
}

func (t ToDoStore) CreateTask(todo models.ToDo) (string, error) {
	res, err := t.collection.InsertOne(context.Background(), todo)

	if err != nil {
		return "", fmt.Errorf("%sCreateTask: %w", path, err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (t ToDoStore) UpdateTask(do models.ToDo) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoStore) DeleteTask(s string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoStore) MakeTaskDone(s string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoStore) TaskList(s string) ([]models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}
