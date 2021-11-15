package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"parts/graph/generated"
	"parts/graph/model"
)

func (r *mutationResolver) CreateTenant(ctx context.Context, input model.NewTenant) (*model.Tenant, error) {
	return r.Svc.CreateTenant(input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Svc.CreateUser(input)
}

func (r *queryResolver) Tenants(ctx context.Context, id *string) ([]*model.Tenant, error) {
	return r.Svc.ListTenants(id)
}

func (r *queryResolver) Users(ctx context.Context, id *string) ([]*model.User, error) {
	var ids *[]string

	if id != nil {
		ids = &[]string{*id}
	}
	return r.Svc.ListUsers(ids)
}

func (r *tenantResolver) Users(ctx context.Context, obj *model.Tenant) ([]*model.User, error) {
	fmt.Println(obj)

	users, err := r.Svc.ListUsers(&obj.UserIDs)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userResolver) Tenant(ctx context.Context, obj *model.User) (*model.Tenant, error) {
	tenantId := obj.TenantID
	tenants, err := r.Svc.ListTenants(&tenantId)

	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, errors.New("Tenant not found")
	}

	return tenants[0], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tenant returns generated.TenantResolver implementation.
func (r *Resolver) Tenant() generated.TenantResolver { return &tenantResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tenantResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
