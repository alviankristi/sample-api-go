package model

type TransactionModel struct {
	CustomerId int                      `json:"customer_id" validate:"required"`
	Orders     []TransactionDetailModel `json:"orders" validate:"required"`
}

type TransactionDetailModel struct {
	ProductId  int `json:"product_id" validate:"required"`
	TotalOrder int `json:"total_order" validate:"required"`
}

type TransactionResponseModel struct {
	Id         int                              `json:"id" `
	CustomerId int                              `json:"customer_id"`
	Amount     int                              `json:"amount"`
	Orders     []TransactionDetailResponseModel `json:"orders"`
}

type TransactionDetailResponseModel struct {
	Id         int `json:"id" `
	ProductId  int `json:"product_id"`
	TotalOrder int `json:"total_order"`
	TotalPrice int `json:"total_price"`
}

type TransactionInformationModel struct {
	Id           int                                  `json:"id" `
	CustomerId   int                                  `json:"customer_id"`
	CustomerName string                               `json:"customer_name"`
	Amount       int                                  `json:"amount"`
	Orders       []*TransactionDetailInformationModel `json:"orders"`
}

type TransactionDetailInformationModel struct {
	Id          int    `json:"id" `
	ProductId   int    `json:"product_id" `
	ProductName string `json:"product_name" `
	TotalOrder  int    `json:"total_order" `
	Price       int    `json:"price" `
	TotalPrice  int    `json:"total_price" `
	BrandId     int    `json:"brand_id" `
	BrandName   string `json:"brand_name" `
}
