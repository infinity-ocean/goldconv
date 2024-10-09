package controller

import (
	"net/http"
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
		handler(w, r)
	}
}
