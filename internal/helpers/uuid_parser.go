package helpers

import "github.com/google/uuid"

func UUIDParser(id string) uuid.UUID {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}
	return parsedUUID
}
