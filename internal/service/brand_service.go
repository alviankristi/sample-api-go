package service

import (
	"context"

	"github.com/alviankristi/catalyst-backend-task/internal/repository"
	model "github.com/alviankristi/catalyst-backend-task/internal/service/model"
)

type BrandService interface {
	// create new brand
	Create(ctx context.Context, brand model.BrandModel) (*model.BrandResponseModel, error)
}

type brandService struct {
	brandRepository repository.BrandRepository
}

//create brand service
func NewBrandService(brandRepository repository.BrandRepository) BrandService {
	return &brandService{
		brandRepository: brandRepository,
	}
}

//create brand
func (service *brandService) Create(ctx context.Context, brand model.BrandModel) (*model.BrandResponseModel, error) {
	//call repo
	entity, err := service.brandRepository.Create(ctx, brand.Name)
	if err != nil {
		return nil, err
	}

	return &model.BrandResponseModel{
		Name: entity.Name,
		Id:   entity.Id,
	}, nil
}
