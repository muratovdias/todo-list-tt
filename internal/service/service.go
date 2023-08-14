package service

import "github.com/muratovdias/todo-list-tt/internal/models"

type ToDo interface {
	CreateTask(models.ToDo) (string, error)
	UpdateTask(models.ToDo) (int64, error)
	DeleteTask(string) (int64, error)
	MakeTaskDone(string) (int64, error)
	TaskList(string) ([]models.ToDo, error)
}

type Service struct {
	ToDo
}
