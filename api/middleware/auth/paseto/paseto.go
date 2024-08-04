package paseto

import (
	"crypto/rand"
	"fmt"
	"main/middleware/auth"
	"main/middleware/cors"
	"net/http"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type pasetoAuth struct {
	symmetricKey []byte
}

func NewPasetoAuth() pasetoAuth {
	pl := pasetoAuth{}
	pl.initialize()
	return pl
}

func (p pasetoAuth) initialize() {
	// Gera uma chave simétrica de 32 bytes para PASETO
	p.symmetricKey = make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(p.symmetricKey); err != nil {
		fmt.Println("Erro ao gerar a chave simétrica:", err)
	}
}

func (p pasetoAuth) GetToken(user auth.AuthData, interval int) (map[string]string, error) {
	// Gera um token PASETO com tempo de expiração
	now := time.Now()
	expiration := now.Add(time.Duration(interval) * time.Minute)

	// Define as reivindicações
	jsonToken := paseto.JSONToken{
		Expiration: expiration,
		Subject:    user.UserId,
	}

	// Adiciona dados personalizados
	jsonToken.Set("pass", user.Pwd)

	token, err := paseto.NewV2().Encrypt(p.symmetricKey, jsonToken, nil)
	if err != nil {
		return nil, err
	}

	resp := map[string]string{"token": token}
	return resp, nil
}

// Middleware para autenticar o token PASETO
func (p pasetoAuth) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.EnableCors(&w)

		if r.Method == "OPTIONS" {
			return
		}

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			fmt.Println("token is empty")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var jsonToken paseto.JSONToken
		var footer string

		err := paseto.NewV2().Decrypt(tokenString, p.symmetricKey, &jsonToken, &footer)
		if err != nil {
			fmt.Println("error: ", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if jsonToken.Expiration.Before(time.Now()) {
			fmt.Println("token expired")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
