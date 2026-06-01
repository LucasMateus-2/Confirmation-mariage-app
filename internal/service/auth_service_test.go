package service

import (
	"context"
	"testing"

	"github.com/LucasMateus-2/Confirmation-mariage-app/internal/repository"
)

func TestAuthService_RegisterAndLogin(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := NewAuthService(repo, "secret_key")
	ctx := context.Background()

	email := "user@test.com"
	password := "supersecret123"

	// 1. Testar Registro
	err := svc.Register(ctx, email, password)
	if err != nil {
		t.Fatalf("registro falhou: %v", err)
	}

	// 2. Testar Login Válido
	token, err := svc.Login(ctx, email, password)
	if err != nil {
		t.Fatalf("login falhou para credenciais válidas: %v", err)
	}
	if token == "" {
		t.Error("token retornado está vazio")
	}

	// 3. Testar Login Inválido
	_, err = svc.Login(ctx, email, "senha_errada")
	if err == nil {
		t.Error("esperava erro para senha inválida, mas veio nil")
	}
}
