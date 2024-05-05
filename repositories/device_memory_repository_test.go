package repositories

import (
	"reflect"
	"simple-api-go/models"
	"testing"
)

type fields struct {
	devices map[string]*models.Device
}
type args struct {
	device *models.Device
}

func TestDeviceMemoryRepository_CreateDevice(t *testing.T) {
	testsCreate := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Device
		wantErr bool
	}{
		{
			name: "CreateNewDevice",
			fields: fields{
				devices: make(map[string]*models.Device),
			},
			args: args{
				device: &models.Device{
					ID:          "1",
					Name:        "Device 1",
					DeviceModel: "Model A",
					Note:        "This is a test device",
					Serial:      "ABC123",
				},
			},
			want: &models.Device{
				ID:          "1",
				Name:        "Device 1",
				DeviceModel: "Model A",
				Note:        "This is a test device",
				Serial:      "ABC123",
			},
			wantErr: false,
		},
		{
			name: "CreateExistingDevice",
			fields: fields{
				devices: map[string]*models.Device{
					"1": {
						ID:          "1",
						Name:        "Device 1",
						DeviceModel: "Model A",
						Note:        "This is a test device",
						Serial:      "ABC123",
					},
				},
			},
			args: args{
				device: &models.Device{
					ID:          "1",
					Name:        "Device 1",
					DeviceModel: "Model A",
					Note:        "This is a test device",
					Serial:      "ABC123",
				},
			},
			want: &models.Device{
				ID:          "1",
				Name:        "Device 1",
				DeviceModel: "Model A",
				Note:        "This is a test device",
				Serial:      "ABC123",
			},
			wantErr: false,
		},
	}

	for _, tt := range testsCreate {
		t.Run(tt.name, func(t *testing.T) {
			r := &DeviceMemoryRepository{
				devices: tt.fields.devices,
			}
			got, err := r.CreateDevice(tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDevice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceMemoryRepository_GetDevice(t *testing.T) {
	testsGet := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Device
		wantErr bool
	}{
		{
			name: "GetExistingDevice",
			fields: fields{
				devices: map[string]*models.Device{
					"1": {
						ID:          "1",
						Name:        "Device 1",
						DeviceModel: "Model A",
						Note:        "This is a test device",
						Serial:      "ABC123",
					},
				},
			},
			args: args{
				device: &models.Device{
					ID: "1",
				},
			},
			want: &models.Device{
				ID:          "1",
				Name:        "Device 1",
				DeviceModel: "Model A",
				Note:        "This is a test device",
				Serial:      "ABC123",
			},
			wantErr: false,
		},
	}

	for _, tt := range testsGet {
		t.Run(tt.name, func(t *testing.T) {
			r := &DeviceMemoryRepository{
				devices: tt.fields.devices,
			}
			got, err := r.GetDevice(tt.args.device.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDevice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceMemoryRepository_UpdateDevice(t *testing.T) {
	testsUpdate := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Device
		wantErr bool
	}{
		{
			name: "UpdateExistingDevice",
			fields: fields{
				devices: map[string]*models.Device{
					"1": {
						ID: "1",
					},
				},
			},
			args: args{
				device: &models.Device{
					ID:          "1",
					Name:        "New Device 1",
					DeviceModel: "New Model A",
					Note:        "New This is a test device",
					Serial:      "NewABC123",
				},
			},
			want: &models.Device{
				ID:          "1",
				Name:        "New Device 1",
				DeviceModel: "New Model A",
				Note:        "New This is a test device",
				Serial:      "NewABC123",
			},
			wantErr: false,
		},
	}

	for _, tt := range testsUpdate {
		t.Run(tt.name, func(t *testing.T) {
			r := &DeviceMemoryRepository{
				devices: tt.fields.devices,
			}
			got, err := r.UpdateDevice(tt.args.device.ID, tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateDevice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceMemoryRepository_DeleteDevice(t *testing.T) {
	testDelete := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Device
		wantErr bool
	}{
		{
			name: "DeleteExistingDevice",
			fields: fields{
				devices: map[string]*models.Device{
					"1": {
						ID: "1",
					},
				},
			},
			args: args{
				device: &models.Device{
					ID: "1",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range testDelete {
		t.Run(tt.name, func(t *testing.T) {
			r := &DeviceMemoryRepository{
				devices: tt.fields.devices,
			}
			if err := r.DeleteDevice(tt.args.device.ID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDevice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
