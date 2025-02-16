package services

import "github.com/gin-gonic/gin"

type categoryService struct{}

func GetCategoryService() *categoryService {
	return &categoryService{}
}

func (ch *categoryService) GetAll(ctx *gin.Context) {}

func (ch *categoryService) Create(ctx *gin.Context) {}

func (ch *categoryService) Update(ctx *gin.Context) {}

func (ch *categoryService) Remove(ctx *gin.Context) {}
