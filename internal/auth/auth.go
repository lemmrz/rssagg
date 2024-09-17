package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Api key should be provided from header "Authorization"
// Format "Athorization":"ApiKey {actual key}"
func GetApiKeyFromHeader(header http.Header) (string, error) {
	key := header.Get("Authorization")

	if key == "" {
		return "", errors.New("empty authorization header")
	}
	vals := strings.Split(key, " ")

	if len(vals) != 2 {
		return "", errors.New("incorrect header values amout")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("incorrect header value")
	}
	return vals[1], nil
}
