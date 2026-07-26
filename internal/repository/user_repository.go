package repository

import (
	"errors"

	"github.com/lucas/confirmation-mariage-app/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	// Usa LOWER para ignorar case-sensitivity (PostgreSQL é case-sensitive)
	result := r.db.Where("LOWER(email) = LOWER(?)", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Não encontrou – retorna nil sem erro
		return nil, nil
	}
	if result.Error != nil {
		// Outro erro (conexão, etc.)
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id int) (*model.User, error) {
	var user model.User

	result := r.db.First(&user, id) // ou r.db.Where("id = ?", id).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
