package service

import (
	"fmt"

	"github.com/infinity-ocean/goldconv/internal/model"
)

func (s *service) AddBalance(id int, balance model.Balance) error{
	err := s.repo.InsertBalance(id, balance)
	if err != nil {
		return fmt.Errorf("service: insertBalance returned error: %w", err)
	}
	return nil
}