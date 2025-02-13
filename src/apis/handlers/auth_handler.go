package handlers

import "github.com/gin-gonic/gin"

type AuthHandlers struct{}

func GetAuthHandlers() *AuthHandlers {
	return &AuthHandlers{}
}

func (ah *AuthHandlers) Login(ctx *gin.Context) {}

func (ah *AuthHandlers) Register(ctx *gin.Context) {}
