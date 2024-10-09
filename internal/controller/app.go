package controller

import (
	"encoding/json"
	"errors"
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
	AddAccount(model.AccountSmall) error
}

func NewController(svc service, port string) *controller {
	return &controller{service: svc, listenPort: port}
}

func (c *controller) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/goldconv/account", makeHTTPHandleFunc(c.handleAccount)) // POST [Make account]
	//TODO router.HandleFunc("/goldconv/login", makeHTTPHandleFunc(c.handleLogin)) // POST [Send JWT]
	//TODO router.HandleFunc("/goldconv/account/{id}", withJWTAuth(makeHTTPHandleFunc(c.handleAccountWithID))) // GET [Get account]
	fmt.Println("Starting server on ", c.listenPort)
	if err := http.ListenAndServe(c.listenPort, router); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}

func (c *controller) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("method for account creation isn't POST")
	}
	req := new(accountRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	if req.Username == "" || req.Email == "" || req.Password == "" || req.Balance == "" {
		return errors.New("one or more fields are empty")
	}
	acc := model.AccountSmall{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
		Balance: req.Balance,
	}
	err := c.service.AddAccount(acc)
	if err != nil {
		return err
	}
	return WriteJSONtoHTTP(w, http.StatusOK, acc)
	}

// 	func toID(r *http.Request) (int, error) {
// 		id := mux.Vars(r)["id"]
// 		return strconv.Atoi(id)
// }