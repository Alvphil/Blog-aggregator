package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("authorization")
	if authHeader == "" {
		return "", errors.New("missing auth header")
	}
	const keyText = "ApiKey "
	if !strings.HasPrefix(authHeader, keyText) {
		return "", errors.New("invalid Authorization header format")
	}
	ApiKey := strings.TrimPrefix(authHeader, keyText)

	return ApiKey, nil
}
