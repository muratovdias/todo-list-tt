package service

import (
	"errors"
	"fmt"
	"github.com/muratovdias/todo-list-tt/internal/models"
	"github.com/muratovdias/todo-list-tt/internal/storage"
	"strings"
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
	// проверяем количество символов в заголовке
	if err := validateTitle(todo.Title); err != nil {
		return "", fmt.Errorf("service.CreateTask: %w", err)
	}

	// проверяем формат входящей даты
	if err := validateDate(todo.ActiveAt); err != nil {
		return "", fmt.Errorf("service.CreateTask: %w", err)
	}

	todo.Status = "active"
	return t.store.CreateTask(todo)
}

func (t ToDoService) UpdateTask(todo models.ToDo) (int64, error) {
	// проверяем количество символов в заголовке
	if err := validateTitle(todo.Title); err != nil {
		return 0, fmt.Errorf("service.CreateTask: %w", err)
	}

	// проверяем формат входящей даты
	if err := validateDate(todo.ActiveAt); err != nil {
		return 0, fmt.Errorf("service.CreateTask: %w", err)
	}

	return t.store.UpdateTask(todo)
}

func (t ToDoService) DeleteTask(id string) (int64, error) {
	return t.store.DeleteTask(id)
}

func (t ToDoService) MakeTaskDone(id string) (int64, error) {
	return t.store.MakeTaskDone(id)
}

func (t ToDoService) TaskList(status string) ([]models.ToDo, error) {
	list, err := t.store.TaskList(status)
	if err != nil {
		return nil, err
	}

	// проверяем день каждого элемента в слайсе на выходной (суббота/воскресенье)
	for i := range list {
		ok, err := isWeekend(list[i].ActiveAt)
		if err != nil {
			return nil, err
		}

		//если этот день выходной, помечаем этот день как ВЫХОДНОЙ
		if ok {
			// эффективная конкатенация строк
			var builder strings.Builder
			builder.WriteString("ВЫХОДНОЙ - ")
			builder.WriteString(list[i].Title)

			list[i].Title = builder.String()
		}
	}

	return list, nil
}
