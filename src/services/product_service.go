package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type productService struct{}

func GetProductService() *productService {
	return &productService{}
}

// func (ph *productService) GetAll(ctx *gin.Context) *helpers.Result {}

// func (ph *productService) GetOne(ctx *gin.Context) *helpers.Result {}

func (ph *productService) Create(ctx *gin.Context) *helpers.Result {
	// db := mysql_db.GetDB()
	// cover, err := ctx.FormFile("cover")
	// if err != nil {
	// 	return &helpers.Result{Ok: false, Status: 400, Message: "یک تصویر برای محصول الزامی است", Data: nil}
	// }

	var recived_product dto.CreateProductDTO
	var new_product dto.CreateProductDTO

	title := ctx.PostForm("title")
	recived_product.Title = title
	slug := strings.Replace(title, " ", "-", -1)
	recived_product.Slug = slug
	description := ctx.PostForm("description")
	recived_product.Description = description
	count := ctx.PostForm("count")
	count_int, _ := strconv.Atoi(count)
	recived_product.Count = count_int
	price := ctx.PostForm("price")
	price_int, _ := strconv.Atoi(price)
	recived_product.Price = price_int
	brandID := ctx.PostForm("brand_id")
	brandID_int, _ := strconv.Atoi(brandID)
	recived_product.BrandID = brandID_int
	SubCategoryID := ctx.PostForm("sub_category_id")
	SubCategoryID_int, _ := strconv.Atoi(SubCategoryID)
	recived_product.SubCategoryID = SubCategoryID_int

	err := ctx.(&recived_product, &new_product)
	if err != nil {
		return &helpers.Result{Ok: false, Status: 400, Message: err.Error(), Data: nil}
	}

	// check_slug_exist := new(schemas.Product)
	// db.Model(check_slug_exist).Where("slug = ?", slug).First(check_slug_exist)
	// if check_slug_exist.ID != 0 {
	// 	return &helpers.Result{Ok: false, Status: 400, Message: "عنوان از قبل وجود دارد لطفا یکی دیگر انتخاب کنید", Data: nil}
	// }

	return &helpers.Result{Ok: true, Status: 201, Message: "محصول ایجاد گردید", Data: recived_product}

}

// func (ph *productService) Update(ctx *gin.Context) *helpers.Result {}

// func (ph *productService) Remove(ctx *gin.Context) *helpers.Result {}
