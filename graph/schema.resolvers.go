package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"parts/graph/generated"
	"parts/graph/model"
)

func (r *componentResolver) Tenant(ctx context.Context, obj *model.Component) (*model.Tenant, error) {
	var ids *[]string = &[]string{obj.TenantID}

	tenants, err := r.Svc.ListTenants(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, errors.New("Tenant not found")
	}

	return tenants[0], nil
}

func (r *componentResolver) ComponentType(ctx context.Context, obj *model.Component) (*model.ComponentType, error) {
	var ids *[]string = &[]string{obj.ComponentTypeID}

	componentTypes, err := r.Svc.ListComponentTypes(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(componentTypes) == 0 {
		return nil, errors.New("ComponentType not found")
	}

	return componentTypes[0], nil
}

func (r *componentResolver) Quantities(ctx context.Context, obj *model.Component) ([]*model.Quantity, error) {
	return r.Svc.ListQuantities(ctx, &obj.QuantityIDs)
}

func (r *componentTypeResolver) Tenant(ctx context.Context, obj *model.ComponentType) (*model.Tenant, error) {
	var ids *[]string = &[]string{obj.TenantID}

	tenants, err := r.Svc.ListTenants(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, errors.New("Tenant not found")
	}

	return tenants[0], nil
}

func (r *componentTypeResolver) Components(ctx context.Context, obj *model.ComponentType) ([]*model.Component, error) {
	return r.Svc.ListComponents(ctx, &obj.ComponentIDs)
}

func (r *containerResolver) Tenant(ctx context.Context, obj *model.Container) (*model.Tenant, error) {
	var ids *[]string = &[]string{obj.TenantID}

	tenants, err := r.Svc.ListTenants(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, errors.New("Tenant not found")
	}

	return tenants[0], nil
}

func (r *containerResolver) Parent(ctx context.Context, obj *model.Container) (*model.Container, error) {
	if obj.ParentID == nil {
		return nil, nil
	}

	var ids *[]string = &[]string{*obj.ParentID}

	containers, err := r.Svc.ListContainers(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(containers) == 0 {
		return nil, errors.New("Container not found")
	}

	return containers[0], nil
}

func (r *containerResolver) ContainerType(ctx context.Context, obj *model.Container) (*model.ContainerType, error) {
	var ids *[]string = &[]string{obj.ContainerTypeID}

	containerTypes, err := r.Svc.ListContainerTypes(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(containerTypes) == 0 {
		return nil, errors.New("ContainerType not found")
	}

	return containerTypes[0], nil
}

func (r *containerResolver) Children(ctx context.Context, obj *model.Container) ([]*model.Container, error) {
	return r.Svc.ListContainers(ctx, &obj.ChildIDs)
}

func (r *containerResolver) Quantities(ctx context.Context, obj *model.Container) ([]*model.Quantity, error) {
	return r.Svc.ListQuantities(ctx, &obj.QuantityIDs)
}

func (r *containerTypeResolver) Tenant(ctx context.Context, obj *model.ContainerType) (*model.Tenant, error) {
	var ids *[]string = &[]string{obj.TenantID}

	tenants, err := r.Svc.ListTenants(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, errors.New("Tenant not found")
	}

	return tenants[0], nil
}

func (r *containerTypeResolver) Containers(ctx context.Context, obj *model.ContainerType) ([]*model.Container, error) {
	return r.Svc.ListContainers(ctx, &obj.ContainerIDs)
}

func (r *mutationResolver) UpsertTenant(ctx context.Context, input model.NewTenant) (*model.Tenant, error) {
	return r.Svc.UpsertTenant(ctx, input)
}

func (r *mutationResolver) UpsertUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Svc.UpsertUser(ctx, input)
}

func (r *mutationResolver) UpsertContainerType(ctx context.Context, input model.NewContainerType) (*model.ContainerType, error) {
	return r.Svc.UpsertContainerType(ctx, input)
}

func (r *mutationResolver) UpsertComponentType(ctx context.Context, input model.NewComponentType) (*model.ComponentType, error) {
	return r.Svc.UpsertComponentType(ctx, input)
}

func (r *mutationResolver) UpsertContainer(ctx context.Context, input model.NewContainer) (*model.Container, error) {
	return r.Svc.UpsertContainer(ctx, input)
}

func (r *mutationResolver) UpsertComponent(ctx context.Context, input model.NewComponent) (*model.Component, error) {
	return r.Svc.UpsertComponent(ctx, input)
}

func (r *mutationResolver) UpsertQuantity(ctx context.Context, input model.NewQuantity) (*model.Quantity, error) {
	return r.Svc.UpsertQuantity(ctx, input)
}

func (r *quantityResolver) Container(ctx context.Context, obj *model.Quantity) (*model.Container, error) {
	var ids *[]string = &[]string{obj.ContainerID}

	containers, err := r.Svc.ListContainers(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(containers) == 0 {
		return nil, errors.New("Container not found")
	}

	return containers[0], nil
}

func (r *quantityResolver) Component(ctx context.Context, obj *model.Quantity) (*model.Component, error) {
	var ids *[]string = &[]string{obj.ComponentID}

	components, err := r.Svc.ListComponents(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(components) == 0 {
		return nil, errors.New("Component not found")
	}

	return components[0], nil
}

func (r *queryResolver) Tenants(ctx context.Context, id *string) ([]*model.Tenant, error) {
	var ids *[]string

	if id != nil {
		ids = &[]string{*id}
	}

	return r.Svc.ListTenants(ctx, ids)
}

func (r *queryResolver) Users(ctx context.Context, id *string) ([]*model.User, error) {
	var ids *[]string

	if id != nil {
		ids = &[]string{*id}
	}
	return r.Svc.ListUsers(ctx, ids)
}

func (r *queryResolver) ContainerTypes(ctx context.Context, id *string) ([]*model.ContainerType, error) {
	var ids *[]string

	if id != nil {
		ids = &[]string{*id}
	}
	return r.Svc.ListContainerTypes(ctx, ids)
}

func (r *queryResolver) ComponentTypes(ctx context.Context, id *string) ([]*model.ComponentType, error) {
	var ids *[]string

	if id != nil {
		ids = &[]string{*id}
	}
	return r.Svc.ListComponentTypes(ctx, ids)
}

func (r *queryResolver) Containers(ctx context.Context, id *string) ([]*model.Container, error) {
	var ids *[]string

	if id != nil {
		ids = &[]string{*id}
	}
	return r.Svc.ListContainers(ctx, ids)
}

func (r *queryResolver) Components(ctx context.Context, id *string) ([]*model.Component, error) {
	var ids *[]string

	if id != nil {
		ids = &[]string{*id}
	}
	return r.Svc.ListComponents(ctx, ids)
}

func (r *queryResolver) Quantities(ctx context.Context, id *string) ([]*model.Quantity, error) {
	var ids *[]string

	if id != nil {
		ids = &[]string{*id}
	}
	return r.Svc.ListQuantities(ctx, ids)
}

func (r *tenantResolver) Users(ctx context.Context, obj *model.Tenant) ([]*model.User, error) {
	return r.Svc.ListUsers(ctx, &obj.UserIDs)
}

func (r *tenantResolver) ContainerTypes(ctx context.Context, obj *model.Tenant) ([]*model.ContainerType, error) {
	return r.Svc.ListContainerTypes(ctx, &obj.ContainerTypeIDs)
}

func (r *tenantResolver) ComponentTypes(ctx context.Context, obj *model.Tenant) ([]*model.ComponentType, error) {
	return r.Svc.ListComponentTypes(ctx, &obj.ComponentTypeIDs)
}

func (r *userResolver) Tenant(ctx context.Context, obj *model.User) (*model.Tenant, error) {
	var ids *[]string = &[]string{obj.TenantID}

	tenants, err := r.Svc.ListTenants(ctx, ids)

	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, errors.New("Tenant not found")
	}

	return tenants[0], nil
}

// Component returns generated.ComponentResolver implementation.
func (r *Resolver) Component() generated.ComponentResolver { return &componentResolver{r} }

// ComponentType returns generated.ComponentTypeResolver implementation.
func (r *Resolver) ComponentType() generated.ComponentTypeResolver { return &componentTypeResolver{r} }

// Container returns generated.ContainerResolver implementation.
func (r *Resolver) Container() generated.ContainerResolver { return &containerResolver{r} }

// ContainerType returns generated.ContainerTypeResolver implementation.
func (r *Resolver) ContainerType() generated.ContainerTypeResolver { return &containerTypeResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Quantity returns generated.QuantityResolver implementation.
func (r *Resolver) Quantity() generated.QuantityResolver { return &quantityResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tenant returns generated.TenantResolver implementation.
func (r *Resolver) Tenant() generated.TenantResolver { return &tenantResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type componentResolver struct{ *Resolver }
type componentTypeResolver struct{ *Resolver }
type containerResolver struct{ *Resolver }
type containerTypeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type quantityResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tenantResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
