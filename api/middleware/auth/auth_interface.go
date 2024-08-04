package auth

import (
	"net/http"
)

type Auth interface {
	Authenticate(next http.HandlerFunc) http.HandlerFunc
	GetToken(user AuthData, interval int) (map[string]string, error)
}
