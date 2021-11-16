package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) CreateTenant(ctx context.Context, nt model.NewTenant) (*model.Tenant, error) {
	return s.repo.CreateTenant(ctx, nt)
}

func (s *Service) ListTenants(ctx context.Context, ids *[]string) ([]*model.Tenant, error) {
	return s.repo.ListTenants(ctx, ids)
}
