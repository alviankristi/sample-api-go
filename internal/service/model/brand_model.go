package model

type BrandModel struct {
	Name string `json:"name" validate:"required"`
}

type BrandResponseModel struct {
	Name string `json:"name" `
	Id   int    `json:"id" `
}
