package services

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"fmt"
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

func (bh *brandService) Create(ctx *gin.Context) *helpers.Result {

}

func (bh *brandService) Remove(ctx *gin.Context) *helpers.Result {
	brandID_str := ctx.Param("id")
	BrandID, err := strconv.Atoi(brandID_str)
	if err != nil {
		return &helpers.Result{Ok: false, Status: 404, Message: "شناسه برند یافت نشد", Data: nil}
	}

	removing_brand := new(schemas.Brand)
	db := mysql_db.GetDB()
	db.Model(&schemas.Brand{}).Where("id = ?", BrandID).First(removing_brand)
	if removing_brand.ID == 0 {
		return &helpers.Result{Ok: false, Status: 404, Message: "شناسه برند یافت نشد", Data: nil}
	}

	err = helpers.RemoveFile(removing_brand.Logo)
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما رخ داد و بزودی رفع خواهد شد", Data: nil}
	}
	err = helpers.RemoveFile(removing_brand.Clip)
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما رخ داد و بزودی رفع خواهد شد", Data: nil}
	}

	err = db.Model(&schemas.Brand{}).Delete(removing_brand).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما رخ داد و بزودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "برند حذف شد", Data: nil}
}
