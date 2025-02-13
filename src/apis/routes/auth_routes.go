package routes

import (
	"arayeshyab/src/apis/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	ah := handlers.GetAuthHandlers()

	r.POST("/auth/login", ah.Login)
	r.POST("/auth/register", ah.Register)
}
