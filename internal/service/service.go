package service

import (
	"github.com/alviankristi/catalyst-backend-task/internal/repository"
)

type Service struct {
	BrandService BrandService
}

//create service
func NewService(repository *repository.Repository) *Service {
	return &Service{
		BrandService: NewBrandService(repository.BrandRepository),
	}
}
