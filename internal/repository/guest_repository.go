package repository

import (
	"database/sql"
	"errors"

	"github.com/lucas/confirmation-mariage-app/internal/model"
)

type GuestRepository struct {
	db *sql.DB
}

func NewGuestRepository(db *sql.DB) *GuestRepository {
	return &GuestRepository{db: db}
}

func (r *GuestRepository) FindAll() ([]model.Guest, error) {
	rows, err := r.db.Query(`SELECT id, name, responded, attending FROM guests ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guests []model.Guest
	for rows.Next() {
		var g model.Guest
		if err := rows.Scan(&g.ID, &g.Name, &g.Responded, &g.Attending); err != nil {
			return nil, err
		}

		plusOnes, err := r.findPlusOnesByGuestID(g.ID)
		if err != nil {
			return nil, err
		}
		g.PlusOnes = plusOnes

		guests = append(guests, g)
	}
	return guests, nil
}

func (r *GuestRepository) FindByID(id int) (*model.Guest, error) {
	g := &model.Guest{}
	row := r.db.QueryRow(`SELECT id, name, responded, attending FROM guests WHERE id = $1`, id)
	err := row.Scan(&g.ID, &g.Name, &g.Responded, &g.Attending)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	plusOnes, err := r.findPlusOnesByGuestID(g.ID)
	if err != nil {
		return nil, err
	}
	g.PlusOnes = plusOnes

	return g, nil
}

func (r *GuestRepository) SearchByName(name string) ([]model.Guest, error) {
	rows, err := r.db.Query(
		`SELECT id, name, responded, attending FROM guests WHERE name ILIKE '%' || $1 || '%' ORDER BY name`,
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guests []model.Guest
	for rows.Next() {
		var g model.Guest
		if err := rows.Scan(&g.ID, &g.Name, &g.Responded, &g.Attending); err != nil {
			return nil, err
		}

		plusOnes, err := r.findPlusOnesByGuestID(g.ID)
		if err != nil {
			return nil, err
		}
		g.PlusOnes = plusOnes

		guests = append(guests, g)
	}
	return guests, nil
}

// Confirm atualiza o status do convidado principal e de cada agregado,
// dentro de uma transação para garantir atomicidade.
func (r *GuestRepository) Confirm(guestID int, input model.ConfirmInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		`UPDATE guests SET responded = true, attending = $1 WHERE id = $2`,
		input.Attending, guestID,
	)
	if err != nil {
		return err
	}

	for _, po := range input.PlusOnes {
		_, err = tx.Exec(
			`UPDATE plus_ones SET attending = $1 WHERE id = $2 AND guest_id = $3`,
			po.Attending, po.ID, guestID,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *GuestRepository) findPlusOnesByGuestID(guestID int) ([]model.PlusOne, error) {
	rows, err := r.db.Query(
		`SELECT id, guest_id, name, attending FROM plus_ones WHERE guest_id = $1`,
		guestID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plusOnes []model.PlusOne
	for rows.Next() {
		var p model.PlusOne
		if err := rows.Scan(&p.ID, &p.GuestID, &p.Name, &p.Attending); err != nil {
			return nil, err
		}
		plusOnes = append(plusOnes, p)
	}
	return plusOnes, nil
}

// Dashboard summary
type DashboardSummary struct {
	Total     int `json:"total"`
	Attending int `json:"attending"`
	Declined  int `json:"declined"`
	Pending   int `json:"pending"`
}

func (r *GuestRepository) Summary() (*DashboardSummary, error) {
	row := r.db.QueryRow(`
		SELECT
			COUNT(*) AS total,
			COUNT(*) FILTER (WHERE responded = true AND attending = true) AS attending,
			COUNT(*) FILTER (WHERE responded = true AND attending = false) AS declined,
			COUNT(*) FILTER (WHERE responded = false) AS pending
		FROM guests
	`)

	s := &DashboardSummary{}
	err := row.Scan(&s.Total, &s.Attending, &s.Declined, &s.Pending)
	return s, err
}
