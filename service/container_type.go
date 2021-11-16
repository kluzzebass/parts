package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) UpsertContainerType(ctx context.Context, input model.NewContainerType) (*model.ContainerType, error) {
	return s.repo.UpsertContainerType(ctx, input)
}

func (s *Service) ListContainerTypes(ctx context.Context, ids *[]string) ([]*model.ContainerType, error) {
	return s.repo.ListContainerTypes(ctx, ids)
}
