package repository

import (
	"database/sql"
	"errors"

	"github.com/lucas/confirmation-mariage-app/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	row := r.db.QueryRow(`SELECT id, email, password FROM users WHERE email = $1`, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id int) (*model.User, error) {
	user := &model.User{}
	row := r.db.QueryRow(`SELECT id, email FROM users WHERE id = $1`, id)
	err := row.Scan(&user.ID, &user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
