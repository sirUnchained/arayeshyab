package main

import (
	"arayeshyab/src/apis"
	"arayeshyab/src/configs"
)

func main() {
	configs.InitConfigs()
	cfg := configs.GetConfigs()

	apis.StartServer(cfg)
}
