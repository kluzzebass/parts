package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) UpsertComponentType(ctx context.Context, input model.NewComponentType) (*model.ComponentType, error) {
	return s.repo.UpsertComponentType(ctx, input)
}

func (s *Service) ListComponentTypes(ctx context.Context, ids *[]string) ([]*model.ComponentType, error) {
	return s.repo.ListComponentTypes(ctx, ids)
}
