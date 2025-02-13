package main

import (
	"arayeshyab/src/apis"
	"arayeshyab/src/configs"
	"arayeshyab/src/databases/mysql_db"
)

func main() {
	configs.InitConfigs()
	cfg := configs.GetConfigs()

	// config database
	mysql_db.InitMysql(cfg)

	apis.StartServer(cfg)
}
