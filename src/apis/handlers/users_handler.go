package handlers

import "github.com/gin-gonic/gin"

type userHandler struct{}

func GetUserHandler() *userHandler {
	return &userHandler{}
}

func (us *userHandler) GetAll(ctx *gin.Context) {
	// todo
}

func (us *userHandler) Update(ctx *gin.Context) {
	// todo
}

func (us *userHandler) Ban(ctx *gin.Context) {
	// todo
}
