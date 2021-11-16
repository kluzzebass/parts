package model

import "time"

type Tenant struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	Name             string    `json:"name"`
	UserIDs          []string  `json:"userIds"`
	ContainerTypeIDs []string  `json:"containerTypeIds"`
	ComponentTypeIDs []string  `json:"componentTypeIds"`
}
