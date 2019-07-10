package utils

import(
	guuid "github.com/google/uuid"
)

func generateUUID() string {

	uuid := guuid.New()
	return uuid.String()
}
