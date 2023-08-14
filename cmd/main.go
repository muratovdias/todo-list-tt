package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/muratovdias/todo-list-tt/internal/config"
	"github.com/muratovdias/todo-list-tt/internal/handler"
	"github.com/muratovdias/todo-list-tt/internal/service"
	"github.com/muratovdias/todo-list-tt/internal/storage"
	"github.com/muratovdias/todo-list-tt/pkg/db"
	"github.com/muratovdias/todo-list-tt/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slog"
	"os"
)

func main() {

	cfg := config.LoadConfig()
	fmt.Println(cfg)

	newLogger := logger.SetupLogger()

	connectDB, err := db.ConnectDB(cfg.DB)
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Error(err.Error())
		}
	}(connectDB.Client(), context.Background())

	if err != nil {
		log.Error("failed to connect mongo", slog.Any("error", err.Error()))
		os.Exit(1)
	}

	newStorage := storage.NewStorage(connectDB)
	newService := service.NewService(*newStorage)
	newHandler := handler.NewHandler(*newService, newLogger)
	app := newHandler.Routes(cfg)

	err = app.Listen(cfg.Address)
	if err != nil {
		log.Error(err.Error())
		return
	}

}
