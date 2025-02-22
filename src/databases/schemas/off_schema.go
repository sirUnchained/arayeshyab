package schemas

import "time"

type Off struct {
	ID        uint   `gorm:"primarykey"`
	Amount    uint   `gorm:"max:100;default:0"`
	Code      string `gorm:"type:varchar(250)"`
	CreatedAt time.Time
	ExpiresAt time.Time
	UpdatedAt time.Time
}
