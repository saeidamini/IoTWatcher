package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
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

var (
	ErrInvalidDatabaseType = errors.New("invalid database type")
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

	switch os.Getenv("RUNNING_MODE") {
	case "local":
		serverInstance := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
		log.Println("Starting server on " + serverInstance)
		log.Fatal(http.ListenAndServe(serverInstance, router))
	case "aws":
		lambda.Start(httpadapter.New(router).ProxyWithContext)
	default:
		log.Fatalf("Could not runnig application.")
	}
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
		return nil, ErrInvalidDatabaseType
	}
}
