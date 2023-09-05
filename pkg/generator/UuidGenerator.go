package generator

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.NewString()
}

func UUIDFrom(uuidStr string) (string, error) {
	uuidValue, err := uuid.Parse(uuidStr)
	if err != nil {
		return "", ErrParseUUID(uuidStr, err)
	}

	return uuidValue.String(), nil
}
