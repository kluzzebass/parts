package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"parts/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Svc *service.Service
	// tenants []*model.Tenant
	// users   []*model.User
	// containerTypes []*model.ContainerType
	// containers     []*model.Container
	// componentTypes []*model.ComponentType
	// components     []*model.Component
}
