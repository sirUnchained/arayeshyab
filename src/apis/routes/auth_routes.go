package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/auth/login")
	r.POST("/auth/register")
}
