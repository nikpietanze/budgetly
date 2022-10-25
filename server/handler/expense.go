package handler

import (
	"budgetly/models"

	"encoding/json"
	"net/http"
)

func Expense(env *Env, w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		expenses := []models.Expense{}
		env.DB.Find(&expenses)

		data, err := json.Marshal(expenses)
		if err != nil {
			return StatusError{500, err}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case http.MethodPost:
		expense := &models.Expense{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(expense)
		if err != nil {
			return StatusError{500, err}
		}
		env.DB.Create(expense)

		data, err := json.Marshal(expense)
		if err != nil {
			return StatusError{500, err}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case http.MethodPut:
		break
	case http.MethodDelete:
		break
	}
	return nil
}
