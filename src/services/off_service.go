package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"time"

	"github.com/gin-gonic/gin"
)

type offService struct{}

func GetOffService() *offService {
	return &offService{}
}

func (oh *offService) GetAll() *helpers.Result {
	return &helpers.Result{Ok: true, Status: 201, Message: "بفرماییذ", Data: nil}
}

func (oh *offService) Create(ctx *gin.Context) *helpers.Result {
	var new_off_dto dto.CreateOffDto

	if err := ctx.ShouldBindBodyWithJSON(&new_off_dto); err != nil {
		errs := dto.CreateOffDto_validate(err)
		return &helpers.Result{Ok: false, Status: 400, Message: "لطفا ورودی ها را بررسی و مجدد وارد کنید", Data: errs}
	}

	if new_off_dto.ExpiresAt.Unix() <= time.Now().Unix() {
		return &helpers.Result{Ok: false, Status: 400, Message: "زمان وارد شده معتبر نیست", Data: nil}
	}

	new_off := new(schemas.Off)
	new_off.Amount = new_off_dto.Amount
	new_off.Code = new_off_dto.Code
	new_off.ExpiresAt = new_off_dto.ExpiresAt

	db := mysql_db.GetDB()
	err := db.Create(new_off).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 201, Message: "تخفیف ایجاد شد", Data: new_off}
}

func (oh *offService) Remove(ctx *gin.Context) *helpers.Result {
	return &helpers.Result{Ok: true, Status: 201, Message: "تخفیف حذف شد", Data: nil}
}
