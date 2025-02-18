package routes

import "github.com/gin-gonic/gin"

func ProductRoutes(r *gin.RouterGroup) {
	r.GET("/product")
	r.POST("/product")
	r.PUT("/product")
	r.DELETE("/product")
}
