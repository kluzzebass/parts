package model

import "time"

type ContainerType struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenantId"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
}
