package db

import (
    "budgetly/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=admin password=super_secret_password dbname=budgetly port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
    db.AutoMigrate(&models.Account{}, &models.Revenue{}, &models.Expense{})

	return db
}
