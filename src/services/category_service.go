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
	var categories []schemas.Category

	// get all categories
	db := mysql_db.GetDB()
	db.Model(&schemas.Category{}).Preload("Children").Find(&categories)

	// put sub categories inside parent categories, then remove them
	for i := 0; i < len(categories); i++ {
		for j := 0; j < len(categories); j++ {
			if categories[j].ParentID == categories[i].ID {
				categories[i].Children = append(categories[i].Children, categories[j])
				categories = append(categories[:j], categories[j+1:]...)
			}
		}
	}

	return &helpers.Result{Ok: false, Status: 200, Message: "بفرمایید", Data: categories}
}

func (ch *categoryService) Create(ctx *gin.Context) *helpers.Result {
	new_category := new(schemas.Category)

	// get and validate title
	title := ctx.PostForm("title")
	if title == "" || len(title) > 50 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان دسته بندی الزلمی است و باید حداکثر ۵۰ حرف باشد", Data: nil}
	}
	// create slug from title
	slug := strings.Replace(title, " ", "-", -1)
	// check slug is duplicated between categories
	check_category := new(schemas.Category)
	check_sub_category := new(schemas.SubCategory)
	db := mysql_db.GetDB()

	db.Model(&schemas.Category{}).Where("slug = ?", slug).First(check_category)
	db.Model(&schemas.SubCategory{}).Where("slug = ?", slug).First(check_sub_category)

	if check_sub_category.ID != 0 || check_category.ID != 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان دسته بندی از قبل موجود است", Data: nil}
	}
	new_category.Title = title
	new_category.Slug = slug

	// ok this one is a bit hard, i check if client send an id i go check it in db, if it exist then i add it as parent if dont i'll send error
	parent_ID_str := ctx.PostForm("parent")
	parent_ID, err := strconv.Atoi(parent_ID_str)
	if err == nil {
		parent_category := new(schemas.Category)
		db.Model(&schemas.Category{}).Where("id = ?", parent_ID).First(parent_category)
		if parent_category.ID == 0 {
			return &helpers.Result{Ok: false, Status: 404, Message: "دسته بندی پدر یافت نشد", Data: nil}
		}

		new_sub_category := new(schemas.Category)
		new_sub_category.ParentID = uint(parent_ID)
		new_sub_category.Slug = slug
		new_sub_category.Title = title
		err = db.Model(&schemas.Category{}).Save(new_sub_category).Error
		if err != nil {
			fmt.Println(err)
			return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
		}

		return &helpers.Result{Ok: true, Status: 201, Message: "دسته بندی فرزند ایجاد شد", Data: new_category}
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
	new_category.Pic = fileName

	// save file
	err = ctx.SaveUploadedFile(cover, fmt.Sprintf("./public/categories/%s", fileName))
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}

	// save datas to db
	err = db.Model(&schemas.Category{}).Create(new_category).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 201, Message: "دسته بندی پدر ایجاد شد", Data: new_category}
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
