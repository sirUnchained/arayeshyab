package handlers

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

type userHandler struct{}

func GetUserHandler() *userHandler {
	return &userHandler{}
}

func (uh *userHandler) GetAll(ctx *gin.Context) {
	us := services.GetUserService()

	result := us.GetAll(ctx)
	if !result.Ok {
		helpers.SendResult(result, ctx)
		return
	}

	helpers.SendResult(result, ctx)
}

func (uh *userHandler) Update(ctx *gin.Context) {
	us := services.GetUserService()

	result := us.Update(ctx)
	if !result.Ok {
		helpers.SendResult(result, ctx)
		return
	}

	helpers.SendResult(result, ctx)
}

func (uh *userHandler) Ban(ctx *gin.Context) {
	us := services.GetUserService()

	result := us.Ban(ctx)
	if !result.Ok {
		helpers.SendResult(result, ctx)
		return
	}

	helpers.SendResult(result, ctx)
}
