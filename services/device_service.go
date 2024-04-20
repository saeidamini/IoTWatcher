package services

import (
	"errors"
	"simple-api-go/models"
	"simple-api-go/repositories"
)

var (
	ErrDeviceNotFound = errors.New("device not found")
	// Other custom errors
)

type DeviceService interface {
	CreateDevice(device *models.Device) (*models.Device, error)
	GetDevice(id string) (*models.Device, error)
	UpdateDevice(id string, device *models.Device) (*models.Device, error)
	DeleteDevice(id string) error
}

type deviceService struct {
	repo repositories.DeviceRepository
}

func NewDeviceService(repo repositories.DeviceRepository) DeviceService {
	return &deviceService{
		repo: repo,
	}
}

func (s *deviceService) GetDevice(id string) (*models.Device, error) {
	return s.repo.GetDevice(id)
}

func (s *deviceService) CreateDevice(device *models.Device) (*models.Device, error) {
	return s.repo.CreateDevice(device)
}

func (s *deviceService) UpdateDevice(id string, device *models.Device) (*models.Device, error) {
	return s.repo.UpdateDevice(id, device)
}

func (s *deviceService) DeleteDevice(id string) error {
	return s.repo.DeleteDevice(id)
}
