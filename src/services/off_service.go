package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type offService struct{}

func GetOffService() *offService {
	return &offService{}
}

func (oh *offService) GetAll() *helpers.Result {
	offs := new([]schemas.Off)
	db := mysql_db.GetDB()

	db.Model(&schemas.Off{}).Find(offs)

	return &helpers.Result{Ok: true, Status: 200, Message: "بفرماییذ", Data: offs}
}

func (oh *offService) Create(ctx *gin.Context) *helpers.Result {
	var new_off_dto dto.CreateOffDto

	if err := ctx.ShouldBindBodyWithJSON(&new_off_dto); err != nil {
		errs := dto.CreateOffDto_validate(err)
		return &helpers.Result{Ok: false, Status: 400, Message: "لطفا ورودی ها را بررسی و مجدد وارد کنید", Data: errs}
	}

	given_time := time.Now().Add(time.Hour * 24 * time.Duration(new_off_dto.Days))

	new_off := new(schemas.Off)
	new_off.Amount = new_off_dto.Amount
	new_off.Code = new_off_dto.Code
	new_off.ExpiresAt = given_time

	db := mysql_db.GetDB()
	err := db.Create(new_off).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 201, Message: "تخفیف ایجاد شد", Data: new_off}
}

func (oh *offService) Remove(ctx *gin.Context) *helpers.Result {
	off_id_str := ctx.Param("id")
	var id int
	var err error

	if id, err = strconv.Atoi(off_id_str); err != nil {
		return &helpers.Result{Ok: false, Status: 404, Message: "تخفیف یافت نشد", Data: nil}
	}

	remove_off := new(schemas.Off)
	db := mysql_db.GetDB()

	db.Model(remove_off).Where("id = ?", id).First(remove_off)
	if remove_off.ID == 0 {
		return &helpers.Result{Ok: false, Status: 404, Message: "تخفیف یافت نشد", Data: nil}
	}

	err = db.Delete(remove_off).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی از سمت ما پیش امده و بزودی حل خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "تخفیف حذف شد", Data: nil}
}
