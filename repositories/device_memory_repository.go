package repositories

import (
	"fmt"
	"simple-api-go/models"
)

type DeviceMemoryRepository struct {
	devices map[string]*models.Device
}

func NewDeviceMemoryRepository() *DeviceMemoryRepository {
	return &DeviceMemoryRepository{
		devices: make(map[string]*models.Device),
	}
}

func (r *DeviceMemoryRepository) GetDevice(id string) (*models.Device, error) {
	device, ok := r.devices[id]
	if !ok {
		return nil, fmt.Errorf("device not found")
	}
	return device, nil
}

func (r *DeviceMemoryRepository) CreateDevice(device *models.Device) error {
	r.devices[device.ID] = device
	return nil
}

func (r *DeviceMemoryRepository) UpdateDevice(id string, device *models.Device) error {
	_, ok := r.devices[id]
	//_, ok := r.devices[device.ID]
	if !ok {
		return fmt.Errorf("device not found")
	}
	r.devices[id] = device
	return nil
}

func (r *DeviceMemoryRepository) DeleteDevice(id string) error {
	_, ok := r.devices[id]
	if !ok {
		return fmt.Errorf("device not found")
	}
	delete(r.devices, id)
	return nil
}
