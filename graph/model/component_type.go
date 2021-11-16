package model

import "time"

type ComponentType struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenantId"`
	CreatedAt    time.Time `json:"createdAt"`
	Description  string    `json:"description"`
	ComponentIDs []string  `json:"componentIds"`
}
