package auth

import (
	"net/http"
)

type Auth interface {
	Authenticate(next http.HandlerFunc) http.HandlerFunc
	GetToken(user AuthData, interval int, instance string) (map[string]string, error)
	Renew(tokenString string, w http.ResponseWriter)
}
