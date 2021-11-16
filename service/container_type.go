package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) CreateContainerType(ctx context.Context, nt model.NewContainerType) (*model.ContainerType, error) {
	return s.repo.CreateContainerType(ctx, nt)
}

func (s *Service) ListContainerTypes(ctx context.Context, ids *[]string) ([]*model.ContainerType, error) {
	return s.repo.ListContainerTypes(ctx, ids)
}
