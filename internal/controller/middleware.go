package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct { 
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc { 	// TODO JWT, user creation
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSONtoHTTP(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func withJWTAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// authorisation: token creation
		// authentication: token checking
		// tok := token{}
		// tok.key = []byte("hello")
		// tok.t = jwt.New(jwt.SigningMethodHS256) 
		// tok.s = tok.t.SignedString(tok.key) 

		handler(w, r)
	}
}

func validateJWT(tokenString string) (jwt.Token, error){

}

// type token struct {
// 	key []byte
//   	t   *jwt.Token
//   	s   string
// }