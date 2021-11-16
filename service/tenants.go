package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) UpsertTenant(ctx context.Context, input model.NewTenant) (*model.Tenant, error) {
	return s.repo.UpsertTenant(ctx, input)
}

func (s *Service) ListTenants(ctx context.Context, ids *[]string) ([]*model.Tenant, error) {
	return s.repo.ListTenants(ctx, ids)
}
