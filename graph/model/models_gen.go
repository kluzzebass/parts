// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewComponent struct {
	ID              *string `json:"id"`
	TenantID        string  `json:"tenantId"`
	ComponentTypeID string  `json:"componentTypeId"`
	Description     string  `json:"description"`
}

type NewComponentType struct {
	ID          *string `json:"id"`
	TenantID    string  `json:"tenantId"`
	Description string  `json:"description"`
}

type NewContainer struct {
	ID              *string `json:"id"`
	TenantID        string  `json:"tenantId"`
	ParentID        *string `json:"parentId"`
	ContainerTypeID string  `json:"containerTypeId"`
	Description     string  `json:"description"`
}

type NewContainerType struct {
	ID          *string `json:"id"`
	TenantID    string  `json:"tenantId"`
	Description string  `json:"description"`
}

type NewQuantity struct {
	ID          *string `json:"id"`
	ContainerID string  `json:"containerId"`
	ComponentID string  `json:"componentId"`
	Quantity    int     `json:"quantity"`
}

type NewTenant struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
}

type NewUser struct {
	ID       *string `json:"id"`
	TenantID string  `json:"tenantId"`
	Name     string  `json:"name"`
}
