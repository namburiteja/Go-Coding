package token

import (
	"encoding/base64"
	"errors"
)

// Create converts the last player ID into an opaque pagination token.
func Create(playerID string) (string, error) {
	if playerID == "" {
		return "", errors.New("player ID cannot be empty")
	}

	return base64.RawURLEncoding.EncodeToString(
		[]byte(playerID),
	), nil
}

// Decode converts the pagination token back into a player ID.
func Decode(value string) (string, error) {
	if value == "" {
		return "", errors.New("pagination token cannot be empty")
	}

	decoded, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return "", errors.New("invalid pagination token")
	}

	playerID := string(decoded)

	if playerID == "" {
		return "", errors.New("invalid pagination token")
	}

	return playerID, nil
}