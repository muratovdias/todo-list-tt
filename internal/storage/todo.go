package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/muratovdias/todo-list-tt/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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

func (t ToDoStore) MakeTaskDone(id string) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("%sMakeTaskDone: %w", path, ErrInvalidId)
	}

	status := bson.M{
		"$set": bson.M{
			"status": "done",
		},
	}

	res, err := t.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, status)
	if err != nil {
		return 0, fmt.Errorf("%sMakeTaskDone: %w", path, err)
	}

	return res.ModifiedCount, nil
}

func (t ToDoStore) TaskList(status string) ([]models.ToDo, error) {
	var filter bson.M

	// если статус active, тогда возвращаем все задачи у которых activeAt <= текущего дня;
	if status == "active" {
		filter = bson.M{
			"activeAt": bson.M{
				"$lte": time.Now().Format("2006-01-02"), // $lte - выберает документы < или == указанному значению.
			},
			"status": bson.M{
				"$eq": "active",
			},
		}
	} else {
		filter = bson.M{
			"status": bson.M{
				"$eq": "done",
			},
		}
	}

	// задачи должны быть отсортированы по дате создания;
	opts := options.Find().SetSort(bson.D{{Key: "activeAt", Value: -1}}) //значение -1 сортирует документы по указанному атрибуту в порядке убывания

	cur, err := t.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("%sTaskList: %w", path, err)
	}
	defer cur.Close(context.Background())

	todos := make([]models.ToDo, 0, cur.RemainingBatchLength())
	for cur.Next(context.Background()) {
		var todo models.ToDo
		err := cur.Decode(&todo)
		if err != nil {
			return nil, fmt.Errorf("%sTaskList: %w", path, err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
