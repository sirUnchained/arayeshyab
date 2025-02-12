package main

import (
	"arayeshyab/src/configs"
	"fmt"
)

func main() {
	configs.InitConfigs()
	cfg := configs.GetConfigs()
	fmt.Printf("%+v\n", cfg)
}
