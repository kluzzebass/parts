package service

import (
	"parts/graph/model"
)

func (s *Service) CreateUser(nu model.NewUser) (*model.User, error) {
	return s.repo.CreateUser(nu)
}

func (s *Service) ListUsers(ids *[]string) ([]*model.User, error) {
	return s.repo.ListUsers(ids)
}
