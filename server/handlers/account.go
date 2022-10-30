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
			err := env.DB.Model(&models.Account{}).First(&accounts, id).Error
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			err := env.DB.Model(&models.Account{}).Find(&accounts).Error
			if err != nil {
				return StatusError{500, err}
			}
		}
        _, err := ReturnJson(accounts, w)
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

		_, err2 := ReturnJson([]*models.Account{account}, w)
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
			_, err := ReturnJson([]*models.Account{account}, w)
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
			_, err := ReturnJson([]*models.Account{account}, w)
			if err != nil {
				return StatusError{500, err}
			}

		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	}
	return nil
}
