package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"budgetly/models"
	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"gorm.io/gorm/clause"
)

func User(env *Env, w http.ResponseWriter, r *http.Request) error {
    claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
    if !ok {
        return StatusError{500, errors.New("failed to get validated claims")}
    }

    println(claims)

	id := strings.TrimPrefix(r.URL.Path, "/api/user/")

	switch r.Method {
	case http.MethodGet:
		users := []*models.User{}
		if id != "" {
			err := env.DB.Model(&models.User{}).Preload("Accounts").Preload("Revenue").Preload("Expenses").First(&users, id).Error
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			err := env.DB.Model(&models.User{}).Preload("Accounts").Preload("Revenue").Preload("Expenses").Find(&users).Error
			if err != nil {
				return StatusError{500, err}
			}
		}
        _, err := ReturnJson(users, w)
        if err != nil {
            return StatusError{500, err}
        }
	case http.MethodPost:
		user := &models.User{}
		decoder := json.NewDecoder(r.Body)
        err1 := decoder.Decode(&user)
		if err1 != nil {
			return StatusError{500, err1}
		}
		env.DB.Create(user)

		_, err2 := ReturnJson([]*models.User{user}, w)
		if err2 != nil {
			return StatusError{500, err2}
		}
	case http.MethodPut:
		user := &models.User{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			return StatusError{500, err}
		}

		if id != "" {
			env.DB.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&user)
			_, err := ReturnJson([]*models.User{user}, w)
			if err != nil {
				return StatusError{500, err}
			}
		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	case http.MethodDelete:
		if id != "" {
			user := &models.User{}
			env.DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Delete(&user, id)
			_, err := ReturnJson([]*models.User{user}, w)
			if err != nil {
				return StatusError{500, err}
			}

		} else {
			return StatusError{500, errors.New("a valid ID must be supplied")}
		}
	}
	return nil
}
