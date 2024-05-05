package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"simple-api-go/models"
	"simple-api-go/services"
	"simple-api-go/utils"
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
		utils.ErrorJSONFormat(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateDevice(device); err != nil {
		utils.ErrorJSONFormat(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdDevice, err := h.service.CreateDevice(&device)
	if err != nil {
		utils.ErrorJSONFormat(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.ReturnHttpResponse(w, createdDevice, http.StatusCreated)
}

func (h *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	id := getDeviceIDFromRequest(r)
	device, err := h.service.GetDevice(id)
	if err != nil {
		//utils.ErrorJSONFormat(w, err.Error(), http.StatusNotFound)
		utils.ErrorJSONFormat(w, err.Error(), http.StatusNotFound)
		return
	}

	h.ReturnHttpResponse(w, device, http.StatusOK)
}

func (h *DeviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	id := getDeviceIDFromRequest(r)
	var updatedDevice models.Device
	if err := json.NewDecoder(r.Body).Decode(&updatedDevice); err != nil {
		utils.ErrorJSONFormat(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedDevice.ID = id
	if err := validateDevice(updatedDevice); err != nil {
		utils.ErrorJSONFormat(w, err.Error(), http.StatusBadRequest)
		return
	}

	device, err := h.service.UpdateDevice(id, &updatedDevice)
	if err != nil {
		if errors.Is(err, utils.ErrDeviceNotFound) {
			utils.ErrorJSONFormat(w, err.Error(), http.StatusNotFound)
			return
		}
		utils.ErrorJSONFormat(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.ReturnHttpResponse(w, device, http.StatusOK)
}

func (h *DeviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	id := getDeviceIDFromRequest(r)
	if err := h.service.DeleteDevice(id); err != nil {
		if errors.Is(err, utils.ErrDeviceNotFound) {
			utils.ErrorJSONFormat(w, err.Error(), http.StatusNotFound)
			return
		}
		utils.ErrorJSONFormat(w, err.Error(), http.StatusInternalServerError)
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
		utils.ErrorJSONFormat(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validateDevice(device models.Device) error {
	// Check for required fields. Validate the format of the Serial field
	// Improve: Merge multiple errors.
	alphaNumericRegex := regexp.MustCompile(`^[A-Za-z0-9]+$`)
	if !alphaNumericRegex.MatchString(device.ID) {
		return errors.New("invalid ID format, It's must be alphameric format. ")
	}

	if device.Name == "" {
		return errors.New("device name is required")
	}

	if device.DeviceModel == "" {
		return errors.New("device model is required")
	}

	// Sanitize input fields
	device.Name = utils.SanitizeInput(device.Name)
	device.DeviceModel = utils.SanitizeInput(device.DeviceModel)
	device.Note = utils.SanitizeInput(device.Note)
	device.Serial = utils.SanitizeInput(device.Serial)

	// Validate the format of the Serial field
	if !alphaNumericRegex.MatchString(device.Serial) {
		return errors.New("invalid serial format, It's must be alphameric format. ")
	}

	return nil
}
