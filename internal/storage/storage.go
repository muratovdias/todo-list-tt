package storage

import (
	"github.com/muratovdias/todo-list-tt/internal/models"
	"github.com/muratovdias/todo-list-tt/internal/storage/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToDo interface {
	CreateTask(models.ToDo) (string, error)
	UpdateTask(models.ToDo) (int64, error)
	DeleteTask(string) (int64, error)
	MakeTaskDone(string) (int64, error)
	TaskList(string) ([]models.ToDo, error)
}

type Storage struct {
	ToDo
}

func NewStorage(db *mongo.Database) *Storage {
	return &Storage{
		NewTODOStore(mongodb.CollectionToDo(db, "tasks")),
	}
}
