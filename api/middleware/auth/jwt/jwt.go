package jwt

import (
	"crypto/rand"
	"fmt"
	"main/middleware/auth"
	"main/middleware/cors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtAuth struct {
	secretKey []byte
}

func NewJwtAuth() jwtAuth {
	jr := jwtAuth{}
	jr.initialize()
	return jr
}

func (j jwtAuth) initialize() {
	// Cria um slice de bytes para armazenar a chave secreta
	j.secretKey = make([]byte, 32) // Gera uma chave secreta de 32 bytes (256 bits)

	// Preenche o slice de bytes com valores aleatórios
	_, err := rand.Read(j.secretKey)
	if err != nil {
		fmt.Println("Erro ao gerar a chave secreta:", err)
		return
	}
}

func (j jwtAuth) GetToken(user auth.AuthData, interval int) (map[string]string, error) {
	// Gera o token JWT com tempo de expiração
	expirationTime := time.Now().Add(time.Duration(interval) * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.UserId,
		"pass":   user.Pwd,
		"exp":    expirationTime,
	})

	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return nil, err
	}

	resp := map[string]string{"token": tokenString}
	return resp, nil
}

// Middleware para autenticar o token JWT
func (j jwtAuth) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\nauthenticate")

		cors.EnableCors(&w)

		if r.Method == "OPTIONS" {
			return
		}

		tokenString := r.Header.Get("Authorization")

		fmt.Println(tokenString)

		if tokenString == "" {
			fmt.Println("token is empty")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return j.secretKey, nil
		})

		if err != nil {
			fmt.Println("error: ", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expiration, ok := claims["exp"].(float64)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}
			if time.Now().Unix() > int64(expiration) {
				fmt.Println("token expired")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		} else {
			fmt.Println("token not valid")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
}
