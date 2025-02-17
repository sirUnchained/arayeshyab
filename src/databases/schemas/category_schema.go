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

type SubCategory struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"type:varchar(50)"`
	Slug      string `gorm:"type:varchar(50);unique"`
	ParrentID uint
	Parrent   Category
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SubSubCategory struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"type:varchar(50)"`
	Slug      string `gorm:"type:varchar(50);unique"`
	ParrentID uint
	Parrent   Category
	CreatedAt time.Time
	UpdatedAt time.Time
}
