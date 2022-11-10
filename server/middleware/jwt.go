package middleware

import (
	"context"
	"log"

    "budgetly/utils"
	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func JWT() *jwtmiddleware.JWTMiddleware {
	keyFunc := func(ctx context.Context) (interface{}, error) {
		return []byte(utils.GetEnv("SECRET")), nil
	}

	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
        utils.GetEnv("ISSUER"),
		[]string{utils.GetEnv("AUDIENCE")},
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	return jwtmiddleware.New(jwtValidator.ValidateToken)
}
