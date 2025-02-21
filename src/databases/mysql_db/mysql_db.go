package mysql_db

import (
	"arayeshyab/src/configs"
	"arayeshyab/src/databases/schemas"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func GetDB() *gorm.DB {
	return DBClient
}

func InitMysql(cfg *configs.Configs) {
	conn_str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)

	// connect to db
	var err error
	DBClient, err = gorm.Open(mysql.Open(conn_str), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// test db connection
	sqldb, _ := DBClient.DB()
	err = sqldb.Ping()
	if err != nil {
		panic(err)
	}

	// migrate
	DBClient.AutoMigrate(
		&schemas.User{},
		&schemas.Category{},
		&schemas.SubCategory{},
		&schemas.Off{},
		&schemas.Product{},
		&schemas.Brand{},
	)
}
