package handler

import (
	"budgetly/models"

	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm/clause"
)

func Account(env *Env, w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/api/account/")

	switch r.Method {
	case http.MethodGet:
		accounts := []*models.Account{}
		if id != "" {
			err := env.DB.Model(&models.Account{}).Preload("Revenue").Preload("Expenses").First(&accounts, id).Error
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			err := env.DB.Model(&models.Account{}).Preload("Revenue").Preload("Expenses").Find(&accounts).Error
			if err != nil {
				return StatusError{500, err}
			}
		}
        _, err := returnJson(accounts, w)
        if err != nil {
            return StatusError{500, err}
        }
	case http.MethodPost:
		account := &models.Account{}
		decoder := json.NewDecoder(r.Body)
        err1 := decoder.Decode(&account)
		if err1 != nil {
			return StatusError{500, err1}
		}
		env.DB.Create(account)

		_, err2 := returnJson([]*models.Account{account}, w)
		if err2 != nil {
			return StatusError{500, err2}
		}
	case http.MethodPut:
		account := &models.Account{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&account)
		if err != nil {
			return StatusError{500, err}
		}

		if id != "" {
			env.DB.Model(&account).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&account)
			_, err := returnJson([]*models.Account{account}, w)
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	case http.MethodDelete:
		if id != "" {
			account := &models.Account{}
			env.DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Delete(&account, id)
			_, err := returnJson([]*models.Account{account}, w)
			if err != nil {
				return StatusError{500, err}
			}

		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	}
	return nil
}

func returnJson(account []*models.Account, w http.ResponseWriter) (bool, error) {
	data, err := json.Marshal(account)
	if err != nil {
		return false, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return true, nil
}
