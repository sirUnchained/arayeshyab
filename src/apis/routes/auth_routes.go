package routes

import (
	"arayeshyab/src/apis/handlers"
	"arayeshyab/src/apis/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	ah := handlers.GetAuthHandlers()

	r.POST("/auth/login", ah.Login)
	r.POST("/auth/register", ah.Register)
	r.GET("/auth/get-me", middleware.Authorize(), ah.GetMe)
}
