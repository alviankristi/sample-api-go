package model

type ProductModel struct {
	Name    string `json:"name" validate:"required"`
	BrandId int    `json:"brand_id" validate:"required"`
	Price   int    `json:"price" validate:"required"`
}

type ProductResponseModel struct {
	Name    string `json:"name" `
	BrandId int    `json:"brand_id" validate:"required"`
	Price   int    `json:"price" validate:"required"`
	Id      int    `json:"id" `
}
