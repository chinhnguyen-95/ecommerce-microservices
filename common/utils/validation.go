package utils

import (
	"errors"
	"strings"
)

// ValidateRequiredFields ensures all required fields are present.
func ValidateRequiredFields(data map[string]string, fields []string) error {
	for _, field := range fields {
		if strings.TrimSpace(data[field]) == "" {
			return errors.New("missing required field: " + field)
		}
	}
	return nil
}
