package repositories

import "simple-api-go/models"

type DeviceRepository interface {
	GetDevice(id string) (*models.Device, error)
	CreateDevice(device *models.Device) error
	UpdateDevice(id string, device *models.Device) error
	DeleteDevice(id string) error
}
