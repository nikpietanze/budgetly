package handler

import (
	"budgetly/models"

	"encoding/json"
	"net/http"
)

func Revenue(env *Env, w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		revenue := []models.Revenue{}
		env.DB.Find(&revenue)

		data, err := json.Marshal(revenue)
		if err != nil {
			return StatusError{500, err}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case http.MethodPost:
		revenue := &models.Revenue{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(revenue)
		if err != nil {
			return StatusError{500, err}
		}
		env.DB.Create(revenue)

		data, err := json.Marshal(revenue)
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
