package main

import (
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
	"simple-api-go/db"
	"simple-api-go/handlers"
	"simple-api-go/repositories"
	"simple-api-go/routes"
	"simple-api-go/services"
)

func main() {
	fmt.Println("Simple API!")

	//deviceRepo := repositories.NewDeviceMemoryRepository()
	deviceRepo, err := NewDeviceRepository()
	if err != nil {
		log.Fatalf("failed to connect to the database instance: %v", err)
		return
	}

	deviceSvc := services.NewDeviceService(deviceRepo)
	deviceHandler := handlers.NewDeviceHandler(deviceSvc)

	router := routes.SetupRoutes(deviceHandler)

	serverInstance := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	log.Println("Starting server on " + serverInstance)
	log.Fatal(http.ListenAndServe(serverInstance, router))
}

func NewDeviceRepository() (repositories.DeviceRepository, error) {
	dbType := os.Getenv("DATABASE_TYPE")
	switch dbType {
	case "memory":
		return repositories.NewDeviceMemoryRepository(), nil
	case "dynamodb":
		dbInstance := db.CreateDynamoDBInstance()
		return repositories.NewDynamoDeviceService(dbInstance), nil
	default:
		return nil, errors.New("invalid database type")
	}
}
