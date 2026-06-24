package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lucas/confirmation-mariage-app/internal/model"
	"github.com/lucas/confirmation-mariage-app/internal/repository"

	"github.com/lucas/confirmation-mariage-app/pkg/hash"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

func (s *AuthService) Login(input LoginInput) (*TokenResponse, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("credenciais inválidas")
	}
	if !hash.Check(input.Password, user.Password) {
		return nil, errors.New("credenciais inválidas")
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{Token: token, User: *user}, nil
}

func generateToken(userID int64) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
