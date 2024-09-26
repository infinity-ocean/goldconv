package controller

import (
	"fmt"
	"net/http"
)

func (c *controller) ShowBalance(w http.ResponseWriter, r *http.Request) error {
	id := 1 
	b, err := c.service.GetBalance(id) 
	if err != nil {
		fmt.Println("Service didn't send balance:", err)
		return err
	}
	response := fmt.Sprintf("Your balance. Gold: %d, Silver: %d, Bronze: %d", b.Gold, b.Silver, b.Bronze)
	w.Write([]byte(response))
	return nil
}