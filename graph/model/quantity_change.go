package model

import "time"

type QuantityChange struct {
	ID         string    `json:"id" db:"quantity_change_id"`
	QuantityID string    `json:"quantityId" db:"quantity_id"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	Amount     int       `json:"amount" db:"amount"`
}
