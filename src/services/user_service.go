package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"fmt"
	"strconv"
	"strings"

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
		Select("full_name", "id", "role", "address", "email", "user_name", "created_at").
		Offset((page - 1) * limit).
		Limit(limit).
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

	// getting current validated user and database
	user, _ := ctx.Get("user")
	db := mysql_db.GetDB()

	// check if username or email get duplicated, then send error
	isUserNameDuplicated := new(schemas.User)
	db.Model(&schemas.User{}).
		Where("id != ?", user.(*schemas.User).ID).
		Where(
			"user_name = ? OR email = ?",
			strings.TrimSpace(userData.UserName),
			strings.TrimSpace(userData.Email)).
		First(isUserNameDuplicated)
	if isUserNameDuplicated.ID != 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "لطفا ایمیل یا نام کاربری دیگری وارد کنید", Data: nil}
	}

	//  updating proccess
	user.(*schemas.User).FullName = userData.FullName
	user.(*schemas.User).UserName = userData.UserName
	user.(*schemas.User).Email = userData.Email
	user.(*schemas.User).Address = userData.Address
	// hash password then update
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.(*schemas.User).Password), 5)
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما پیش امده و به زودی رفع خواهد شد", Data: nil}
	}
	user.(*schemas.User).Password = string(hashedPassword)
	// update if we have no error
	err = db.Save(user).Error
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما پیش امده و به زودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "بروزرسانی اطلاعات با موفقیت انجام شد", Data: user}
}

func (us *userService) Ban(ctx *gin.Context) *helpers.Result {
	userID_str := ctx.Param("id")
	userID, err := strconv.Atoi(userID_str)
	if err != nil {
		return &helpers.Result{Ok: false, Status: 404, Message: "شناسه کاربری یافت نشد", Data: nil}
	}

	banUser := new(schemas.User)
	db := mysql_db.GetDB()
	db.Model(&schemas.User{}).Where("id = ?", userID).First(banUser)
	if banUser.ID == 0 {
		return &helpers.Result{Ok: false, Status: 404, Message: "شناسه کاربری یافت نشد", Data: nil}
	}

	// we cannot ban admins, and we dont let admins ban themselvse
	if banUser.Role == "admin" {
		return &helpers.Result{Ok: false, Status: 400, Message: "نمیتوانم ادمین هارا اخراج بکنم", Data: nil}
	}

	err = db.Model(&schemas.User{}).Delete(banUser).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما رخ داد و بزودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "کاربر اخراج شد", Data: nil}

}
