package helpers

import "github.com/gin-gonic/gin"

type Result struct {
	Ok      bool
	Status  int
	Message string
	Data    interface{}
}

func SendResult(result *Result, ctx *gin.Context) {
	ctx.JSON(result.Status, result)
}
