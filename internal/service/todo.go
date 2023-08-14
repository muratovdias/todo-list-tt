package service

import (
	"github.com/muratovdias/todo-list-tt/internal/models"
	"github.com/muratovdias/todo-list-tt/internal/storage"
)

type ToDoService struct {
	store storage.ToDo
}

func NewTODOService(store storage.ToDo) *ToDoService {
	return &ToDoService{
		store: store,
	}
}

func (t ToDoService) CreateTask(do models.ToDo) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoService) UpdateTask(do models.ToDo) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoService) DeleteTask(s string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoService) MakeTaskDone(s string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoService) TaskList(s string) ([]models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}
