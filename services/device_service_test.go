package services_test

import (
	"errors"
	"simple-api-go/utils"
	"testing"

	"simple-api-go/models"
	"simple-api-go/services"
)

type MockDeviceRepository struct {
	devices map[string]*models.Device
}

func (m *MockDeviceRepository) GetDevice(id string) (*models.Device, error) {
	device, ok := m.devices[id]
	if !ok {
		return nil, utils.ErrDeviceNotFound
	}
	return device, nil
}

func (m *MockDeviceRepository) CreateDevice(device *models.Device) (*models.Device, error) {
	m.devices[device.ID] = device
	return device, nil
}

func (m *MockDeviceRepository) UpdateDevice(id string, device *models.Device) (*models.Device, error) {
	_, ok := m.devices[id]
	if !ok {
		return nil, utils.ErrDeviceNotFound
	}
	m.devices[id] = device
	return device, nil
}

func (m *MockDeviceRepository) DeleteDevice(id string) error {
	_, ok := m.devices[id]
	if !ok {
		return utils.ErrDeviceNotFound
	}
	delete(m.devices, id)
	return nil
}

func TestDeviceService(t *testing.T) {
	mockRepo := &MockDeviceRepository{
		devices: map[string]*models.Device{
			"idTest1": {
				ID:          "idTest1",
				Name:        "Device 1",
				DeviceModel: "Model A",
				Note:        "This is a test device",
				Serial:      "ABC123",
			},
		},
	}

	deviceService := services.NewDeviceService(mockRepo)

	t.Run("GetDevice", func(t *testing.T) {
		device, err := deviceService.GetDevice("idTest1")
		if err != nil {
			t.Errorf("GetDevice() error = %v", err)
			return
		}

		if device.Name != "Device 1" {
			t.Errorf("GetDevice() got = %v, want %v", device.Name, "Device 1")
		}
	})

	t.Run("CreateDevice", func(t *testing.T) {
		device := &models.Device{
			ID:          "2",
			Name:        "Device 2",
			DeviceModel: "Model B",
			Note:        "This is another test device",
			Serial:      "DEF456",
		}

		createdDevice, err := deviceService.CreateDevice(device)
		if err != nil {
			t.Errorf("CreateDevice() error = %v", err)
			return
		}

		if createdDevice.ID != device.ID {
			t.Errorf("CreateDevice() got = %v, want %v", createdDevice.ID, device.ID)
		}
	})

	t.Run("UpdateDevice", func(t *testing.T) {
		updatedDevice := &models.Device{
			ID:          "idTest1",
			Name:        "Updated Device",
			DeviceModel: "Model C",
			Note:        "This is an updated device",
			Serial:      "GHI789",
		}

		_, err := deviceService.UpdateDevice("idTest1", updatedDevice)
		if err != nil {
			t.Errorf("UpdateDevice() error = %v", err)
			return
		}

		device, _ := deviceService.GetDevice("idTest1")
		if device.Name != updatedDevice.Name {
			t.Errorf("UpdateDevice() got = %v, want %v", device.Name, updatedDevice.Name)
		}
	})

	t.Run("DeleteDevice", func(t *testing.T) {
		err := deviceService.DeleteDevice("idTest1")
		if err != nil {
			t.Errorf("DeleteDevice() error = %v", err)
			return
		}

		_, err = deviceService.GetDevice("idTest1")
		if !errors.Is(err, utils.ErrDeviceNotFound) {
			t.Errorf("DeleteDevice() device should not exist")
		}
	})

	t.Run("DeviceNotFound", func(t *testing.T) {
		_, err := deviceService.GetDevice("non-existent-id")
		if !errors.Is(err, utils.ErrDeviceNotFound) {
			t.Errorf("GetDevice() error = %v, want %v", err, utils.ErrDeviceNotFound)
		}
	})
}
