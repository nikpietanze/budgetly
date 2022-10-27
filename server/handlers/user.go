package handler

import (
	"budgetly/models"

	"encoding/json"
	"net/http"
)

func User(env *Env, w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		users := []models.User{}
        err := env.DB.Model(&models.User{}).Preload("Accounts").Preload("Revenue").Preload("Expenses").Find(&users).Error
        if err != nil {
            return StatusError{500, err}
        }

		data, err := json.Marshal(users)
		if err != nil {
		    return StatusError{500, err}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case http.MethodPost:
		user := &models.User{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(user)
		if err != nil {
            return StatusError{500, err}
		}
		env.DB.Create(user)

		data, err := json.Marshal(user)
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

