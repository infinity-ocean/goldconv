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
	InsertAccount(model.AccountSmall) error
	Login(model.AccountLogin) (string, error)
}

func NewService(repo repo) *service {
	return &service{repo: repo}
}

func (s *service) AddAccount(small model.AccountSmall) error {
	return s.repo.InsertAccount(small)
}

func (s *service) Login(loginAcc model.AccountLogin) (string, error) {
	return s.repo.Login(loginAcc)
}