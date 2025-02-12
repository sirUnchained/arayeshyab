package apis

import (
	"arayeshyab/src/configs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(cfg *configs.Configs) {
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())

	server.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
