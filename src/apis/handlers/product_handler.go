package handlers

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/services"

	"github.com/gin-gonic/gin"
)

type productHandler struct{}

func GetProductHandler() *productHandler {
	return &productHandler{}
}

func (ph *productHandler) GetAll(ctx *gin.Context) {}

func (ph *productHandler) GetOne(ctx *gin.Context) {
	ps := services.GetProductService()

	result := ps.GetOne(ctx)

	helpers.SendResult(result, ctx)
}

func (ph *productHandler) Create(ctx *gin.Context) {
	ps := services.GetProductService()

	result := ps.Create(ctx)

	helpers.SendResult(result, ctx)
}

func (ph *productHandler) Update(ctx *gin.Context) {}

func (ph *productHandler) Remove(ctx *gin.Context) {}
