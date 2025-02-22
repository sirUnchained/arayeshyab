package handlers

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

type offHandler struct{}

func GetOffHandler() *offHandler {
	return &offHandler{}
}

func (oh *offHandler) GetAll(ctx *gin.Context) {
	off_se := services.GetOffService()

	result := off_se.GetAll()

	helpers.SendResult(result, ctx)
}

func (oh *offHandler) Create(ctx *gin.Context) {
	off_se := services.GetOffService()

	result := off_se.Create(ctx)

	helpers.SendResult(result, ctx)
}

func (oh *offHandler) Remove(ctx *gin.Context) {
	off_se := services.GetOffService()

	result := off_se.Remove(ctx)

	helpers.SendResult(result, ctx)
}
