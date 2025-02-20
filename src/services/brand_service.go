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
	new_brand := new(schemas.Brand)

	// get and validate title
	title := ctx.PostForm("title")
	if title == "" || len(title) > 50 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان برند الزلمی است و باید حداکثر ۵۰ حرف باشد", Data: nil}
	}
	new_brand.Title = title

	slug := strings.Replace(title, " ", "-", -1)

	// check slug is duplicated between categories
	check_brand := new(schemas.Brand)
	db := mysql_db.GetDB()
	db.Model(&schemas.Brand{}).Where("slug = ?", slug).First(check_brand)
	if check_brand.ID != 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان برند از قبل موجود است", Data: nil}
	}
	new_brand.Slug = slug

	// ** logo
	logo, err := ctx.FormFile("logo")
	if err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "یک تصویر برای برند الزامی است", Data: nil}
	}
	// make sure logo foramt is jpg
	if !strings.Contains(logo.Filename, "jpg") {
		return &helpers.Result{Ok: false, Status: 400, Message: "فرمت لوگو فقط باید jpg باشد", Data: nil}
	}
	// make sure logo size is 2mb
	if logo.Size > (2 << 20) {
		return &helpers.Result{Ok: false, Status: 400, Message: "اندازه تصویر وارد شده بیش از ۲ مگابایت است", Data: nil}
	}
	logo_name := fmt.Sprintf("%s-%d-%s", time.Now().Format("2020122921093"), rand.Intn(10e10), logo.Filename)
	new_brand.Logo = fmt.Sprintf("/public/brands/%s", logo_name)

	// ** clip
	clip, err := ctx.FormFile("clip")
	if err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "یک کلیپ برای برند الزامی است", Data: nil}
	}
	// make sure clip foramt is jpg
	if !strings.Contains(clip.Filename, "mp4") {
		return &helpers.Result{Ok: false, Status: 400, Message: "فرمت کلیپ فقط باید mp4 باشد", Data: nil}
	}
	// make sure clip size is 10mb
	if clip.Size > (10 << 20) {
		return &helpers.Result{Ok: false, Status: 400, Message: "اندازه تصویر وارد شده بیش از ۲ مگابایت است", Data: nil}
	}
	clip_name := fmt.Sprintf("%s-%d-%s", time.Now().Format("2020122921093"), rand.Intn(10e10), clip.Filename)
	new_brand.Clip = fmt.Sprintf("/public/brands/%s", clip_name)

	// ** save clip and logo
	err = ctx.SaveUploadedFile(clip, fmt.Sprintf("./public/brands/%s", clip_name))
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}
	err = ctx.SaveUploadedFile(logo, fmt.Sprintf("./public/brands/%s", logo_name))
	if err != nil {
		helpers.RemoveFile(new_brand.Clip)
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}

	err = db.Model(new_brand).Save(new_brand).Error
	if err != nil {
		helpers.RemoveFile(new_brand.Clip)
		helpers.RemoveFile(new_brand.Logo)

		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و به زوذی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 201, Message: "برند ایجاد شد", Data: new_brand}

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
