package configs

import (
	"encoding/json"
	"io"
	"os"
)

type Configs struct {
	Mysql  Mysql
	Redis  Redis
	Jwt    Jwt
	Server Server
}

type Mysql struct {
	Host     string
	Port     string
	Password string
	Username string
	Database string
}

type Redis struct {
	Port string
	Host string
}

type Jwt struct {
	AccessTokenKey             string
	AccessTokenExpirePerMinute string
	RefreshTokenKey            string
	RefreshTokenExpirePerDay   string
}

type Server struct {
	Port string
	Host string
}

var cfg Configs

func GetConfigs() *Configs {
	return &cfg
}

func InitConfigs() {
	env := os.Getenv("APP_ENV")
	address := getFileAddress(&env)
	readFileInitCfg(&address)
}

func getFileAddress(env *string) string {
	if *env == "production" {
		return "configs/config.prod.json"
	}
	return "configs/config.dev.json"
}

func readFileInitCfg(path *string) {
	jsonFile, err := os.Open(*path)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &cfg)
}
