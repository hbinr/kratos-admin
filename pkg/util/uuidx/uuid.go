package uuidx

import "github.com/google/uuid"

func GenID() string {
	return uuid.New().String()
}
