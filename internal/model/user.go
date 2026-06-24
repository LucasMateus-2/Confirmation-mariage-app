package model

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // O hífen impede que a senha vaze no JSON
}
