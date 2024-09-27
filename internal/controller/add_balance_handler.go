package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



func (c *controller) AddBalance(w http.ResponseWriter, r *http.Request) error {
	var mapBalance map[string]uint
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&mapBalance); err != nil{
		return fmt.Errorf("controller: JSON not fits balance: %w", err)
	}
	if len(mapBalance) != 3{
		return errors.New("controller: JSON must contain exactly 3 fields (Gold, Silver, Bronze)")
	}
	var balance balance
	balance.Gold = mapBalance["Gold"]
	balance.Silver = uint8(mapBalance["Silver"])
	balance.Bronze = uint8(mapBalance["Bronze"])
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		return fmt.Errorf("controller: invalid user ID: %w", err)
	}

	if err := c.service.AddBalance(id, balance.toModel()); err != nil{
		return fmt.Errorf("controller: failed to add balance: %w", err)
	}
	return nil
}