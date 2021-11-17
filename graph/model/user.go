package model

import "time"

type User struct {
	ID        string    `json:"id" db:"user_id"`
	TenantID  string    `json:"tenantId" db:"tenant_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	Name      string    `json:"name" db:"name"`
}
