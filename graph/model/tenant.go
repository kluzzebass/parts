package model

type Tenant struct {
	ID               string   `json:"id"`
	CreatedAt        string   `json:"createdAt"`
	Name             string   `json:"name"`
	UserIDs          []string `json:"userIds"`
	ContainerTypeIDs []string `json:"containerTypeIds"`
}
