package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"github.com/infinity-ocean/goldconv/internal/config"
	"github.com/infinity-ocean/goldconv/internal/model"
)

type CtxKey string

type repo interface {
	InsertAccount(small model.AccountSmall) error	
	Login(acc model.AccountLogin) (string, int, error) 
	SelectAccount(id int) (model.Account, error)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct { 
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc { 	// 
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSONtoHTTP(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func withJWTAuth(handler http.HandlerFunc, repo repo, conf config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secret := conf.JWTSecret
		tokenString := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenString, "Bearer ") {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("jwt token in 'Authorization' header is not provided"))
		}
		tokenTrim := strings.TrimPrefix(tokenString, "Bearer ")
		
		token, err := jwt.Parse(tokenTrim, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			} 
			return []byte(secret), nil
		}) 

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("access restricted"))
			return 
		}

		if !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("token is not valid"))
		}
		id, err := getID(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("id in http request is not correct"))
		}
		account, err := repo.SelectAccount(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("pg didn't return account"))
		}
		claims := token.Claims.(jwt.MapClaims) // type assertion, проверять ok
		if account.ID != int(claims["accountID"].(float64)) { // ?
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("id from db and id in token aren't match"))
		}
		// r = r.WithContext(context.WithValue(r.Context(), CtxKey{}, account)) // не знаю как зашивать в контекст
		handler(w, r)
	}
}


func WriteJSONtoHTTP(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}