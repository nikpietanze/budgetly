package handler

import (
	"budgetly/models"

	"encoding/json"
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

type ModelData interface {
	[]*models.User |
	[]*models.Account |
	[]*models.Expense |
	[]*models.Revenue
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

func ReturnJson[T ModelData](m T, w http.ResponseWriter) (bool, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return false, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return true, nil
}
