package model

import "time"

type Container struct {
	ID              string    `json:"id"`
	TenantID        string    `json:"tenant"`
	ParentID        *string   `json:"parentId"`
	ContainerTypeID string    `json:"containerType"`
	CreatedAt       time.Time `json:"createdAt"`
	Description     string    `json:"description"`
	ChildIDs        []string  `json:"childIds"`
	QuantityIDs     []string  `json:"quantityIds"`
}
