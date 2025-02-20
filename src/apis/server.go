package apis

import (
	"arayeshyab/src/apis/middleware"
	"arayeshyab/src/apis/routes"
	"arayeshyab/src/configs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(cfg *configs.Configs) {
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery(), middleware.AddHeadersSecurity(cfg))

	server.Static("/public", "./public")

	initRoutes(&server.RouterGroup)

	server.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func initRoutes(r *gin.RouterGroup) {
	route_group := r.Group("/v1")

	routes.AuthRoutes(route_group)
	routes.UsersRoutes(route_group)
	routes.CategoryRoutes(route_group)
	routes.ProductRoutes(route_group)
	routes.BrandRoutes(route_group)
}
