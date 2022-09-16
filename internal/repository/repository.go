package repository

import "database/sql"

type Repository struct {
	BrandRepository        BrandRepository
	ProductRepository      ProductRepository
	TransactionRespository TransactionRespository
}

//create repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		BrandRepository:        NewBrandRepository(db),
		ProductRepository:      NewProductRepository(db),
		TransactionRespository: NewTransactionRepository(db),
	}
}
