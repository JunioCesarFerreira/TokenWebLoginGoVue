package controller

import (
	"encoding/json"
	"fmt"
	"main/data"
	"main/middleware/auth"
	"main/middleware/cors"
	"net/http"
)

type Controller struct {
	authProvider auth.Auth
	repository   data.Repository
}

func New(a auth.Auth, r data.Repository) Controller {
	return Controller{
		authProvider: a,
		repository:   r,
	}
}

// @Summary Tenta validar usuário e fazer login
// @Description Verifica se usuário é registrado e realiza login
// @Tags Login
// @Accept  json
// @Produce  json
// @Param LoginRequest body auth.AuthData true "Usuário e Hash"
// @Success 200
// @Router /login [post]
func (c Controller) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---\nlogin endpoint\n---")

	cors.EnableCors(&w)

	if r.Method == "OPTIONS" {
		return
	}

	var user auth.AuthData
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("nameid:", user.UserId)
	fmt.Println("pwdHashed:", user.Pwd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	check, err := c.repository.CheckUser(user.UserId, user.Pwd)
	if !check || err != nil {
		http.Error(w, "Invalid name or password", http.StatusUnauthorized)
		return
	}

	resp, err := c.authProvider.GetToken(user, 1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

// @Summary Rota protegida com autenticação utilizando Token
// @Description Testando Token
// @Tags Login
// @Produce json
// @Param Authorization header string true "Token"
// @Success 200
// @Router /protected [get]
func (c Controller) ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---\nprotected endpoint\n---")

	resp := map[string]string{"message": "Protected route.\nIt works!"}
	json.NewEncoder(w).Encode(resp)
}

// @Summary Renova o token de autenticação
// @Description Renova o token de autenticação se o token atual for válido e está prestes a expirar
// @Tags Login
// @Produce json
// @Param Authorization header string true "Token"
// @Success 200 {string} string "Token renovado com sucesso"
// @Router /renew [get]
func (c Controller) RenewToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---\nrenew endpoint\n---")
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	c.authProvider.Renew(tokenString, w)
}
