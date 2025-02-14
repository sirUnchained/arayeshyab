package handlers

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/schemas"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct{}

func GetAuthHandlers() *AuthHandlers {
	return &AuthHandlers{}
}

func (ah *AuthHandlers) Login(ctx *gin.Context) {
	as := services.GetAuthServices()

	auth_result := as.Login(ctx)
	if !auth_result.Ok {
		helpers.SendResult(auth_result, ctx)
		return
	}

	ts := services.GetTokenService()

	token_result := ts.GenerateNewTokens(auth_result.Data.(*schemas.User))
	if !token_result.Ok {
		helpers.SendResult(token_result, ctx)
		return
	}

	helpers.SendResult(token_result, ctx)
}

func (ah *AuthHandlers) Register(ctx *gin.Context) {
	as := services.GetAuthServices()

	auth_result := as.Register(ctx)
	if !auth_result.Ok {
		helpers.SendResult(auth_result, ctx)
		return
	}

	ts := services.GetTokenService()

	token_result := ts.GenerateNewTokens(auth_result.Data.(*schemas.User))
	if !token_result.Ok {
		helpers.SendResult(token_result, ctx)
		return
	}

	helpers.SendResult(token_result, ctx)
}
