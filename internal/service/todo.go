package service

import (
	"errors"
	"fmt"
	"github.com/muratovdias/todo-list-tt/internal/models"
	"github.com/muratovdias/todo-list-tt/internal/storage"
)

var (
	ErrInvalidTitle = errors.New("invalid title, must not be more than 200 character")
	ErrInvalidDate  = errors.New("invalid date")
)

type ToDoService struct {
	store storage.ToDo
}

func NewTODOService(store storage.ToDo) *ToDoService {
	return &ToDoService{
		store: store,
	}
}

func (t ToDoService) CreateTask(todo models.ToDo) (string, error) {
	if err := validateTitle(todo.Title); err != nil {
		return "", fmt.Errorf("service.CreateTask: %w", err)
	}

	if err := validateDate(todo.ActiveAt); err != nil {
		return "", fmt.Errorf("service.CreateTask: %w", err)
	}

	todo.Status = "active"
	return t.store.CreateTask(todo)
}

func (t ToDoService) UpdateTask(todo models.ToDo) (int64, error) {
	if err := validateTitle(todo.Title); err != nil {
		return 0, fmt.Errorf("service.CreateTask: %w", err)
	}

	if err := validateDate(todo.ActiveAt); err != nil {
		return 0, fmt.Errorf("service.CreateTask: %w", err)
	}

	return t.store.UpdateTask(todo)
}

func (t ToDoService) DeleteTask(id string) (int64, error) {
	return t.store.DeleteTask(id)
}

func (t ToDoService) MakeTaskDone(s string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoService) TaskList(s string) ([]models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}
