package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: Bearer token
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header empty")
	}
	authHeaderSplit := strings.Split(authHeader, " ")
	if len(authHeaderSplit) != 2 || authHeaderSplit[0] != "Bearer" || authHeaderSplit[1] == "" {
		return "", errors.New("malformed authorization header")
	} else {
		return authHeaderSplit[1], nil
	}
}
