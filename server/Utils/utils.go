package utils

import (
	"github.com/google/uuid"
)

func GenerateShortURL() string {
	newUUID := uuid.New()
	return newUUID.String()[:8]
}
