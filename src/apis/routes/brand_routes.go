package routes

import "github.com/gin-gonic/gin"

func BrandRoutes(r *gin.RouterGroup) {
	r.GET("/brand")
	r.POST("/brand")
	r.DELETE("/brand/:id")
}
