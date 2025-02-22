package handlers

import "github.com/gin-gonic/gin"

type offHandler struct{}

func GetOffHandler() *offHandler {
	return &offHandler{}
}

func (oh *offHandler) GetAll(ctx *gin.Context) {}

func (oh *offHandler) Create(ctx *gin.Context) {}

func (oh *offHandler) Remove(ctx *gin.Context) {}
