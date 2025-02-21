package schemas

import "time"

type Off struct {
	ID        uint `gorm:"primarykey"`
	ProductID uint `gorm:"unique"`
	Amount    uint `gorm:"max:100;default:0"`
	CreatedAt time.Time
	ExpiresAt time.Time
	UpdatedAt time.Time
}
