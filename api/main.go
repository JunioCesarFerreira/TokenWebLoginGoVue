package main

import (
	"fmt"
	"log"
	"main/controller"
	"main/data"
	"main/middleware/auth/jwt"
	"net/http"
	"strconv"

	_ "main/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	port = 8082
)

// @title JWT Login API
// @version 1.0.0
// @contact.name Junio Cesar Ferreira
// @license.name MIT
// @BasePath /
func main() {
	fmt.Println("initializing go server")

	authProvider := jwt.NewJwtAuth()
	//authProvider := paseto.NewPasetoAuth()
	respository := data.NewDummy()
	controller := controller.New(authProvider, respository)

	// Inicialização do roteador do Gorilla Mux
	router := mux.NewRouter()

	router.HandleFunc("/login", controller.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/protected", authProvider.Authenticate(controller.ProtectedEndpoint)).Methods("GET", "OPTIONS")

	//Documentação da API
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	fmt.Printf("listening port %d\n", port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
