package service

import (
	"parts/graph/model"
)

func (s *Service) CreateTenant(nt model.NewTenant) (*model.Tenant, error) {
	return s.repo.CreateTenant(nt)
}

func (s *Service) ListTenants(id *string) ([]*model.Tenant, error) {
	return s.repo.ListTenants(id)
}
