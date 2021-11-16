package model

import "time"

type Quantity struct {
	ID          string    `json:"id"`
	ContainerID string    `json:"containerId"`
	ComponentID string    `json:"componentId"`
	CreatedAt   time.Time `json:"createdAt"`
	Quantity    int       `json:"quantity"`
}
