package handlers

import "github.com/gin-gonic/gin"

type categoryHandler struct{}

func GetCategoryHandler() *categoryHandler {
	return &categoryHandler{}
}

func (ch *categoryHandler) GetAll(ctx *gin.Context) {}

func (ch *categoryHandler) Create(ctx *gin.Context) {}

func (ch *categoryHandler) Update(ctx *gin.Context) {}

func (ch *categoryHandler) Remove(ctx *gin.Context) {}
