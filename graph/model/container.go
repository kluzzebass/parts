package model

import "time"

type Container struct {
	ID              string    `json:"id" db:"container_id"`
	TenantID        string    `json:"tenant" db:"tenant_id"`
	ParentID        *string   `json:"parentId" db:"parent_container_id"`
	ContainerTypeID string    `json:"containerType" db:"container_type_id"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	Description     string    `json:"description" db:"description"`
	ChildIDs        []string  `json:"childIds" db:"children"`
	QuantityIDs     []string  `json:"quantityIds" db:"quantities"`
}
