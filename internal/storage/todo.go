package storage

import "go.mongodb.org/mongo-driver/mongo"

type ToDoStore struct {
	collection *mongo.Collection
}

func NewTODOStore(collection *mongo.Collection) *ToDoStore {
	return &ToDoStore{
		collection: collection,
	}
}
