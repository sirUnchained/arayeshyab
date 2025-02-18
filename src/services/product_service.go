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

type productService struct{}

func GetProductService() *productService {
	return &productService{}
}

func (ph *productService) GetAll(ctx *gin.Context) *helpers.Result {}

func (ph *productService) GetOne(ctx *gin.Context) *helpers.Result {
	id_str := ctx.Param("id")
	var id int
	var err error
	if id, err = strconv.Atoi(id_str); err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "ای دی محصول معتبر نیست"}
	}

	result := new(schemas.Product)
	db := mysql_db.GetDB()
	err = db.Model(result).Where("id = ?", id).Preload("Brand").Preload("SubCategory").First(result).Error
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &helpers.Result{Ok: false, Status: 404, Message: "محصول پیدا نشد", Data: nil}
		}
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 201, Message: "محصول ایجاد گردید", Data: result}
}

func (ph *productService) Create(ctx *gin.Context) *helpers.Result {
	title := ctx.PostForm("title")
	if title == "" || len(title) > 250 {
		fmt.Println(title)
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان الزامی بوده و باید کمتر از ۲۵۰ حرف باشد", Data: nil}
	}

	slug := strings.Replace(title, " ", "-", -1)
	db := mysql_db.GetDB()
	check_slug_exist := new(schemas.Product)
	db.Model(check_slug_exist).Where("slug = ?", slug).First(check_slug_exist)
	if check_slug_exist.ID != 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان از قبل وجود دارد لطفا یکی دیگر انتخاب کنید", Data: nil}
	}

	description := ctx.PostForm("description")
	if description == "" || len(description) > 500 {
		return &helpers.Result{Ok: false, Status: 400, Message: "توضیحات الزامی بوده و باید کمتر از ۵۰۰ حرف باشد", Data: nil}
	}

	count := ctx.PostForm("count")
	count_int, err := strconv.Atoi(count)
	if err != nil || count_int < 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ورودی تعداد معتبر نیست", Data: nil}
	}

	price := ctx.PostForm("price")
	price_int, err := strconv.Atoi(price)
	if err != nil || price_int < 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ورودی مبلغ معتبر نیست", Data: nil}
	}

	brandID := ctx.PostForm("brand_id")
	brandID_int, err := strconv.Atoi(brandID)
	if err != nil || brandID_int <= 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ای دی برند معتبر نیست", Data: nil}
	}

	SubCategoryID := ctx.PostForm("sub_category_id")
	SubCategoryID_int, err := strconv.Atoi(SubCategoryID)
	if err != nil || SubCategoryID_int <= 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ای دی زیر دسته بندی معتبر نیست", Data: nil}
	}

	check_sub_cat := new(schemas.SubCategory)
	db.Model(check_sub_cat).Where("id = ?", SubCategoryID).First(check_sub_cat)
	if check_sub_cat.ID == 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ای دی زیر دسته بندی یافت نشد", Data: nil}
	}

	cover, err := ctx.FormFile("cover")
	if err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "یک تصویر برای محصول الزامی است", Data: nil}
	}

	filename := fmt.Sprintf("%s-%d-%s", time.Now().Format("2020122921093"), rand.Intn(10e10), cover.Filename)

	err = ctx.SaveUploadedFile(cover, fmt.Sprintf("./public/%s/%s", check_sub_cat.Slug, filename))
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}

	new_product := new(schemas.Product)
	new_product.Slug = slug
	new_product.Title = title
	new_product.Description = description
	new_product.Pic = fmt.Sprintf("/public/%s/%s", check_sub_cat.Slug, filename)
	new_product.Count = uint(count_int)
	new_product.Price = uint(price_int)
	new_product.SubCategoryID = uint(SubCategoryID_int)
	new_product.BrandID = uint(brandID_int)

	err = db.Create(new_product).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 201, Message: "بفرمایید", Data: new_product}

}

func (ph *productService) Update(ctx *gin.Context) *helpers.Result {}

func (ph *productService) Remove(ctx *gin.Context) *helpers.Result {}
