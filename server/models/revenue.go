package models

type Revenue struct {
	ID                uint    `gorm:"primaryKey"`
	UserID            uint    `gorm:"not null"`
	AccountID         uint    `gorm:"not null"`
	Name              string  `gorm:"not null"`
	Description       string  `gorm:"not null"`
	Amount            float32 `gorm:"not null"`
	DepositDate       int     `gorm:"not null"`
	Recurring         bool    `gorm:"not null"`
	RecurringInterval int     `gorm:"not null"`
	CreatedAt         int
	UpdatedAt         int
}
