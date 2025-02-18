package services

type productService struct{}

func GetProductService() *productService {
	return &productService{}
}

// func (ph *productService) GetAll(ctx *gin.Context) *helpers.Result {}

// func (ph *productService) GetOne(ctx *gin.Context) *helpers.Result {}

// func (ph *productService) Create(ctx *gin.Context) *helpers.Result {
// 	cover, err := ctx.FormFile("cover")
// 	if err != nil {
// 		return &helpers.Result{Ok: false, Status: 400, Message: "یک تصویر برای محصول الزامی است", Data: nil}
// 	}

// }

// func (ph *productService) Update(ctx *gin.Context) *helpers.Result {}

// func (ph *productService) Remove(ctx *gin.Context) *helpers.Result {}
