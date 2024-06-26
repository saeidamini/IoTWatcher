package repositories

import (
	"simple-api-go/models"
	"simple-api-go/utils"
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
		return nil, utils.ErrDeviceNotFound
	}
	return device, nil
}

func (r *DeviceMemoryRepository) CreateDevice(device *models.Device) (*models.Device, error) {
	r.devices[device.ID] = device
	return device, nil
}

func (r *DeviceMemoryRepository) UpdateDevice(id string, device *models.Device) (*models.Device, error) {
	_, ok := r.devices[id]
	//_, ok := r.devices[device.ID]
	if !ok {
		return nil, utils.ErrDeviceNotFound
	}
	r.devices[id] = device
	return device, nil
}

func (r *DeviceMemoryRepository) DeleteDevice(id string) error {
	_, ok := r.devices[id]
	if !ok {
		return utils.ErrDeviceNotFound
	}
	delete(r.devices, id)
	return nil
}
