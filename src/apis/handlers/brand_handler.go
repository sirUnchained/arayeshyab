package handlers

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

type brandHandler struct{}

func GetBrandHandler() *brandHandler {
	return &brandHandler{}
}

func (bh *brandHandler) GetAll(ctx *gin.Context) {
	bs := services.GetBrandService()

	result := bs.GetAll(ctx)

	helpers.SendResult(result, ctx)
}

func (bh *brandHandler) Create(ctx *gin.Context) {
	bs := services.GetBrandService()

	result := bs.Create(ctx)

	helpers.SendResult(result, ctx)
}

func (bh *brandHandler) Remove(ctx *gin.Context) {
	bs := services.GetBrandService()

	result := bs.Remove(ctx)

	helpers.SendResult(result, ctx)
}
