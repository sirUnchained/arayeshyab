package routes

import (
	"arayeshyab/src/apis/handlers"
	"arayeshyab/src/apis/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	ch := handlers.GetCategoryHandler()

	r.GET("/category", ch.GetAll)
	r.POST("/category", middleware.Authorize(), middleware.RoleProtect(), ch.Create)
	r.PUT("/category/:id", middleware.Authorize(), middleware.RoleProtect(), ch.Update)
	r.DELETE("/category/:id", middleware.Authorize(), middleware.RoleProtect(), ch.Remove)
}
