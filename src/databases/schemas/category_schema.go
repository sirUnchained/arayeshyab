package schemas

import "time"

type Category struct {
	ID       uint      `gorm:"primarykey"`
	Title    string    `gorm:"type:varchar(50)"`
	Slug     string    `gorm:"type:varchar(50);unique"`
	Pic      string    `gorm:"type:varchar(255)"`
	ParentID *uint     `gorm:"column:parent_id"`
	Parent   *Category `gorm:"foreignKey:ParentID"`
	// Children  []Category `gorm:"foreignKey:ParentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SubCategory struct {
	ID        uint      `gorm:"primarykey"`
	Title     string    `gorm:"type:varchar(50)"`
	Slug      string    `gorm:"type:varchar(50);unique"`
	ParentID  *uint     `gorm:"column:parent_id"`
	Parent    *Category `gorm:"foreignKey:ParentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
