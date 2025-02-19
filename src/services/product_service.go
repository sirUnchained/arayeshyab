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

func (ph *productService) GetAll(ctx *gin.Context) *helpers.Result {
	query_str := "price > 0 "

	brandID_str := ctx.Query("brand-id")
	SubCategoryID_str := ctx.Query("sub-category-id")

	newest := ctx.Query("newest")
	reachest := ctx.Query("reachest")

	limit_str := ctx.Query("limit")
	page_str := ctx.Query("page")

	if id, err := strconv.Atoi(brandID_str); err == nil {
		query_str += fmt.Sprintf("AND brand_id = %d ", id)
	}
	if id, err := strconv.Atoi(SubCategoryID_str); err == nil {
		query_str += fmt.Sprintf("AND sub_category_id = %d ", id)
	}

	if isNewest, err := strconv.Atoi(newest); err == nil {
		if isNewest == 0 {
			query_str += "ORDER BY created_at asc "
		} else {
			query_str += "ORDER BY created_at desc "
		}
	} else if isReachest, err := strconv.Atoi(reachest); err == nil {
		if isReachest == 0 {
			query_str += "ORDER BY price asc "
		} else {
			query_str += "ORDER BY price desc "
		}
	}

	var page, limit int
	page, err := strconv.Atoi(page_str)
	if err != nil {
		page = 1
	}
	limit, err = strconv.Atoi(limit_str)
	if err != nil {
		limit = 10
	}

	products := new([]schemas.Product)
	db := mysql_db.GetDB()
	err = db.
		Model(&schemas.Product{}).
		Where(query_str).
		Limit(limit).
		Offset((page-1)*limit).
		Select("id", "title", "slug", "pic", "count", "price", "brand_id", "sub_category_id").
		Find(products).Error
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "بفرمایید", Data: products}

}

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
	check_sub_brand := new(schemas.SubCategory)
	db.Model(check_sub_brand).Where("id = ?", brandID).First(check_sub_brand)
	if check_sub_brand.ID == 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ای دی برند بندی یافت نشد", Data: nil}
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

func (ph *productService) Update(ctx *gin.Context) *helpers.Result {
	productID_str := ctx.Param("id")
	id, err := strconv.Atoi(productID_str)
	if err != nil {
		return &helpers.Result{Ok: false, Status: 404, Message: "محصولی با ای دی مورد نظر یافت نشد", Data: nil}
	}

	updating_product := new(schemas.Product)
	db := mysql_db.GetDB()
	db.Model(updating_product).Where("id = ?", id).First(updating_product)
	if updating_product.ID == 0 {
		return &helpers.Result{Ok: false, Status: 404, Message: "محصولی با ای دی مورد نظر یافت نشد", Data: nil}
	}

	title := ctx.PostForm("title")
	if title == "" || len(title) > 250 {
		fmt.Println(title)
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان الزامی بوده و باید کمتر از ۲۵۰ حرف باشد", Data: nil}
	}
	updating_product.Title = title

	slug := strings.Replace(title, " ", "-", -1)
	check_slug_exist := new(schemas.Product)
	db.Model(check_slug_exist).
		Where("id != ?", updating_product.ID).
		Where("slug = ?", slug).
		First(check_slug_exist)
	if check_slug_exist.ID != 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "عنوان از قبل وجود دارد لطفا یکی دیگر انتخاب کنید", Data: nil}
	}
	updating_product.Slug = slug

	description := ctx.PostForm("description")
	if description == "" || len(description) > 500 {
		return &helpers.Result{Ok: false, Status: 400, Message: "توضیحات الزامی بوده و باید کمتر از ۵۰۰ حرف باشد", Data: nil}
	}
	updating_product.Description = description

	count := ctx.PostForm("count")
	count_int, err := strconv.Atoi(count)
	if err != nil || count_int < 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ورودی تعداد معتبر نیست", Data: nil}
	}
	updating_product.Count = uint(count_int)

	price := ctx.PostForm("price")
	price_int, err := strconv.Atoi(price)
	if err != nil || price_int < 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ورودی مبلغ معتبر نیست", Data: nil}
	}
	updating_product.Price = uint(price_int)

	brandID := ctx.PostForm("brand_id")
	brandID_int, err := strconv.Atoi(brandID)
	if err != nil || brandID_int <= 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ای دی برند معتبر نیست", Data: nil}
	}
	check_sub_brand := new(schemas.SubCategory)
	db.Model(check_sub_brand).Where("id = ?", brandID).First(check_sub_brand)
	if check_sub_brand.ID == 0 {
		return &helpers.Result{Ok: false, Status: 400, Message: "ای دی برند بندی یافت نشد", Data: nil}
	}
	updating_product.BrandID = uint(brandID_int)

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
	updating_product.SubCategoryID = uint(SubCategoryID_int)

	cover, err := ctx.FormFile("cover")
	if err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: "یک تصویر برای محصول الزامی است", Data: nil}
	}

	filename := fmt.Sprintf("%s-%d-%s", time.Now().Format("2020122921093"), rand.Intn(10e10), cover.Filename)

	err = ctx.SaveUploadedFile(cover, fmt.Sprintf("./public/%s/%s", check_sub_cat.Slug, filename))
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}

	err = helpers.RemoveFile(updating_product.Pic)
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}
	updating_product.Pic = fmt.Sprintf("/public/%s/%s", check_sub_cat.Slug, filename)

	err = db.Save(updating_product).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "بروز رسانی با موفقیت انجام شد", Data: updating_product}

}

func (ph *productService) Remove(ctx *gin.Context) *helpers.Result {
	productID_str := ctx.Param("id")
	id, err := strconv.Atoi(productID_str)
	if err != nil {
		return &helpers.Result{Ok: false, Status: 404, Message: "محصولی با ای دی مورد نظر یافت نشد", Data: nil}
	}

	removing_product := new(schemas.Product)
	db := mysql_db.GetDB()
	db.Model(removing_product).Where("id = ?", id).First(removing_product)
	if removing_product.ID == 0 {
		return &helpers.Result{Ok: false, Status: 404, Message: "محصولی با ای دی مورد نظر یافت نشد", Data: nil}
	}

	err = db.Delete(removing_product).Error
	if err != nil {
		return &helpers.Result{Ok: false, Status: 500, Message: "مشکلی پیش امده و بزودی رفع خواهد شد", Data: nil}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "عملیات حذف با موفقیت انجام شد", Data: removing_product}
}
