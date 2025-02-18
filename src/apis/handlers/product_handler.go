package handlers

import "github.com/gin-gonic/gin"

type productHandler struct{}

func GetProductHandler() *productHandler {
	return &productHandler{}
}

func (ph *productHandler) GetAll(ctx *gin.Context) {}

func (ph *productHandler) GetOne(ctx *gin.Context) {}

func (ph *productHandler) Create(ctx *gin.Context) {}

func (ph *productHandler) Update(ctx *gin.Context) {}

func (ph *productHandler) Remove(ctx *gin.Context) {}
