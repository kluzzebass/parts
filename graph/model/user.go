package model

import "time"

type User struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenantId"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}
