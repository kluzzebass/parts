package model

type User struct {
	ID        string `json:"id"`
	TenantID  string `json:"tenantId"`
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
}
