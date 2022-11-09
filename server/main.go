package main

import (
	"budgetly/db"
	"budgetly/handlers"

	"log"
	"net/http"
	"os"
)

func main() {
	db := db.Connect()
	os.Setenv("PORT", "5050")

	env := &handler.Env{
		DB:   db,
        Port: ":" + os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to the Budgetly API"))
    })
	http.Handle("/api/user/", handler.Handler{env, handler.User})
	http.Handle("/api/account/", handler.Handler{env, handler.Account})
	http.Handle("/api/revenue/", handler.Handler{env, handler.Revenue})
	http.Handle("/api/expense/", handler.Handler{env, handler.Expense})

	log.Printf("\nServer running on http://localhost%s\n", env.Port)
	log.Fatal(http.ListenAndServe(env.Port, nil))
}
