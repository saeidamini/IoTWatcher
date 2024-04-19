package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
	"simple-api-go/handlers"
	"simple-api-go/repositories"
	"simple-api-go/routes"
	"simple-api-go/services"
)

func main() {
	fmt.Println("Simple API!")

	deviceRepo := repositories.NewDeviceMemoryRepository()
	deviceSvc := services.NewDeviceService(deviceRepo)
	deviceHandler := handlers.NewDeviceHandler(deviceSvc)

	router := routes.SetupRoutes(deviceHandler)

	serverPath := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	log.Println("Starting server on " + serverPath)
	log.Fatal(http.ListenAndServe(serverPath, router))
}
