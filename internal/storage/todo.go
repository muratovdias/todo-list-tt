package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/muratovdias/todo-list-tt/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const path = "storage.todo."

var ErrInvalidId = errors.New("invalid object id")

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

func (t ToDoStore) UpdateTask(todo models.ToDo) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(todo.ID)
	if err != nil {
		return 0, fmt.Errorf("%sUpdateTask: %w", path, ErrInvalidId)
	}

	todo.ID = "" // так как поле '_id' неизменяемое, нужно его обнулить, чтобы не было ошибки
	update := bson.M{
		"$set": bson.M{
			"title":    todo.Title,
			"activeAt": todo.ActiveAt,
		},
	}

	res, err := t.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return 0, fmt.Errorf("%sUpdateTask: %w", path, err)
	}

	return res.ModifiedCount, nil
}

func (t ToDoStore) DeleteTask(id string) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("%sDeleteTask: %w", path, ErrInvalidId)
	}

	res, err := t.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return 0, fmt.Errorf("%sDeleteTask: %w", path, err)
	}
	return res.DeletedCount, nil
}

func (t ToDoStore) MakeTaskDone(s string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoStore) TaskList(s string) ([]models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}
