package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/infinity-ocean/goldconv/internal/model"
)

type controller struct {
	service service
	listenPort string
}

type service interface {
	AddAccount(model.AccountSmall) error
	Login(model.AccountLogin) (string, int, error)
	GetAccount(int) (model.Account, error)
}

func NewController(svc service, port string) *controller {
	return &controller{service: svc, listenPort: port}
}

func (c *controller) Run(repo repo) {
	router := mux.NewRouter()
	router.HandleFunc("/goldconv/account", makeHTTPHandleFunc(c.handleAccount)) // POST [Make account]
	router.HandleFunc("/goldconv/login", makeHTTPHandleFunc(c.handleLogin)) // POST [Send JWT] 
	router.HandleFunc("/goldconv/account/{id}", withJWTAuth(makeHTTPHandleFunc(c.handleAccountWithID), repo)) // GET [Get account] With JWT auth
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

func (c *controller) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("method for account creation isn't POST")
	}
	req := new(AccountLogin)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	if req.Username == "" || req.Password == "" {
		return errors.New("one or more fields are empty")
	}
	acc := model.AccountLogin{
		Username: req.Username,
		Password: req.Password,
	}
	jwt, id, err := c.service.Login(acc)
	if err != nil {
		return err
	} 
	answer := make(map[string]string)
	answer["jwt"] = jwt
	answer["id"] = strconv.Itoa(id)
	return WriteJSONtoHTTP(w, http.StatusOK, answer)
}

func (c *controller) handleAccountWithID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	account, err := c.service.GetAccount(id)
	if err != nil {
		return err
	}
	err = WriteJSONtoHTTP(w, http.StatusOK, account) // PAY ATTENTION, MAY BE ERROR
	if err != nil {
		return fmt.Errorf("failed to write account into http: %v", err)
	}
	return nil
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}