package utils

import (
	"encoding/json"
	"errors"
	"golang.org/x/text/unicode/norm"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
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
	_ = json.NewEncoder(w).Encode(ErrorJSON{Message: err, Status: code})
}

func SanitizeInput(input string) string {
	// Remove leading and trailing spaces
	input = strings.TrimSpace(input)

	// Convert to title case
	title := cases.Title(language.Und)
	input, _, _ = transform.String(title, input)

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	sanitized, _, _ := transform.String(t, input)
	return sanitized
}
