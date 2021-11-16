package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) UpsertQuantity(ctx context.Context, input model.NewQuantity) (*model.Quantity, error) {
	return s.repo.UpsertQuantity(ctx, input)
}

func (s *Service) ListQuantities(ctx context.Context, ids *[]string) ([]*model.Quantity, error) {
	return s.repo.ListQuantities(ctx, ids)
}
