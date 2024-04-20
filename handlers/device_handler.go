package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"simple-api-go/models"
	"simple-api-go/services"
)

type DeviceHandler struct {
	service services.DeviceService
}

func NewDeviceHandler(service services.DeviceService) *DeviceHandler {
	return &DeviceHandler{service: service}
}

func (h *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//createdDevice := h.service.CreateDevice(&device)
	createdDevice, err := h.service.CreateDevice(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.ReturnHttpResponse(w, createdDevice, http.StatusCreated)
}

func (h *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	id := getDeviceIDFromRequest(r)
	device, err := h.service.GetDevice(id)
	if err != nil {
		if errors.Is(err, services.ErrDeviceNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.ReturnHttpResponse(w, device, http.StatusOK)
}

func (h *DeviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	id := getDeviceIDFromRequest(r)
	var updatedDevice models.Device
	if err := json.NewDecoder(r.Body).Decode(&updatedDevice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//device := h.service.UpdateDevice(id, &updatedDevice)
	device, err := h.service.UpdateDevice(id, &updatedDevice)
	if err != nil {
		if errors.Is(err, services.ErrDeviceNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.ReturnHttpResponse(w, device, http.StatusOK)
}

func (h *DeviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	id := getDeviceIDFromRequest(r)
	if err := h.service.DeleteDevice(id); err != nil {
		if errors.Is(err, services.ErrDeviceNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.ReturnHttpResponse(w, nil, http.StatusNoContent)
}

func getDeviceIDFromRequest(r *http.Request) string {
	// Example: /devices/{id}
	return r.PathValue("id") // Device ID
}

func (h *DeviceHandler) ReturnHttpResponse(w http.ResponseWriter, createdDevice *models.Device, httpCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	err := json.NewEncoder(w).Encode(createdDevice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}