package model

type Guest struct {
	ID        int       `json:"id"                  gorm:"primaryKey"`
	Name      string    `json:"name"                gorm:"not null"`
	Responded bool      `json:"responded"           gorm:"default:false"`
	Attending bool      `json:"attending"           gorm:"default:false"`
	PlusOnes  []PlusOne `json:"plus_ones,omitempty" gorm:"foreignKey:GuestID"`
}

type PlusOne struct {
	ID        int    `json:"id"        gorm:"primaryKey"`
	GuestID   int    `json:"guest_id"  gorm:"not null;index"`
	Name      string `json:"name"      gorm:"not null"`
	Attending bool   `json:"attending" gorm:"default:false"`
}

// PlusOneConfirmation é o payload de um acompanhante dentro da confirmação
// enviada pelo convidado (não é uma tabela, só DTO de entrada).
type PlusOneConfirmation struct {
	ID        int  `json:"id"`
	Attending bool `json:"attending"`
}

// ConfirmInput é o payload enviado pelo convidado ao confirmar presença.
type ConfirmInput struct {
	Attending bool                  `json:"attending"`
	PlusOnes  []PlusOneConfirmation `json:"plus_ones,omitempty"`
}
