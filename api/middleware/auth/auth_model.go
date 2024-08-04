package auth

// Estrutura do usuário
type AuthData struct {
	UserId string `json:"userId"`
	Pwd    string `json:"pass"`
}
