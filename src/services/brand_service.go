package services

import (
	"arayeshyab/src/apis/helpers"

	"github.com/gin-gonic/gin"
)

type brandService struct{}

func GetBrandService() *brandService {
	return &brandService{}
}

func (bh *brandService) GetAll(ctx *gin.Context) *helpers.Result {}

func (bh *brandService) Create(ctx *gin.Context) *helpers.Result {}

func (bh *brandService) Remove(ctx *gin.Context) *helpers.Result {}
