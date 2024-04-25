package repositories

import (
	"errors"
	"os"
	"simple-api-go/db"
	"simple-api-go/utils"

	//_ "simple-api-go/repositories"
	"simple-api-go/models"
	"testing"
	//"simple-api-go/repositories"
)

func TestDeviceRepository(t *testing.T) {
	t.Run("MemoryRepository", func(t *testing.T) {
		testDeviceRepository(t, NewDeviceMemoryRepository())
	})

	t.Run("DynamoDBRepository", func(t *testing.T) {
		//db := setupDynamoDBInstance()
		//defer teardownDynamoDBInstance(db)
		// TODO: make another table for test.
		os.Setenv("DYNAMODB_TABLE", "saeid-amn-Devices")
		os.Setenv("REGION", "local")
		os.Setenv("ENDPOINT_URL", "http://localhost:8000")

		testDeviceRepository(t, NewDynamoDeviceService(db.CreateDynamoDBInstance()))
	})
}

func testDeviceRepository(t *testing.T, repo DeviceRepository) {
	t.Run("CreateDevice", func(t *testing.T) {
		device := &models.Device{
			ID:          "idTest1",
			Name:        "Device 1",
			DeviceModel: "Model A",
			Note:        "This is a test device",
			Serial:      "ABC123",
		}

		createdDevice, err := repo.CreateDevice(device)
		if err != nil {
			t.Errorf("CreateDevice() error = %v", err)
			return
		}

		if createdDevice.ID != device.ID {
			t.Errorf("CreateDevice() got = %v, want %v", createdDevice.ID, device.ID)
		}
	})

	t.Run("GetDevice", func(t *testing.T) {
		device, err := repo.GetDevice("idTest1")
		if err != nil {
			t.Errorf("GetDevice() error = %v", err)
			return
		}

		if device.Name != "Device 1" {
			t.Errorf("GetDevice() got = %v, want %v", device.Name, "Device 1")
		}
	})

	t.Run("UpdateDevice", func(t *testing.T) {
		updatedDevice := &models.Device{
			ID:          "idTest1",
			Name:        "Updated Device",
			DeviceModel: "Model B",
			Note:        "This is an updated device",
			Serial:      "DEF456",
		}

		_, err := repo.UpdateDevice("idTest1", updatedDevice)
		if err != nil {
			t.Errorf("UpdateDevice() error = %v", err)
			return
		}

		device, _ := repo.GetDevice("idTest1")
		if device.Name != updatedDevice.Name {
			t.Errorf("UpdateDevice() got = %v, want %v", device.Name, updatedDevice.Name)
		}
	})

	t.Run("DeleteDevice", func(t *testing.T) {
		err := repo.DeleteDevice("idTest1")
		if err != nil {
			t.Errorf("DeleteDevice() error = %v", err)
			return
		}

		_, err = repo.GetDevice("idTest1")
		if !errors.Is(err, utils.ErrDeviceNotFound) {
			t.Errorf("DeleteDevice() device should not exist")
		}
	})
}
