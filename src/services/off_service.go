package services

import (
	"arayeshyab/src/apis/helpers"

	"github.com/gin-gonic/gin"
)

type offService struct{}

func GetOffService() *offService {
	return &offService{}
}

func (oh *offService) GetAll() *helpers.Result {}

func (oh *offService) Create(ctx *gin.Context) *helpers.Result {

}

func (oh *offService) Remove(ctx *gin.Context) *helpers.Result {}
