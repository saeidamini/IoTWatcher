package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"simple-api-go/models"
	"simple-api-go/services"
	"simple-api-go/utils"
	"strings"
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
	//Check if the device ID already exists
	existingDevice, err := h.service.GetDevice(device.ID)
	if err == nil && existingDevice != nil {
		utils.ErrorJSONFormat(w, utils.ErrDeviceDuplicate.Error(), http.StatusConflict)
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
	id := "/devices/" + getDeviceIDFromRequest(r)
	device, err := h.service.GetDevice(id)
	if err != nil {
		//utils.ErrorJSONFormat(w, err.Error(), http.StatusNotFound)
		utils.ErrorJSONFormat(w, err.Error(), http.StatusNotFound)
		return
	}

	h.ReturnHttpResponse(w, device, http.StatusOK)
}

func (h *DeviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	id := "/devices/" + getDeviceIDFromRequest(r)
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
	id := "/devices/" + getDeviceIDFromRequest(r)
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
	if len(r.PathValue("id")) > 1 {
		return r.PathValue("id")
	} else {
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) >= 3 {
			return parts[3]
		}
	}
	return ""
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
	var errorMessages []string

	deviceIDRegex := regexp.MustCompile(`^/devices/[A-Za-z0-9]+$`)
	if !deviceIDRegex.MatchString(device.ID) {
		errorMessages = append(errorMessages, "invalid ID format, It must be in the format '/devices/alphanumeric'")
	}

	if device.Name == "" {
		errorMessages = append(errorMessages, "device name is required")
	}

	if device.DeviceModel == "" {
		errorMessages = append(errorMessages, "device model is required")
	}

	deviceModelRegex := regexp.MustCompile(`^/devicemodels/[A-Za-z0-9]+$`)
	if !deviceModelRegex.MatchString(device.DeviceModel) {
		errorMessages = append(errorMessages, "invalid DeviceModel format, It must be in the format '/devicemodels/alphanumeric'")
	}

	// Sanitize input fields
	device.Name = utils.SanitizeInput(device.Name)
	//device.DeviceModel = utils.SanitizeInput(device.DeviceModel)
	device.Note = utils.SanitizeInput(device.Note)
	device.Serial = utils.SanitizeInput(device.Serial)

	// Validate the format of the Serial field
	alphaNumericRegex := regexp.MustCompile(`^[A-Za-z0-9]+$`)
	if !alphaNumericRegex.MatchString(device.Serial) {
		errorMessages = append(errorMessages, "invalid serial format, It's must be alphameric format")
	}

	if len(errorMessages) > 0 {
		return errors.New(strings.Join(errorMessages, "; "))
	}

	return nil
}
