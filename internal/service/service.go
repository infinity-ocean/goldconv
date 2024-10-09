package service

import (
	"github.com/infinity-ocean/goldconv/internal/model"

)

type service struct {
	repo repo
}

type repo interface {
	SelectBalance(int) (model.Balance, error)
	InsertBalance(int, model.Balance) error
	InsertAccount(model.AccountSmall) error // ? should we return pointer or not
}

func NewService(repo repo) *service {
	return &service{repo: repo}
}

func (s *service) AddAccount(small model.AccountSmall) error {
	return s.repo.InsertAccount(small)
}