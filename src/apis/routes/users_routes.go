package routes

import (
	"arayeshyab/src/apis/handlers"
	"arayeshyab/src/apis/middleware"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.RouterGroup) {
	uh := handlers.GetUserHandler()

	r.GET("/user", middleware.Authorize(), uh.GetAll)
	r.PUT("/user", middleware.Authorize(), uh.Update)
	r.DELETE("/user/:id", middleware.Authorize(), uh.Ban)
}
