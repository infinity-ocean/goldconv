package service

import (
	"github.com/infinity-ocean/goldconv/internal/model"
)

type service struct {
	repo repo
}

type repo interface {
	InsertAccount(model.AccountSmall) error
	Login(model.AccountLogin) (string, int, error)
	SelectAccount(int) (model.Account, error)
}

func NewService(repo repo) *service {
	return &service{repo: repo}
}

func (s *service) AddAccount(small model.AccountSmall) error {
	return s.repo.InsertAccount(small)
}

func (s *service) Login(loginAcc model.AccountLogin) (string, int, error) {
	return s.repo.Login(loginAcc)
}

func (s *service) GetAccount(id int) (model.Account, error) {
	return s.repo.SelectAccount(id)
}