package service

import (
	"fmt"

	"github.com/infinity-ocean/goldconv/internal/model"

)

func (s *service) GetBalance(id int) (model.Balance, error) {
	b, err := s.repo.SelectBalance(id)
	if err != nil {
		return b, fmt.Errorf("service: selectBalance returned error: %w", err)
	}
	return b, nil
}