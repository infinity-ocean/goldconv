package service

import (
	"fmt"

	"github.com/infinity-ocean/goldconv/internal/model"

)

func (s *service) GetBalance(id int) (model.Balance, error) {
	b, err := s.repo.SelectBalance(id)
	if err != nil {
		fmt.Println("Repo returned error:", err)
		return b, err
	}
	return b, nil
}