package models

type Revenue struct {
	ID          uint    `gorm:"primaryKey"`
	AccountID   uint    `gorm:"not null"`
	Name        string  `gorm:"not null"`
	Amount      float32 `gorm:"not null"`
	DepositDate int     `gorm:"not null"`
	CreatedAt   int
	UpdatedAt   int
}
