package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"parts/graph/generated"
	"parts/graph/model"
)

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

func (r *mutationResolver) CreateTenant(ctx context.Context, input model.NewTenant) (*model.Tenant, error) {
	return r.Svc.CreateTenant(ctx, input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Svc.CreateUser(ctx, input)
}

func (r *mutationResolver) CreateContainerType(ctx context.Context, input model.NewContainerType) (*model.ContainerType, error) {
	return r.Svc.CreateContainerType(ctx, input)
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

func (r *tenantResolver) Users(ctx context.Context, obj *model.Tenant) ([]*model.User, error) {
	return r.Svc.ListUsers(ctx, &obj.UserIDs)
}

func (r *tenantResolver) ContainerTypes(ctx context.Context, obj *model.Tenant) ([]*model.ContainerType, error) {
	return r.Svc.ListContainerTypes(ctx, &obj.ContainerTypeIDs)
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

// ContainerType returns generated.ContainerTypeResolver implementation.
func (r *Resolver) ContainerType() generated.ContainerTypeResolver { return &containerTypeResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tenant returns generated.TenantResolver implementation.
func (r *Resolver) Tenant() generated.TenantResolver { return &tenantResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type containerTypeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tenantResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
