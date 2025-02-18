package dto

type CreateProductDTO struct {
	Title         string `json:"title" binding:"required,max=250"`
	Slug          string `json:"slug"`
	Description   string `json:"description" binding:"max=500"`
	Count         int    `json:"count" binding:"required,numeric,min=0"`
	Price         int    `json:"price" binding:"required,numeric,min=0"`
	Cover         string `json:"cover" binding:"required"`
	BrandID       int    `json:"brand_id" binding:"required,numeric,min=1"`
	SubCategoryID int    `json:"sub_category_id" binding:"required,numeric,min=1"`
}
