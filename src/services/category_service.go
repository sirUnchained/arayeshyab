package services

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/databases/mysql_db"
	"arayeshyab/src/databases/schemas"
	"fmt"
	"math/rand"
	"strconv"
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
	// get and validate title
	title := ctx.PostForm("title")
	if title == "" || len(title) > 50 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان دسته بندی الزلمی است و باید حداکثر ۵۰ حرف باشد", Data: nil}
	}
	// create slug from title and check it is unique in db
	slug := strings.Replace(title, " ", "-", -1)
	checkCategory := new(schemas.Category)
	db := mysql_db.GetDB()
	db.Model(&schemas.Category{}).Where("slug = ?", slug).First(checkCategory)
	if checkCategory.ID != 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان دسته بندی از قبل موجود است", Data: nil}
	}
	// get cover make sure it exist
	cover, err := ctx.FormFile("cover")
	if err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "تصویر دسته بندی الزامی است", Data: nil}
	}
	// make sure cover foramt is jpg
	if !strings.Contains(cover.Filename, "jpg") {
		return &helpers.Result{Ok: false, Status: 400, Message: "فرمت فایل فقطط باید jpg باشد", Data: nil}
	}
	// make sure cover size is 2mb
	if cover.Size > (2 << 20) {
		return &helpers.Result{Ok: false, Status: 400, Message: "اندازه تصویر وارد شده بیش از ۲ مگابایت است", Data: nil}
	}
	// create unique name for file
	fileName := fmt.Sprintf("%s-%d-%s", time.Now().Format("202503032460"), rand.Intn(10e10), cover.Filename)
	cover.Filename = fileName
	// save file
	err = ctx.SaveUploadedFile(cover, fmt.Sprintf("./public/%s", fileName))
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}
	// save datas to db
	newCategory := &schemas.Category{Title: title, Pic: fileName, Slug: slug}
	db.Model(&schemas.Category{}).Create(newCategory)
	// done
	return &helpers.Result{Ok: true, Status: 201, Message: "دسته بندی ایجاد شد", Data: newCategory}
}

func (ch *categoryService) Remove(ctx *gin.Context) *helpers.Result {
	id_str := ctx.Param("id")

	var id int
	var err error
	if id, err = strconv.Atoi(id_str); err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "شناسه دسته بندی معتبر نیست", Data: nil}
	}

	db := mysql_db.GetDB()
	err = db.Delete(&schemas.Category{}, id).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return &helpers.Result{Ok: false, Status: 404, Message: "دسته بندی یافت نشد", Data: nil}
		}

		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "دسته بندی حذف شد", Data: nil}
}
