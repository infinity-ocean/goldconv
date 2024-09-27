package controller

import "github.com/infinity-ocean/goldconv/internal/model"

type balance struct {
	Gold   uint  `json:"gold"`
	Silver uint8 `json:"silver"`
	Bronze uint8 `json:"bronze"`
}

func (b *balance) toModel() model.Balance {
	return model.Balance{
		Gold:   b.Gold,
		Silver: b.Silver,
		Bronze: b.Bronze,
	}
}
