package main_test

import (
	"os"
	main "simple-api-go"
	"simple-api-go/db"
	"simple-api-go/repositories"
	. "testing"
)

func TestNewDeviceRepository(t *T) {
	testCases := []struct {
		name           string
		databaseType   string
		expected       repositories.DeviceRepository
		expectedError  error
		setupEnvVars   func()
		teardownEnvVar func()
	}{
		{
			name:          "MemoryRepository",
			databaseType:  "memory",
			expected:      repositories.NewDeviceMemoryRepository(),
			expectedError: nil,
			setupEnvVars: func() {
				_ = os.Setenv("DATABASE_TYPE", "memory")
			},
			teardownEnvVar: func() {
				_ = os.Unsetenv("DATABASE_TYPE")
			},
		},
		{
			name:          "DynamoDBRepository",
			databaseType:  "dynamodb",
			expected:      repositories.NewDynamoDeviceService(&db.DynamoDBInstance{}),
			expectedError: nil,
			setupEnvVars: func() {
				_ = os.Setenv("DATABASE_TYPE", "dynamodb")
			},
			teardownEnvVar: func() {
				_ = os.Unsetenv("DATABASE_TYPE")
			},
		},
		{
			name:           "InvalidDatabaseType",
			databaseType:   "invalid",
			expected:       nil,
			expectedError:  main.ErrInvalidDatabaseType,
			setupEnvVars:   func() {},
			teardownEnvVar: func() {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *T) {
			tc.setupEnvVars()
			defer tc.teardownEnvVar()

			repo, err := main.NewDeviceRepository()
			if err != nil && tc.expectedError == nil {
				t.Errorf("NewDeviceRepository() unexpected error: %v", err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("NewDeviceRepository() expected error: %v, got nil", tc.expectedError)
			} else if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("NewDeviceRepository() expected error: %v, got: %v", tc.expectedError, err)
			}

			if repo == nil && tc.expected != nil {
				t.Errorf("NewDeviceRepository() expected repository: %v, got nil", tc.expected)
			}
			//else if repo != nil && repo != tc.expected {
			//	t.Errorf("NewDeviceRepository() expected repository: %v, got: %v", tc.expected, repo)
			//}
		})
	}
}
