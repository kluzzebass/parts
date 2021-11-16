package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) CreateComponentType(ctx context.Context, nt model.NewComponentType) (*model.ComponentType, error) {
	return s.repo.CreateComponentType(ctx, nt)
}

func (s *Service) ListComponentTypes(ctx context.Context, ids *[]string) ([]*model.ComponentType, error) {
	return s.repo.ListComponentTypes(ctx, ids)
}
