package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple-api-go/models"
	"simple-api-go/utils"
	"testing"
)

type MockDeviceService struct {
	GetDeviceFunc    func(id string) (*models.Device, error)
	CreateDeviceFunc func(device *models.Device) (*models.Device, error)
	UpdateDeviceFunc func(id string, device *models.Device) (*models.Device, error)
	DeleteDeviceFunc func(id string) error
}

func (m *MockDeviceService) GetDevice(id string) (*models.Device, error) {
	return m.GetDeviceFunc(id)
}

func (m *MockDeviceService) CreateDevice(device *models.Device) (*models.Device, error) {
	return m.CreateDeviceFunc(device)
}

func (m *MockDeviceService) UpdateDevice(id string, device *models.Device) (*models.Device, error) {
	return m.UpdateDeviceFunc(id, device)
}

func (m *MockDeviceService) DeleteDevice(id string) error {
	return m.DeleteDeviceFunc(id)
}

func TestDeviceHandler_CreateDevice(t *testing.T) {
	mockService := &MockDeviceService{}
	handler := NewDeviceHandler(mockService)

	t.Run("CreateDevice", func(t *testing.T) {
		newDevice := &models.Device{
			ID:          "/devices/idTest1",
			Name:        "Device 1",
			DeviceModel: "/devicemodels/Model2",
			Note:        "This is a test device",
			Serial:      "ABC123",
		}

		mockService.CreateDeviceFunc = func(device *models.Device) (*models.Device, error) {
			return device, nil
		}
		mockService.GetDeviceFunc = func(id string) (*models.Device, error) {
			return nil, nil
		}

		reqBody, err := json.Marshal(newDevice)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/api/devices", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.CreateDevice(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusCreated)
		}

		var responseDevice models.Device
		err = json.NewDecoder(rr.Body).Decode(&responseDevice)
		if err != nil {
			t.Errorf("failed to decode response: %v", err)
		}
		if responseDevice != *newDevice {
			t.Errorf("unexpected response device: got %v, want %v", responseDevice, *newDevice)
		}
	})

	t.Run("DuplicateDevice", func(t *testing.T) {
		existingDevice := &models.Device{
			ID:          "/devices/idTest1",
			Name:        "Device 2",
			DeviceModel: "/devicemodels/Model3",
			Note:        "This is another test device",
			Serial:      "DEF456",
		}

		mockService.CreateDeviceFunc = func(device *models.Device) (*models.Device, error) {
			return existingDevice, nil
		}

		mockService.GetDeviceFunc = func(id string) (*models.Device, error) {
			return existingDevice, nil
		}

		reqBody, err := json.Marshal(existingDevice)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/api/devices", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.CreateDevice(rr, req)

		if rr.Code != http.StatusConflict {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusConflict)
		}
	})

	t.Run("ValidateDevice", func(t *testing.T) {
		existingDevice := &models.Device{
			ID:          "Bad idTest1",
			DeviceModel: "Bad Model3",
			Note:        "This is another test device",
			Serial:      "__ Bad DEF456",
		}

		mockService.CreateDeviceFunc = func(device *models.Device) (*models.Device, error) {
			return existingDevice, nil
		}

		mockService.GetDeviceFunc = func(id string) (*models.Device, error) {
			return existingDevice, nil
		}

		reqBody, err := json.Marshal(existingDevice)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/api/devices", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.CreateDevice(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusConflict)
		}
	})
}

func TestDeviceHandler_GetDevice(t *testing.T) {
	mockService := &MockDeviceService{}
	handler := NewDeviceHandler(mockService)

	t.Run("ExistingDevice", func(t *testing.T) {
		expectedDevice := &models.Device{
			ID:          "/devices/idTest1",
			Name:        "Device 1",
			DeviceModel: "/devicemodels/Model2",
			Note:        "This is a test device",
			Serial:      "ABC123",
		}

		mockService.GetDeviceFunc = func(id string) (*models.Device, error) {
			return expectedDevice, nil
		}
		_ = getDeviceIDFromRequest
		_ = func(_ *http.Request) string {
			return "idTest1"
		}

		req, err := http.NewRequest("GET", "/api/devices/idTest1", nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.GetDevice(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusOK)
		}

		var responseDevice models.Device
		err = json.NewDecoder(rr.Body).Decode(&responseDevice)
		if err != nil {
			t.Errorf("failed to decode response: %v", err)
		}

		if responseDevice != *expectedDevice {
			t.Errorf("unexpected response device: got %v, want %v", responseDevice, *expectedDevice)
		}
	})

	t.Run("NonExistingDevice", func(t *testing.T) {
		mockService.GetDeviceFunc = func(id string) (*models.Device, error) {
			return nil, utils.ErrDeviceNotFound
		}

		req, err := http.NewRequest("GET", "/api/devices/ddd2", nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.GetDevice(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusNotFound)
		}
	})
}

func TestDeviceHandler_UpdateDevice(t *testing.T) {
	mockService := &MockDeviceService{}
	handler := NewDeviceHandler(mockService)

	t.Run("UpdateExistingDevice", func(t *testing.T) {
		updatedDevice := &models.Device{
			Name:        "Updated Device",
			DeviceModel: "/devicemodels/Model2",
			Note:        "This is an updated device",
			Serial:      "XYZ789",
		}

		expectedDevice := &models.Device{
			ID:          "/devices/idTest1",
			Name:        "Updated Device",
			DeviceModel: "/devicemodels/Model2",
			Note:        "This is an updated device",
			Serial:      "XYZ789",
		}

		mockService.UpdateDeviceFunc = func(id string, device *models.Device) (*models.Device, error) {
			return expectedDevice, nil
		}

		reqBody, err := json.Marshal(updatedDevice)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}

		req, err := http.NewRequest("PUT", "/api/devices/idTest1", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.UpdateDevice(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusOK)
		}

		var responseDevice models.Device
		err = json.NewDecoder(rr.Body).Decode(&responseDevice)
		if err != nil {
			t.Errorf("failed to decode response: %v", err)
		}

		if responseDevice != *expectedDevice {
			t.Errorf("unexpected response device: got %v, want %v", responseDevice, *expectedDevice)
		}
	})

	t.Run("NonExistingDevice", func(t *testing.T) {
		mockService.UpdateDeviceFunc = func(id string, device *models.Device) (*models.Device, error) {
			return nil, utils.ErrDeviceNotFound
		}

		updatedDevice := &models.Device{
			Name:        "Updated Device",
			DeviceModel: "/devicemodels/Model2",
			Note:        "This is an updated device",
			Serial:      "XYZ789",
		}

		reqBody, err := json.Marshal(updatedDevice)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}

		req, err := http.NewRequest("PUT", "/api/devices/idTest2XYZ", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.UpdateDevice(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusNotFound)
		}
	})
}

func TestDeviceHandler_DeleteDevice(t *testing.T) {
	mockService := &MockDeviceService{}
	handler := NewDeviceHandler(mockService)

	t.Run("DeleteExistingDevice", func(t *testing.T) {
		mockService.DeleteDeviceFunc = func(id string) error {
			return nil
		}

		req, err := http.NewRequest("DELETE", "/api/devices/idTest1", nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.DeleteDevice(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusNoContent)
		}
	})

	t.Run("NonExistingDevice", func(t *testing.T) {
		mockService.DeleteDeviceFunc = func(id string) error {
			return utils.ErrDeviceNotFound
		}

		req, err := http.NewRequest("DELETE", "/api/devices/idTest1", nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler.DeleteDevice(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusNotFound)
		}
	})
}
