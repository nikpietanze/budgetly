package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	AvatarUrl string `gorm:"not null"`
	Accounts  []Account
	Revenue   []Revenue
	Expenses  []Expense
	CreatedAt int
	UpdatedAt int
}
