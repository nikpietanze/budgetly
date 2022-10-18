package models

type Expense struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"not null"`
	Amount    float32 `gorm:"not null"`
	Paid      bool    `gorm:"not null"`
	CreatedAt int
	UpdatedAt int
}
