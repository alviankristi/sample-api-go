package entity

type TransactionEntity struct {
	BaseEntity
	CustomerId int `db:"customer_id"`
	Amount     int `db:"amount"`
	Orders     []TransactionDetailEntity
}
type TransactionDetailEntity struct {
	Id         int `db:"id"`
	ProductId  int `db:"product_id"`
	TotalOrder int `db:"total_order"`
	Price      int `db:"price"`
	TotalPrice int `db:"total_price"`
}
type CreatedTransactionEntity struct {
	CustomerId int `db:"customer_id"`
	Orders     []CreatedTransactionDetailEntity
}

type CreatedTransactionDetailEntity struct {
	ProductId  int `db:"product_id"`
	TotalOrder int `db:"total_order"`
}

type CalculatedTransaction struct {
	ProductId  int `db:"product_id"`
	Price      int `db:"price"`
	TotalOrder int `db:"total_order"`
	TotalPrice int `db:"total_price"`
}

type TransactionInformationEntity struct {
	BaseEntity
	CustomerId   int    `db:"customer_id"`
	CustomerName string `db:"customer_name"`
	Amount       int    `db:"amount"`
	Orders       []*TransactionDetailInformationEntity
}

type TransactionDetailInformationEntity struct {
	Id          int    `db:"id"`
	ProductId   int    `db:"product_id"`
	ProductName string `db:"product_name"`
	TotalOrder  int    `db:"total_order"`
	Price       int    `db:"price"`
	TotalPrice  int    `db:"total_price"`
	BrandId     int    `db:"brand_id"`
	BrandName   string `db:"brand_name"`
}
