package utils

import uuid "github.com/gofrs/uuid"

func GenUUID() string {
	result, _ := uuid.NewV4()
	return result.String()
}
