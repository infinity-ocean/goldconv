package service

import (
	"github.com/infinity-ocean/goldconv/internal/model"

)

type service struct {
	repo repo
}

type repo interface {
	SelectBalance(id int) (model.Balance, error)
}

func NewService(repo repo) *service {
	return &service{repo: repo}
}
