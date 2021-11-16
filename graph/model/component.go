package model

import "time"

type Component struct {
	ID              string    `json:"id"`
	TenantID        string    `json:"tenantId"`
	ComponentTypeID string    `json:"componentTypeId"`
	CreatedAt       time.Time `json:"createdAt"`
	Description     string    `json:"description"`
	QuantityIDs     []string  `json:"quantityIds"`
}
