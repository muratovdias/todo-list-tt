package handler

import (
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/muratovdias/todo-list-tt/internal/models"
	"github.com/muratovdias/todo-list-tt/internal/service"
	mock_service "github.com/muratovdias/todo-list-tt/internal/service/mock"
	"github.com/muratovdias/todo-list-tt/pkg/logger"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

const invalidTitle = "Ipd7VSJUJYQj43GbvnN2jslXV7VhwRPMZaxfaGCTlhJahDCSinldI1aE15nXs7jutakDxdQU5F0YPUZ5P9JQNlmk2A1zwNtdpooDgez4V0x9A2DGp3yLl6AWd0MOZCgCXPXEr2ZtAWlD58GPggamk0Db759MOBD5HqSKMhg5f3JI7CZMgnYAbOxRmSVTHGcX2f089UXCU"

func TestHandler_CreateTask(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockToDo, todo models.ToDo)

	testCases := []struct {
		name           string
		inputBody      string
		inputTask      models.ToDo
		mockBehaviour  mockBehaviour
		expectedStatus int
	}{
		{
			name:      "OK",
			inputBody: `{"title":"Test","activeAt":"2023-08-20"}`,
			inputTask: models.ToDo{
				Title:    "Test",
				ActiveAt: "2023-08-20",
			},
			mockBehaviour: func(s *mock_service.MockToDo, todo models.ToDo) {
				s.EXPECT().CreateTask(todo).Return("64d21e4a7ae4d05cf058fef2", nil)
			},
			expectedStatus: 201,
		},
		{
			name:      "Invalid Date",
			inputBody: `{"title":"Test","activeAt":"2023--07"}`,
			inputTask: models.ToDo{
				Title:    "Test",
				ActiveAt: "2023--07",
			},
			mockBehaviour: func(s *mock_service.MockToDo, todo models.ToDo) {
				s.EXPECT().CreateTask(todo).Return("", service.ErrInvalidDate)
			},
			expectedStatus: 400,
		},
		{
			name:      "Invalid Title",
			inputBody: fmt.Sprintf(`{"title":"%s","activeAt":"2023-08-20"}`, invalidTitle),
			inputTask: models.ToDo{
				Title:    invalidTitle,
				ActiveAt: "2023-08-20",
			},
			mockBehaviour: func(s *mock_service.MockToDo, todo models.ToDo) {
				s.EXPECT().CreateTask(todo).Return("", service.ErrInvalidTitle)
			},
			expectedStatus: 400,
		},
		{
			name:           "Empty title",
			inputBody:      `{"title":"","activeAt":"2023-20-08"}`,
			inputTask:      models.ToDo{},
			mockBehaviour:  func(s *mock_service.MockToDo, todo models.ToDo) {},
			expectedStatus: 400,
		},
		{
			name:           "Empty activeAt",
			inputBody:      `{"title":"Test","activeAt":""}`,
			inputTask:      models.ToDo{},
			mockBehaviour:  func(s *mock_service.MockToDo, todo models.ToDo) {},
			expectedStatus: 400,
		},
		{
			name:           "Bad JSON",
			inputBody:      `"title":"Test","activeAt":""}`,
			inputTask:      models.ToDo{},
			mockBehaviour:  func(s *mock_service.MockToDo, todo models.ToDo) {},
			expectedStatus: 500,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			todo := mock_service.NewMockToDo(c)
			tt.mockBehaviour(todo, tt.inputTask)

			s := service.Service{ToDo: todo}
			lg := logger.SetupLogger()
			handler := NewHandler(s, lg)
			handler.validator = validator.New()

			app := fiber.New()
			app.Post("/api/todo-list/tasks", handler.CreateTask)

			req := httptest.NewRequest("POST", "/api/todo-list/tasks", bytes.NewBufferString(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req, -1)

			require.NoError(t, err)
			require.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}
