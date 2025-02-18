package routes

import (
	"arayeshyab/src/apis/handlers"
	"arayeshyab/src/apis/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	ph := handlers.GetProductHandler()

	r.GET("/product", ph.GetAll)
	r.GET("/product/:id", ph.GetOne)
	r.POST("/product", middleware.Authorize(), middleware.RoleProtect(), ph.Create)
	r.PUT("/product/:id", middleware.Authorize(), middleware.RoleProtect(), ph.Update)
	r.DELETE("/product/:id", middleware.Authorize(), middleware.RoleProtect(), ph.Remove)
}
