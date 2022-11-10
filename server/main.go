package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"budgetly/db"
	"budgetly/handlers"

	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func main() {
	db := db.Connect()
	os.Setenv("PORT", "5050")
	env := &handler.Env{
		DB:   db,
		Port: ":" + os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}

    // TODO: move middleware to separate module
	keyFunc := func(ctx context.Context) (interface{}, error) {
		return []byte("secret"), nil
	}

	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
        "https://<issuer-url>/",
		[]string{"<audience>"},
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Budgetly API"))
	})
	http.Handle("/api/user/", middleware.CheckJWT(handler.Handler{env, handler.User}))
	http.Handle("/api/account/", middleware.CheckJWT(handler.Handler{env, handler.Account}))
	http.Handle("/api/revenue/", middleware.CheckJWT(handler.Handler{env, handler.Revenue}))
	http.Handle("/api/expense/", middleware.CheckJWT(handler.Handler{env, handler.Expense}))

	log.Printf("\nServer running on http://localhost%s\n", env.Port)
	log.Fatal(http.ListenAndServe(env.Port, nil))
}
