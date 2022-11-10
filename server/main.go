package main

import (
	"log"
	"net/http"
	"os"

	"budgetly/db"
    "budgetly/middleware"
	"budgetly/handlers"
)

func main() {
	db := db.Connect()
	os.Setenv("PORT", "5050")
	env := &handler.Env{
		DB:   db,
		Port: ":" + os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}

    middleware := middleware.JWT()

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

