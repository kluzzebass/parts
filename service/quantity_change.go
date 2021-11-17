package service

import (
	"context"
	"parts/graph/model"
)

func (s *Service) ListQuantityChanges(ctx context.Context, ids *[]string) ([]*model.QuantityChange, error) {
	return s.repo.ListQuantityChanges(ctx, ids)
}
