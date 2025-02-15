package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

func GetUserService() *userService {
	return &userService{}
}

func (us *userService) GetAll(ctx *gin.Context) *helpers.Result {
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

	var users []schemas.User
	db := mysql_db.GetDB()
	db.Model(&schemas.User{}).
		Where("deleted_at = null").
		Take(limit).
		Offset((page - 1) * limit).
		Find(&users)

	return &helpers.Result{Ok: true, Status: 200, Message: "بفرمایید", Data: users}
}

func (us *userService) Update(ctx *gin.Context) *helpers.Result {
	var userData dto.UpdateUserDTO
	err := ctx.ShouldBindBodyWithJSON(&userData)
	if err != nil {
		errs := dto.UpdateUserDTO_validation(err)
		return &helpers.Result{Ok: false, Status: 400, Message: "لطفا با دقت اطلاعات را وارد کنید", Data: errs}
	}

	user, _ := ctx.Get("user")
	updateUser := new(schemas.User)
	db := mysql_db.GetDB()

	db.Model(&schemas.User{}).Where("id = ?", user.(schemas.User).ID).First(updateUser)
	if updateUser.ID == 0 {
		fmt.Printf("%+v\n", user)
		return &helpers.Result{Ok: false, Status: 404, Message: "کاربری با مشخصات شما یافت نشد", Data: nil}
	}

	updateUser.FullName = userData.FullName
	updateUser.UserName = userData.UserName
	updateUser.Email = userData.Email
	updateUser.Address = userData.Address

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), 15)
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما پیش امده و به زودی رفع خواهد شد", Data: nil}
	}
	updateUser.Password = string(hashedPassword)

	err = db.Save(&updateUser).Error
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما پیش امده و به زودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "بروزرسانی اطلاعات با موفقیت انجام شد", Data: updateUser}
}

func (us *userService) Ban(ctx *gin.Context) *helpers.Result {
	// todo
}
