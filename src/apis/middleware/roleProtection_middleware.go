package middleware

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/schemas"

	"github.com/gin-gonic/gin"
)

func RoleProtect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")

		if user.(*schemas.User).Role != "admin" {
			helpers.SendResult(&helpers.Result{Ok: false, Status: 403, Message: "شما اجازه دسترسی به این بخش را ندارید", Data: nil}, ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
