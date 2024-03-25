package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	floodControl "task"
	"task/pkg/handler"
	"task/pkg/repository"
	"task/pkg/service"
)

func main() {
	db, err := repository.NewMongoDB()
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(floodControl.Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Printf("flood-control started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("flood-control shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Client().Disconnect(context.Background()); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
