package controller

import (
	"github.com/infinity-ocean/goldconv/internal/model"

)

type controller struct {
	service service
}

type service interface {
	GetBalance(id int) (model.Balance, error)
}

func NewController (svc service) *controller {
	return &controller{service: svc}
}


