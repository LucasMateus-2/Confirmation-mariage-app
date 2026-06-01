package repository

import (
	"context"
	"testing"

	"github.com/LucasMateus-2/Confirmation-mariage-app/internal/domain"
)

func TestInMemoryUserRepository(t *testing.T) {
	repo := NewInMemoryUserRepository()
	ctx := context.Background()

	user := &domain.User{Email: "teste@email.com", Password: "hashedpassword"}

	// Teste Criar
	err := repo.Create(ctx, user)
	if err != nil {
		t.Fatalf("esperava erro nil, obteve: %v", err)
	}

	// Teste Buscar por Email
	found, err := repo.GetByEmail(ctx, "teste@email.com")
	if err != nil {
		t.Fatalf("esperava encontrar usuário, obteve erro: %v", err)
	}

	if found.Email != user.Email {
		t.Errorf("esperava email %s, obteve %s", user.Email, found.Email)
	}
}
