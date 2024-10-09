package controller

import (

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
	router.HandleFunc("/goldconv/login", makeHTTPHandleFunc(c.handleAccount)) // POST
	router.HandleFunc("/goldconv/account", makeHTTPHandleFunc(c.handleAccount)) // POST
	router.HandleFunc("/goldconv/account/{id}", withJWTAuth(makeHTTPHandleFunc(c.handleBalance))) // GET
	fmt.Println("Starting server on ", c.listenPort)
	if err := http.ListenAndServe(c.listenPort, router); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
