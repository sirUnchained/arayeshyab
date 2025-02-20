package services

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

type brandService struct{}

func GetBrandService() *brandService {
	return &brandService{}
}

func (bh *brandService) GetAll(ctx *gin.Context) *helpers.Result {
	limit_str := ctx.Query("limit")
	page_str := ctx.Query("page")

	limit, err := strconv.Atoi(limit_str)
	if err != nil {
		limit = 5
	}

	page, err := strconv.Atoi(page_str)
	if err != nil {
		page = 1
	}

	var brands []schemas.Brand
	db := mysql_db.GetDB()
	db.Model(&schemas.Brand{}).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&brands)

	return &helpers.Result{Ok: true, Status: 200, Message: "بفرمایید", Data: brands}
}

func (bh *brandService) Create(ctx *gin.Context) *helpers.Result {}

func (bh *brandService) Remove(ctx *gin.Context) *helpers.Result {}
