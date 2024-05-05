package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple-api-go/models"
	"simple-api-go/services"
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

func TestDeviceHandler_GetDevice(t *testing.T) {
	mockService := &MockDeviceService{}
	handler := NewDeviceHandler(mockService)

	t.Run("ExistingDevice", func(t *testing.T) {
		expectedDevice := &models.Device{
			ID:          "1",
			Name:        "Device 1",
			DeviceModel: "Model A",
			Note:        "This is a test device",
			Serial:      "ABC123",
		}

		mockService.GetDeviceFunc = func(id string) (*models.Device, error) {
			return expectedDevice, nil
		}

		req, err := http.NewRequest("GET", "/devices/1", nil)
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
			return nil, services.ErrDeviceNotFound
		}

		req, err := http.NewRequest("GET", "/devices/2", nil)
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

func TestDeviceHandler_CreateDevice(t *testing.T) {
	mockService := &MockDeviceService{}
	handler := NewDeviceHandler(mockService)

	t.Run("CreateDevice", func(t *testing.T) {
		newDevice := &models.Device{
			ID:          "1",
			Name:        "Device 1",
			DeviceModel: "Model A",
			Note:        "This is a test device",
			Serial:      "ABC123",
		}

		mockService.CreateDeviceFunc = func(device *models.Device) (*models.Device, error) {
			return device, nil
		}

		reqBody, err := json.Marshal(newDevice)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/devices", bytes.NewBuffer(reqBody))
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
}
