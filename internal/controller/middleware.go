package controller

import "net/http"

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct { 
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w, r); err != nil {
			WriteJSONtoHTTP(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}