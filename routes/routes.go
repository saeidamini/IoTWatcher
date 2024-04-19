package routes

import (
	"net/http"
	"simple-api-go/handlers"
)

func SetupRoutes(handler *handlers.DeviceHandler) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /api/devices", handler.CreateDevice)
	router.HandleFunc("GET /api/devices/{id}", handler.GetDevice)
	router.HandleFunc("PUT /api/devices/{id}", handler.UpdateDevice)
	router.HandleFunc("DELETE /api/devices/{id}", handler.DeleteDevice)

	return router
}
