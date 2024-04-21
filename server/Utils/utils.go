package utils

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
)

func GenerateShortURL() string {
	newUUID, _ := uuid.NewRandom()
	return newUUID.String()[:8]
}

func ValidateURL(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return fmt.Errorf("invalid URL: %s", err.Error())
	}
	return nil
}

func ValidateExpiryDate(date string) (int64, error) {
	if date == "" {
		return 0, nil
	}

	expiryDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return expiryDate.Unix(), fmt.Errorf("invalid expiry date format")
	}

	currentTime := time.Now().Unix()

	// Check if expiry date is in the past
	if expiryDate.Unix() < currentTime {
		return expiryDate.Unix(), fmt.Errorf("expiry date cannot be in the past")
	}

	return expiryDate.Unix(), nil
}
