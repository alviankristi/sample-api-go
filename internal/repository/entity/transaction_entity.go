package entity

type TransactionEntity struct {
	BaseEntity
	CustomerId int
	ProductId  int
	TotalOrder int
	Price      int
	TotalPrice int
}
