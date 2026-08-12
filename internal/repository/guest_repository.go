package repository

import (
	"errors"

	"github.com/lucas/confirmation-mariage-app/internal/model"
	"gorm.io/gorm"
)

type GuestRepository struct {
	db *gorm.DB
}

func NewGuestRepository(db *gorm.DB) *GuestRepository {
	return &GuestRepository{db: db}
}

func (r *GuestRepository) FindAll() ([]model.Guest, error) {
	var guests []model.Guest
	err := r.db.Preload("PlusOnes").Order("name").Find(&guests).Error
	if err != nil {
		return nil, err
	}
	return guests, nil
}

func (r *GuestRepository) FindByID(id int) (*model.Guest, error) {
	var guest model.Guest
	err := r.db.Preload("PlusOnes").First(&guest, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &guest, nil
}

func (r *GuestRepository) SearchByName(name string) ([]model.Guest, error) {
	var guests []model.Guest
	err := r.db.
		Preload("PlusOnes").
		Where("name ILIKE ?", "%"+name+"%").
		Order("name").
		Find(&guests).Error
	if err != nil {
		return nil, err
	}
	return guests, nil
}

// Confirm atualiza o status do convidado principal e de cada agregado,
// dentro de uma transação para garantir atomicidade.
func (r *GuestRepository) Confirm(guestID int, input model.ConfirmInput) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Guest{}).
			Where("id = ?", guestID).
			Updates(map[string]interface{}{
				"responded": true,
				"attending": input.Attending,
			}).Error; err != nil {
			return err
		}

		for _, po := range input.PlusOnes {
			if err := tx.Model(&model.PlusOne{}).
				Where("id = ? AND guest_id = ?", po.ID, guestID).
				Update("attending", po.Attending).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// DashboardSummary resume o total de convidados por status.
type DashboardSummary struct {
	Total     int `json:"total"`
	Attending int `json:"attending"`
	Declined  int `json:"declined"`
	Pending   int `json:"pending"`
}

func (r *GuestRepository) Summary() (*DashboardSummary, error) {
	s := &DashboardSummary{}

	var total, attending, declined, pending int64

	if err := r.db.Model(&model.Guest{}).Count(&total).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&model.Guest{}).
		Where("responded = true AND attending = true").
		Count(&attending).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&model.Guest{}).
		Where("responded = true AND attending = false").
		Count(&declined).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&model.Guest{}).
		Where("responded = false").
		Count(&pending).Error; err != nil {
		return nil, err
	}

	s.Total = int(total)
	s.Attending = int(attending)
	s.Declined = int(declined)
	s.Pending = int(pending)

	return s, nil
}
