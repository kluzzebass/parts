package model

import "time"

type Tenant struct {
	ID               string    `json:"id" db:"tenant_id"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
	Name             string    `json:"name" db:"name"`
	UserIDs          []string  `json:"userIds" db:"users"`
	ContainerTypeIDs []string  `json:"containerTypeIds" db:"container_types"`
	ComponentTypeIDs []string  `json:"componentTypeIds" db:"component_types"`
}
