package models

type Expense struct {
	ID                uint    `gorm:"primaryKey"`
	UserID            uint    `gorm:"not null"`
	AccountID         uint    `gorm:"not null"`
	Name              string  `gorm:"not null"`
	Description       string  `gorm:"not null"`
	Amount            float32 `gorm:"not null"`
	Paid              bool    `gorm:"not null"`
	DueDate           int     `gorm:"not null"`
	Recurring         bool    `gorm:"not null"`
	RecurringInterval int     `gorm:"not null"`
	CreatedAt         int
	UpdatedAt         int
}
