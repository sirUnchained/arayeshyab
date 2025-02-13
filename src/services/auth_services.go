package services

import "github.com/gin-gonic/gin"

type AuthServices struct{}

func GetAuthServices() *AuthServices {
	return &AuthServices{}
}

func (ah *AuthServices) Login(ctx *gin.Context) {}

func (ah *AuthServices) Register(ctx *gin.Context) {}
