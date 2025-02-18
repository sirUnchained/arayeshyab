package schemas

import "time"

type Brand struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"type:varchar(100);unique"`
	Slug      string `gorm:"type:varchar(100);unique"`
	Logo      string `gorm:"type:varchar(250)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
