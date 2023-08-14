package service

import "github.com/muratovdias/todo-list-tt/internal/storage"

type ToDoService struct {
	store storage.ToDo
}

func NewTODOService(store storage.ToDo) *ToDoService {
	return &ToDoService{
		store: store,
	}
}
