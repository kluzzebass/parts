package model

import "time"

type Component struct {
	ID              string    `json:"id" db:"component_id"`
	TenantID        string    `json:"tenantId" db:"tenant_id"`
	ComponentTypeID string    `json:"componentTypeId" db:"component_type_id"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	Description     string    `json:"description" db:"description"`
	QuantityIDs     []string  `json:"quantityIds" db:"quantities"`
}
