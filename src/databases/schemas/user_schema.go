package schemas

import "time"

type User struct {
	ID        uint   `gorm:"primarykey"`
	FullName  string `gorm:"type:varchar(250);"`
	UserName  string `gorm:"type:varchar(250);unique; not null"`
	Email     string `gorm:"type:varchar(250);unique; not null"`
	Address   string `gorm:"type:Text"`
	Password  string `gorm:"type:Text; not null"`
	Role      string `gorm:"type:varchar(50);default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
