package handlers

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct{}

func GetCategoryHandler() *categoryHandler {
	return &categoryHandler{}
}

func (ch *categoryHandler) GetAll(ctx *gin.Context) {
	cs := services.GetCategoryService()

	result := cs.GetAll()
	if !result.Ok {
		helpers.SendResult(result, ctx)
		return
	}

	helpers.SendResult(result, ctx)
}

func (ch *categoryHandler) CreateCategory(ctx *gin.Context) {
	cs := services.GetCategoryService()

	result := cs.CreateCategory(ctx)
	if !result.Ok {
		helpers.SendResult(result, ctx)
		return
	}

	helpers.SendResult(result, ctx)
}

func (ch *categoryHandler) Remove(ctx *gin.Context) {
	cs := services.GetCategoryService()

	result := cs.Remove(ctx)
	if !result.Ok {
		helpers.SendResult(result, ctx)
		return
	}

	helpers.SendResult(result, ctx)
}
