package service

import (
	"github.com/lucas/confirmation-mariage-app/internal/model"
	"github.com/lucas/confirmation-mariage-app/internal/repository"
)

type GuestService struct {
	repo *repository.GuestRepository
}

func NewGuestService(repo *repository.GuestRepository) *GuestService {
	return &GuestService{repo: repo}
}

func (s *GuestService) ListAll() ([]model.Guest, error) {
	return s.repo.FindAll()
}

func (s *GuestService) GetByID(id int) (*model.Guest, error) {
	return s.repo.FindByID(id)
}

func (s *GuestService) SearchByName(name string) ([]model.Guest, error) {
	return s.repo.SearchByName(name)
}

func (s *GuestService) Dashboard() (*repository.DashboardSummary, error) {
	return s.repo.Summary()
}

func (s *GuestService) Confirm(id int, input model.ConfirmInput) error {
	// delega tudo ao repository, que já faz transação e atualiza convidados + acompanhantes
	return s.repo.Confirm(id, input)
}
