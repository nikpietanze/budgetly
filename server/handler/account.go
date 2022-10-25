package handler

import (
	"budgetly/models"

	"encoding/json"
	"net/http"
)

func Account(env *Env, w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		accounts := []models.Account{}
        err := env.DB.Model(&models.Account{}).Preload("Revenue").Preload("Expenses").Find(&accounts).Error
        if err != nil {
            return StatusError{500, err}
        }

		data, err := json.Marshal(accounts)
		if err != nil {
		    return StatusError{500, err}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case http.MethodPost:
		account := &models.Account{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(account)
		if err != nil {
            return StatusError{500, err}
		}
		env.DB.Create(account)

		data, err := json.Marshal(account)
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
