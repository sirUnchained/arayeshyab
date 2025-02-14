package middleware

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("AccessToken")
		if err != nil {
			helpers.SendResult(&helpers.Result{Ok: false, Status: 401, Message: "لطفا وارد شوید", Data: nil}, ctx)
			ctx.Abort()
			return
		}

		ts := services.GetTokenService()
		claims, result := ts.GetTokenClaims(token)
		if !result.Ok {
			helpers.SendResult(result, ctx)
			ctx.Abort()
			return
		}

		user := new(schemas.User)
		db := mysql_db.GetDB()
		db.Model(&schemas.User{}).Where("id = ?", claims.ID).Select("user_name", "full_name", "email", "created_at", "role", "id").First(user)
		if user.ID == 0 {
			helpers.SendResult(&helpers.Result{Ok: false, Status: 400, Message: "اطلاعات شما پیدا نشد لطفا وارد شوید", Data: nil}, ctx)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
