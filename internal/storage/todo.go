package storage

import (
	"github.com/muratovdias/todo-list-tt/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToDoStore struct {
	collection *mongo.Collection
}

func NewTODOStore(collection *mongo.Collection) *ToDoStore {
	return &ToDoStore{
		collection: collection,
	}
}

func (t ToDoStore) CreateTask(do models.ToDo) (string, error) {
	//TODO implement me
	panic("implement me")
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
