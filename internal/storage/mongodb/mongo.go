package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const path = "storage.mongodb."

func CollectionToDo(db *mongo.Database, name string) *mongo.Collection {
	collection := db.Collection(name)
	// Создаем уникальные индексы (title, activeAt)
	CreateTodoIndex(collection)

	return collection
}

// CreateTodoIndex - делает атрибуты title и activeAt в коллекции 'tasks' уникальными
func CreateTodoIndex(collection *mongo.Collection) {
	index := mongo.IndexModel{
		Keys:    bson.D{{"title", 1}, {"activeAt", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		log.Fatalf("%screateIndex: %s", path, err.Error())
	}

}
