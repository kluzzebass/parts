package model

import "time"

type Quantity struct {
	ID          string    `json:"id" db:"quantity_id"`
	ContainerID string    `json:"containerId" db:"container_id"`
	ComponentID string    `json:"componentId" db:"component_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	Quantity    int       `json:"quantity" db:"quantity"`
}
