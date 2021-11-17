package model

import "time"

type ContainerType struct {
	ID           string    `json:"id" db:"container_type_id"`
	TenantID     string    `json:"tenantId" db:"tenant_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	Description  string    `json:"description" db:"description"`
	ContainerIDs []string  `json:"containerIds" db:"containers"`
}
