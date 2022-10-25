package handler

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (statusError StatusError) Error() string {
	return statusError.Err.Error()
}

func (statusError StatusError) Status() int {
	return statusError.Code
}

type Env struct {
	DB   *gorm.DB
	Port string
	Host string
}

type Handler struct {
	Env    *Env
	Handle func(e *Env, w http.ResponseWriter, r *http.Request) error
}

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := handler.Handle(handler.Env, w, r)
	if err != nil {
		switch error := err.(type) {
		case Error:
			fmt.Printf("HTTP %d - %s", error.Status(), error)
			http.Error(w, error.Error(), error.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
