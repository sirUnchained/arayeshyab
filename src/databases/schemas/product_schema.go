package schemas

import "time"

type Product struct {
	ID            uint   `gorm:"primarykey"`
	Title         string `gorm:"type:varchar(250);"`
	Slug          string `gorm:"type:varchar(250);unique"`
	Description   string `gorm:"type:varchar(500)"`
	Pic           string `gorm:"type:varchar(250)"`
	Count         uint
	Price         uint
	OffID         uint `gorm:"index"`
	Off           Off
	BrandID       uint `gorm:"index"`
	Brand         Brand
	SubCategoryID uint `gorm:"index"`
	SubCategory   SubCategory
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
