package schemas

import "time"

type Category struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"type:varchar(50)"`
	Slug      string `gorm:"type:varchar(50);unique"`
	Pic       string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
