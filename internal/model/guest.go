package model

type Guest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Responded bool      `json:"responded"`
	Attending bool      `json:"attending"`
	PlusOnes  []PlusOne `json:"plus_ones,omitempty"`
}

type PlusOne struct {
	ID        int    `json:"id"`
	GuestID   int    `json:"guest_id"`
	Name      string `json:"name"`
	Attending bool   `json:"attending"`
}
type PlusOneConfirmation struct {
	ID        int  `json:"id"`
	Attending bool `json:"attending"`
}

// ConfirmInput é o payload enviado pelo convidado ao confirmar presença.
type ConfirmInput struct {
	Attending bool                  `json:"attending"`
	PlusOnes  []PlusOneConfirmation `json:"plus_ones,omitempty"`
}
