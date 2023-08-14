package db

import (
	"context"
	"fmt"
	"github.com/muratovdias/todo-list-tt/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	path = "pkg.db.ConnectDB" //путь чтобы отслеживать ошибку в логах
)

func ConnectDB(cfg config.DB) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	//пингуем нашу базу
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	return client.Database(cfg.Name), nil
}
