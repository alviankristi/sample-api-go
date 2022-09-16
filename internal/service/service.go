package service

import (
	"github.com/alviankristi/catalyst-backend-task/internal/repository"
)

type Service struct {
	BrandService       BrandService
	ProductService     ProductService
	TransactionService TransactionService
}

//create service
func NewService(repository *repository.Repository) *Service {
	return &Service{
		BrandService:       NewBrandService(repository.BrandRepository),
		ProductService:     NewProductService(repository.ProductRepository),
		TransactionService: NewTransactionService(repository.TransactionRespository),
	}
}
