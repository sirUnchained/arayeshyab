package routes

import (
	"arayeshyab/src/apis/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	r.GET("/category")
	r.POST("/category", middleware.Authorize(), middleware.RoleProtect())
	r.PUT("/category/:id", middleware.Authorize(), middleware.RoleProtect())
	r.DELETE("/category/:id", middleware.Authorize(), middleware.RoleProtect())
}
