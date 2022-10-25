package models

type Account struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Type        string  `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Revenue     []Revenue
	Expenses    []Expense
	CreatedAt   int
	UpdatedAt   int
}
