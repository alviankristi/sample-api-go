package service

import (
	"context"

	"github.com/alviankristi/catalyst-backend-task/internal/repository"
	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	model "github.com/alviankristi/catalyst-backend-task/internal/service/model"
)

type ProductService interface {
	// create new product
	Create(ctx context.Context, product model.ProductModel) (*model.ProductResponseModel, error)

	//get product by id
	GetById(ctx context.Context, id int) (*model.ProductResponseModel, error)

	//get product by brand id
	GetByBrandId(ctx context.Context, id int) ([]*model.ProductResponseModel, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

//create product service
func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

//create product
func (service *productService) Create(ctx context.Context, product model.ProductModel) (*model.ProductResponseModel, error) {
	//call repo
	p := entity.CreatedProductEntity{
		Name:    product.Name,
		BrandId: product.BrandId,
		Price:   product.Price,
	}
	entity, err := service.productRepository.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	return &model.ProductResponseModel{
		Name:    entity.Name,
		Id:      entity.Id,
		BrandId: entity.BrandId,
		Price:   entity.Price,
	}, nil
}

//get product by id
func (service *productService) GetById(ctx context.Context, id int) (*model.ProductResponseModel, error) {
	//call repo
	entity, err := service.productRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, nil
	}

	return &model.ProductResponseModel{
		Name:    entity.Name,
		Id:      entity.Id,
		BrandId: entity.BrandId,
		Price:   entity.Price,
	}, nil
}

//get product by brand id
func (service *productService) GetByBrandId(ctx context.Context, id int) ([]*model.ProductResponseModel, error) {
	//call repo
	entities, err := service.productRepository.GetByBrandId(ctx, id)
	if err != nil {
		return nil, err
	}

	results := []*model.ProductResponseModel{}
	for _, e := range entities {
		results = append(results, &model.ProductResponseModel{
			Name:    e.Name,
			Id:      e.Id,
			BrandId: e.BrandId,
			Price:   e.Price,
		})
	}
	return results, nil
}
