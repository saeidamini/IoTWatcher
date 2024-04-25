package repositories

import (
	"reflect"
	"simple-api-go/models"
	"testing"
)

func TestDeviceMemoryRepository_CreateDevice(t *testing.T) {
	type fields struct {
		devices map[string]*models.Device
	}
	type args struct {
		device *models.Device
	}
	tests := []struct {
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

	for _, tt := range tests {
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

func TestDeviceMemoryRepository_DeleteDevice(t *testing.T) {
	type fields struct {
		devices map[string]*models.Device
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DeviceMemoryRepository{
				devices: tt.fields.devices,
			}
			if err := r.DeleteDevice(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDevice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeviceMemoryRepository_GetDevice(t *testing.T) {
	type fields struct {
		devices map[string]*models.Device
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Device
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DeviceMemoryRepository{
				devices: tt.fields.devices,
			}
			got, err := r.GetDevice(tt.args.id)
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
	type fields struct {
		devices map[string]*models.Device
	}
	type args struct {
		id     string
		device *models.Device
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Device
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DeviceMemoryRepository{
				devices: tt.fields.devices,
			}
			got, err := r.UpdateDevice(tt.args.id, tt.args.device)
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

func TestNewDeviceMemoryRepository(t *testing.T) {
	tests := []struct {
		name string
		want *DeviceMemoryRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeviceMemoryRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeviceMemoryRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
