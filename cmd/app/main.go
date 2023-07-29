package main

import (
	"TaskManager/internal/config"
	"TaskManager/internal/storage/sqlite_storage"
	"TaskManager/internal/storage/memory_storage"
	"TaskManager/internal/usecase"
	"log"
	"errors"
	"TaskManager/internal/delivery/http_server"
)

func chooseStorage(appConfig config.Config) (usecase.TaskStorage, error) {
	switch appConfig.StorageType {
	case "memory":
		return memory_storage.New()
	case "sqlite":
		return sqlite_storage.New()
	default:
		return nil, errors.New("wrong storage type")
	}
}

func main() {
	appConfig, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	taskStorage, err := chooseStorage(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	taskInteractor := usecase.NewTaskInteractor(taskStorage)
	handlers := http_server.NewHandler(taskInteractor)

	srv := new(http_server.Server)

	err = srv.Run(appConfig.HttpPort, handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
