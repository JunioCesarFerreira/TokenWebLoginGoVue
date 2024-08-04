package jwt

import (
	"crypto/rand"
	"encoding/json"
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

	fmt.Println("DEBUG new JWT Token")
	fmt.Println(tokenString)

	resp := map[string]string{"token": tokenString}
	return resp, nil
}

// Middleware para autenticar o token JWT
func (j jwtAuth) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.EnableCors(&w)

		if r.Method == "OPTIONS" {
			return
		}

		tokenString := r.Header.Get("Authorization")

		fmt.Println("DEBUG received JWT Token")
		fmt.Println(tokenString)

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

func (j jwtAuth) Renew(tokenString string, w http.ResponseWriter) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verifica se o token está próximo do vencimento
	claims := token.Claims.(jwt.MapClaims)
	expiration := claims["exp"].(float64)
	if time.Now().Unix() > int64(expiration)-60 { // Renova se expirar em menos de 1 minuto
		userId := claims["userId"].(string)
		pass := claims["pass"].(string)

		newToken, err := j.GetToken(auth.AuthData{UserId: userId, Pwd: pass}, 5)
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
