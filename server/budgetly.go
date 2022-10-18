package main

import (
	"budgetly/db"
	"budgetly/models"

	"encoding/json"
	"fmt"
	"net/http"
)

// Migrate the schema
// db.AutoMigrate(&Product{})

// Create
// db.Create(&Product{Code: "D42", Price: 100})

// Read
// var product Product
// db.First(&product, 1) // find product with integer primary key
//db.First(&product, "code = ?", "D42") // find product with code D42

// Update - update product's price to 200
// db.Model(&product).Update("Price", 200)
// Update - update multiple fields
// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

// Delete - delete product
// db.Delete(&product, 1)

func main() {
	db := db.Connect()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	http.HandleFunc("/api/account", func(w http.ResponseWriter, r *http.Request) {
        // TODO: if /num is supplied, get single, also will need for updated/del
		switch r.Method {
		case "GET":
            accounts := []models.Account{}
			db.Find(&accounts)

            data, err := json.Marshal(accounts)
            if err != nil {
                panic("Error processing data")
            }

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "POST":
			account := &models.Account{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(account)
			if err != nil {
				panic(err)
			}
			db.Create(account)

            data, err := json.Marshal(account)
            if err != nil {
                panic("Error processing data")
            }

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "PUT":
			break
		case "DELETE":
			break
		}
	})

	http.HandleFunc("/api/revenue", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
            revenue := []models.Revenue{}
			db.Find(&revenue)

            data, err := json.Marshal(revenue)
            if err != nil {
                panic("Error processing data")
            }

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "POST":
			revenue := &models.Revenue{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(revenue)
			if err != nil {
				panic(err)
			}
			db.Create(revenue)

            data, err := json.Marshal(revenue)
            if err != nil {
                panic("Error processing data")
            }

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "PUT":
			break
		case "DELETE":
			break
		}
	})

	http.HandleFunc("/api/expense", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
            expenses := []models.Expense{}
			db.Find(&expenses)

            data, err := json.Marshal(expenses)
            if err != nil {
                panic("Error processing data")
            }

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "POST":
			expense := &models.Expense{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(expense)
			if err != nil {
				panic(err)
			}
			db.Create(expense)

            data, err := json.Marshal(expense)
            if err != nil {
                panic("Error processing data")
            }

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "PUT":
			break
		case "DELETE":
			break
		}
	})

	port := ":5050"
	fmt.Printf("\nServer running on http://localhost%s\n", port)

	http.ListenAndServe(port, nil)
}
