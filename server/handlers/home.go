package handler

import (
	"errors"
	"net/http"
)

func Home(env *Env, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		w.Write([]byte("Welcome to the Budgetly API"))
    } else {
		return StatusError{400, errors.New("invalid HTTP method/URL pair")}
	}

	return nil
}
