package routes

import (
	"arayeshyab/src/apis/handlers"
	"arayeshyab/src/apis/middleware"

	"github.com/gin-gonic/gin"
)

func OffRoutes(r *gin.RouterGroup) {
	oh := handlers.GetOffHandler()

	r.GET("/off", middleware.Authorize(), middleware.RoleProtect(), oh.GetAll)
	r.POST("/off", middleware.Authorize(), middleware.RoleProtect(), oh.Create)
	r.DELETE("/off", middleware.Authorize(), middleware.RoleProtect(), oh.Remove)
}
