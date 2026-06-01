package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LucasMateus-2/Confirmation-mariage-app/internal/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository cria uma nova instância conectada ao banco
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`

	// Executa o insert e já joga o ID gerado pelo banco de volta na struct
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
