package middleware

import (
	"context"
	"log"

	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func JWT() *jwtmiddleware.JWTMiddleware {
	keyFunc := func(ctx context.Context) (interface{}, error) {
		return []byte("secret"), nil
	}

	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
        "http://localhost:3000/",
		[]string{"https://dev-msewkuc22kp85583.us.auth0.com/api/v2/"},
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	return jwtmiddleware.New(jwtValidator.ValidateToken)
}
