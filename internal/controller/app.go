package controller

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/infinity-ocean/goldconv/internal/model"
)

type controller struct {
	service service
	listenPort string
}

type service interface {
	GetBalance(int) (model.Balance, error)
	AddBalance(int, model.Balance) error
}

func NewController(svc service, port string) *controller {
	return &controller{service: svc, listenPort: port}
}

func (c *controller) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/goldconv/balance", withJWTAuth(makeHTTPHandleFunc(c.handleBalance)))
	router.HandleFunc("/goldconv/balance/{id}", withJWTAuth(makeHTTPHandleFunc(c.handleBalance)))
	fmt.Println("Starting server on ", c.listenPort)
	if err := http.ListenAndServe(c.listenPort, router); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
			
func (c *controller) handleBalance(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return c.AddBalance(w, r)
	}
	if r.Method == "GET" {
		return c.ShowBalance(w, r)
	}
	if r.Method == "DELETE" {
		return nil
	}
	return nil
}

func WriteJSONtoHTTP(w http.ResponseWriter, status int, v any) error { // вроде разобался
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}