package models

type Account struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"not null"`
	Type      string  `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	Revenue   []Revenue
	CreatedAt int
	UpdatedAt int
}
