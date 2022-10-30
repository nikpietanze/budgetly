package handler

import (
	"budgetly/models"

	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm/clause"
)

func Revenue(env *Env, w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/api/revenue/")

	switch r.Method {
	case http.MethodGet:
		revenue := []*models.Revenue{}

		if id != "" {
			err := env.DB.Model(&models.Revenue{}).First(&revenue, id).Error
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			err := env.DB.Model(&models.Revenue{}).Find(&revenue).Error
			if err != nil {
				return StatusError{500, err}
			}
		}
        _, err := ReturnJson(revenue, w)
        if err != nil {
            return StatusError{500, err}
        }
	case http.MethodPost:
		revenue := &models.Revenue{}
		decoder := json.NewDecoder(r.Body)
        err1 := decoder.Decode(&revenue)
		if err1 != nil {
			return StatusError{500, err1}
		}
		env.DB.Create(revenue)

		_, err2 := ReturnJson([]*models.Revenue{revenue}, w)
		if err2 != nil {
			return StatusError{500, err2}
		}
	case http.MethodPut:
		revenue := &models.Revenue{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&revenue)
		if err != nil {
			return StatusError{500, err}
		}

		if id != "" {
			env.DB.Model(&revenue).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&revenue)
			_, err := ReturnJson([]*models.Revenue{revenue}, w)
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	case http.MethodDelete:
		if id != "" {
			revenue := &models.Revenue{}
			env.DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Delete(&revenue, id)
			_, err := ReturnJson([]*models.Revenue{revenue}, w)
			if err != nil {
				return StatusError{500, err}
			}

		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	}
	return nil
}
