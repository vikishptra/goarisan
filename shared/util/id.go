package util

import "github.com/google/uuid"

func GenerateID() string {
	ID := uuid.New().String()

	return ID
}

// https://github.com/matoous/go-nanoid
// https://github.com/segmentio/ksuid
