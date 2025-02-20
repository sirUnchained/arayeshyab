package routes

import (
	"arayeshyab/src/apis/handlers"
	"arayeshyab/src/apis/middleware"

	"github.com/gin-gonic/gin"
)

func BrandRoutes(r *gin.RouterGroup) {
	bh := handlers.GetBrandHandler()

	r.GET("/brand", bh.GetAll)
	r.POST("/brand", middleware.Authorize(), middleware.RoleProtect(), bh.Create)
	r.DELETE("/brand/:id", middleware.Authorize(), middleware.RoleProtect(), bh.Remove)
}
