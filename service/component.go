package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) UpsertComponent(ctx context.Context, input model.NewComponent) (*model.Component, error) {
	return s.repo.UpsertComponent(ctx, input)
}

func (s *Service) ListComponents(ctx context.Context, ids *[]string) ([]*model.Component, error) {
	return s.repo.ListComponents(ctx, ids)
}
