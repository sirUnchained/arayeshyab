package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type authServices struct{}

func GetAuthServices() *authServices {
	return &authServices{}
}

func (ah *authServices) Login(ctx *gin.Context) *helpers.Result {
	userData := new(dto.AuthDTO)
	if err := ctx.ShouldBindBodyWithJSON(userData); err != nil {
		errs := dto.AuthDTO_GenerateFailedMap(err)
		return &helpers.Result{Ok: false, Status: 400, Message: "اعتبار سنجی شکست خورد لطفا ورودی هارا با دقت وارد کنید", Data: errs}
	}
	// checking do user exist
	user := new(schemas.User)
	db := mysql_db.GetDB()
	if db.Model(&schemas.User{}).Where("email = ?", userData.Email).First(user); user.ID == 0 {
		return &helpers.Result{Ok: false, Status: 404, Message: "اطلاعات کاربری یافت نشد", Data: nil}
	}
	// compare founded user's password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 400, Message: "اطلاعات کاربری درست وارد نشده اند", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "خوش آمدید", Data: user}
}

func (ah *authServices) Register(ctx *gin.Context) *helpers.Result {
	userData := new(dto.AuthDTO)
	if err := ctx.ShouldBindBodyWithJSON(userData); err != nil {
		errs := dto.AuthDTO_GenerateFailedMap(err)
		return &helpers.Result{Ok: false, Status: 400, Message: "اعتبار سنجی شکست خورد لطفا ورودی هارا با دقت وارد کنید", Data: errs}
	}
	// checking do user exist
	user := new(schemas.User)
	db := mysql_db.GetDB()
	if db.Model(&schemas.User{}).Where("email = ?", userData.Email).First(user); user.ID != 0 {
		return &helpers.Result{Ok: false, Status: 404, Message: "لطفا ایمیل دیگری وارد کنید", Data: nil}
	}
	// hashing password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 15)
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما پیش امده و به زودی رفع خواهد شد", Data: nil}
	}
	// create user
	user.Email = userData.Email
	user.UserName = fmt.Sprintf("کاربر-%s-%d", time.Now().Format("20060102150405"), (rand.Intn(10e10)))
	user.Password = string(hashedPass)
	// if this user is first one make him admin
	var count int64
	if db.Model(&schemas.User{}).Count(&count); count == 0 {
		user.Role = "admin"
	}
	// save in db
	if err := db.Model(&schemas.User{}).Create(user).Error; err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطایی از سمت ما پیش امده و به زودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 201, Message: "خوش آمدید", Data: user}
}
