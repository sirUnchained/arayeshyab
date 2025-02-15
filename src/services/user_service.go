package services

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
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
	// todo
}

func (us *userService) Ban(ctx *gin.Context) *helpers.Result {
	// todo
}
