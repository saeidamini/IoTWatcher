package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrDeviceNotFound = errors.New("device not found")
)

type ErrorJSON struct {
	Message string `json:"message"`
	Status  int    `json:"status,omitempty"`
}

func ErrorJSONFormat(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorJSON{Message: err, Status: code})
}
