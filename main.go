package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
		config := &aws.Config{
			Region:   aws.String(os.Getenv("REGION")),
			Endpoint: aws.String(os.Getenv("ENDPOINT")),
		}
		sess := session.Must(session.NewSession(config))

		dynamoDBClient := dynamodb.New(sess)
		dbInstance, err := db.NewDynamoDBInstance(dynamoDBClient, os.Getenv("DYNAMO_TABLE_NAME"))
		if err != nil {
			log.Fatalf("failed to create DynamoDB instance: %v", err)
		}
		return repositories.NewDynamoDeviceService(dbInstance), nil
	default:
		return nil, errors.New("invalid database type")
	}
}
