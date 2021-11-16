package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) CreateUser(ctx context.Context, nu model.NewUser) (*model.User, error) {
	return s.repo.CreateUser(ctx, nu)
}

func (s *Service) ListUsers(ctx context.Context, ids *[]string) ([]*model.User, error) {
	return s.repo.ListUsers(ctx, ids)
}
