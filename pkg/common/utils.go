package common

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateID generates a random ID of specified length
func GenerateID(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

// FormatCurrency formats a float64 as currency string
func FormatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

// ParseDate parses a date string in various formats
func ParseDate(dateStr string) (time.Time, error) {
	formats := []string{
		"2006-01-02",
		"02/01/2006",
		"01/02/2006",
		time.RFC3339,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}

// IsValidEmail performs basic email validation
func IsValidEmail(email string) bool {
	// Basic email validation - in production you might want to use a more robust solution
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	atIndex := -1
	dotIndex := -1

	for i, char := range email {
		if char == '@' {
			if atIndex != -1 {
				return false // Multiple @ symbols
			}
			atIndex = i
		} else if char == '.' {
			dotIndex = i
		}
	}

	return atIndex > 0 && dotIndex > atIndex && dotIndex < len(email)-1
}

// TruncateString truncates a string to the specified length
func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength-3] + "..."
}
