package main

import (
	"TaskManager/internal/config"
	"TaskManager/internal/storage/sqlite_storage"
	"TaskManager/internal/storage/postgres_storage"
	"TaskManager/internal/storage/memory_storage"
	"TaskManager/internal/usecase"
	"TaskManager/internal/delivery/cli_application"
	"TaskManager/internal/delivery/http_server"
	"context"
	"log"
	"errors"
	"os"
	"os/signal"
)

// @title Task Manager API
// @version 1.0
// @description API Server for TaskManager Application

// @host localhost:8080
// @BasePath /

func chooseStorage(appConfig config.Config) (usecase.TaskStorage, error) {
	switch appConfig.StorageType {
	case "memory":
		return memory_storage.New()
	case "sqlite":
		return sqlite_storage.New()
	case "postgres":
		return postgres_storage.New(appConfig.Postgres)
	default:
		return nil, errors.New("wrong storage type")
	}
}

func chooseInterface(appConfig config.Config, taskInteractor *usecase.TaskInteractor) {
	switch appConfig.Interface {
	case "cli":
		taskCLI := cli_application.NewTaskCLI(taskInteractor)

		err := taskCLI.Run()
		if err != nil {
			log.Fatal(err)
		}
	case "rest_api":
		handlers := http_server.NewHandler(taskInteractor)

		srv := new(http_server.Server)
		
		go func() {
			if err := srv.Run(appConfig.HttpPort, handlers.InitRoutes()); err != nil {
				log.Fatalf("error occured while running http server: %s", err.Error())
			}
		}()
		log.Print("App Started")

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
	
		log.Print("App Shutting Down")

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatal("error occured on server shutting down: %s", err.Error())
		}
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
	
	chooseInterface(appConfig, taskInteractor)
}
