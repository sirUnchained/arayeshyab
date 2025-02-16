package services

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type categoryService struct{}

func GetCategoryService() *categoryService {
	return &categoryService{}
}

func (ch *categoryService) GetAll() *helpers.Result {
	categories := new([]schemas.Category)

	db := mysql_db.GetDB()
	db.Model(&schemas.Category{}).Find(categories)

	return &helpers.Result{Ok: false, Status: 200, Message: "بفرمایید", Data: categories}
}

func (ch *categoryService) Create(ctx *gin.Context) *helpers.Result {
	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "اندازه اطلاعات وارد شده معتبر نمی باشد", Data: nil}
	}

	title := ctx.PostForm("title")
	if title == "" || len(title) > 50 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان دسته بندی الزلمی است و باید حداکثر ۵۰ حرف باشد", Data: nil}
	}

	cover, err := ctx.FormFile("cover")
	if err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "تصویر دسته بندی الزامی است", Data: nil}
	}

	if !strings.Contains(cover.Filename, "jpg") {
		return &helpers.Result{Ok: false, Status: 400, Message: "فرمت فایل فقطط باید jpg باشد", Data: nil}
	}

	if cover.Size > (2 << 20) {
		return &helpers.Result{Ok: false, Status: 400, Message: "اندازه تصویر وارد شده بیش از ۲ مگابایت است", Data: nil}
	}

	fileName := fmt.Sprintf("%s-%d-%s", time.Now().Format("202503032460"), rand.Intn(10e10), cover.Filename)
	cover.Filename = fileName

	err = ctx.SaveUploadedFile(cover, fmt.Sprintf("./public/%s", fileName))
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}

	slug := strings.Replace(title, " ", "-", -1)

	newCategory := &schemas.Category{Title: title, Pic: fileName, Slug: slug}
	db := mysql_db.GetDB()
	db.Model(&schemas.Category{}).Create(newCategory)

	return &helpers.Result{Ok: true, Status: 201, Message: "دسته بندی ایجاد شد", Data: newCategory}
}

func (ch *categoryService) Remove(ctx *gin.Context)/* *helpers.Result */ {}
