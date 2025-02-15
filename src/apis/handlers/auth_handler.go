package handlers

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/schemas"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

type authHandlers struct{}

func GetAuthHandlers() *authHandlers {
	return &authHandlers{}
}

func (ah *authHandlers) Login(ctx *gin.Context) {
	as := services.GetAuthServices()

	auth_result := as.Login(ctx)
	if !auth_result.Ok {
		helpers.SendResult(auth_result, ctx)
		return
	}

	ts := services.GetTokenService()

	token_result := ts.GenerateNewTokens(auth_result.Data.(*schemas.User), ctx)
	if !token_result.Ok {
		helpers.SendResult(token_result, ctx)
		return
	}

	helpers.SendResult(auth_result, ctx)
}

func (ah *authHandlers) Register(ctx *gin.Context) {
	as := services.GetAuthServices()

	auth_result := as.Register(ctx)
	if !auth_result.Ok {
		helpers.SendResult(auth_result, ctx)
		return
	}

	ts := services.GetTokenService()

	token_result := ts.GenerateNewTokens(auth_result.Data.(*schemas.User), ctx)
	if !token_result.Ok {
		helpers.SendResult(token_result, ctx)
		return
	}

	helpers.SendResult(auth_result, ctx)
}

func (ah *authHandlers) GetMe(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	helpers.SendResult(&helpers.Result{Ok: true, Status: 200, Message: "شما شناخته شده اید", Data: user}, ctx)
}
