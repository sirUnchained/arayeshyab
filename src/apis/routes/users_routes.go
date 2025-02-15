package routes

import "github.com/gin-gonic/gin"

func UsersRoutes(r *gin.RouterGroup) {
	r.GET("/users")
	r.PUT("/user")
	r.DELETE("/user/:id")
}
