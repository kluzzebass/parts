package model

import "time"

type ComponentType struct {
	ID           string    `json:"id" db:"component_type_id"`
	TenantID     string    `json:"tenantId" db:"tenant_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	Description  string    `json:"description" db:"description"`
	ComponentIDs []string  `json:"componentIds" db:"components"`
}
