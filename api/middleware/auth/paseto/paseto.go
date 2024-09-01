package paseto

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"main/middleware/auth"
	"main/middleware/cors"
	"net/http"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoAuth struct {
	symmetricKey []byte
}

func NewPasetoAuth() *PasetoAuth {
	pl := &PasetoAuth{}
	pl.initialize()
	return pl
}

func (p *PasetoAuth) initialize() {
	// Gera uma chave simétrica de 32 bytes para PASETO
	p.symmetricKey = make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(p.symmetricKey); err != nil {
		fmt.Println("Erro ao gerar a chave simétrica:", err)
	}
}

func (p *PasetoAuth) GetToken(user auth.AuthData, interval int, instance string) (map[string]string, error) {
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
	jsonToken.Set("instance", instance)

	token, err := paseto.NewV2().Encrypt(p.symmetricKey, jsonToken, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("DEBUG new PASETO Token")
	fmt.Println(token)

	resp := map[string]string{"token": token}
	return resp, nil
}

// Middleware para autenticar o token PASETO
func (p *PasetoAuth) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.EnableCors(&w)

		if r.Method == "OPTIONS" {
			return
		}

		tokenString := r.Header.Get("Authorization")

		fmt.Println("DEBUG received PASETO Token")
		fmt.Println(tokenString)

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

		// Recupera a instância da sessão do token e utiliza conforme necessário
		instance := jsonToken.Get("instance")
		fmt.Println("Instance ID:", instance)

		// Cria um novo contexto com o valor da instância
		ctx := context.WithValue(r.Context(), "instance", instance)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (p *PasetoAuth) Renew(tokenString string, w http.ResponseWriter) {
	var jsonToken paseto.JSONToken
	var footer string

	err := paseto.NewV2().Decrypt(tokenString, p.symmetricKey, &jsonToken, &footer)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verifica se o token está próximo do vencimento
	if time.Until(jsonToken.Expiration) < time.Minute {
		userId := jsonToken.Subject
		pass := jsonToken.Get("pass")
		instance := jsonToken.Get("instance")

		newToken, err := p.GetToken(auth.AuthData{UserId: userId, Pwd: pass}, 5, instance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newToken)
	} else {
		w.WriteHeader(http.StatusNotModified) // Não renovado, ainda não precisa
	}
}
