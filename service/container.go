package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) UpsertContainer(ctx context.Context, input model.NewContainer) (*model.Container, error) {
	return s.repo.UpsertContainer(ctx, input)
}

func (s *Service) ListContainers(ctx context.Context, ids *[]string) ([]*model.Container, error) {
	return s.repo.ListContainers(ctx, ids)
}
