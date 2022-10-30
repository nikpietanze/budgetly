package handler

import (
	"budgetly/models"

	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm/clause"
)

func Expense(env *Env, w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/api/expense/")

	switch r.Method {
	case http.MethodGet:
		expenses := []*models.Expense{}

		if id != "" {
			err := env.DB.Model(&models.Expense{}).First(&expenses, id).Error
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			err := env.DB.Model(&models.Expense{}).Find(&expenses).Error
			if err != nil {
				return StatusError{500, err}
			}
		}
        _, err := ReturnJson(expenses, w)
        if err != nil {
            return StatusError{500, err}
        }
	case http.MethodPost:
		expense := &models.Expense{}
		decoder := json.NewDecoder(r.Body)
        err1 := decoder.Decode(&expense)
		if err1 != nil {
			return StatusError{500, err1}
		}
		env.DB.Create(expense)

		_, err2 := ReturnJson([]*models.Expense{expense}, w)
		if err2 != nil {
			return StatusError{500, err2}
		}
	case http.MethodPut:
		expense := &models.Expense{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&expense)
		if err != nil {
			return StatusError{500, err}
		}

		if id != "" {
			env.DB.Model(&expense).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&expense)
			_, err := ReturnJson([]*models.Expense{expense}, w)
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	case http.MethodDelete:
		if id != "" {
			expense := &models.Expense{}
			env.DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Delete(&expense, id)
			_, err := ReturnJson([]*models.Expense{expense}, w)
			if err != nil {
				return StatusError{500, err}
			}

		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	}
	return nil
}
