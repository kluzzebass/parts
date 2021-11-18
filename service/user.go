package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) UpsertUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return s.repo.UpsertUser(ctx, input)
}

func (s *Service) ListUsers(ctx context.Context, ids *[]string, sort []*model.UserSort) ([]*model.User, error) {
	return s.repo.ListUsers(ctx, ids, sort)
}
