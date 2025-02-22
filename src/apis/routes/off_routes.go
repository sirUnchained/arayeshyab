package routes

import (
	"arayeshyab/src/apis/middleware"

	"github.com/gin-gonic/gin"
)

func OffRoutes(r *gin.RouterGroup) {

	r.GET("/off")
	r.POST("/off", middleware.Authorize(), middleware.RoleProtect())
	r.DELETE("/off", middleware.Authorize(), middleware.RoleProtect())
}
