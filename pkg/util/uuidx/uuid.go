package uuidx

import "github.com/google/uuid"

func GenID() uint32 {
	return uuid.New().ID()
}
