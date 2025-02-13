package handlers

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct{}

func GetAuthHandlers() *AuthHandlers {
	return &AuthHandlers{}
}

func (ah *AuthHandlers) Login(ctx *gin.Context) {
	as := services.GetAuthServices()

	result := as.Login(ctx)
	if !result.Ok {
		helpers.SendResult(result, ctx)
		return
	}

	helpers.SendResult(result, ctx)
}

func (ah *AuthHandlers) Register(ctx *gin.Context) {
	as := services.GetAuthServices()

	result := as.Register(ctx)
	if !result.Ok {
		helpers.SendResult(result, ctx)
		return
	}

	helpers.SendResult(result, ctx)
}
