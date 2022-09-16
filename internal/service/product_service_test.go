package service

import (
	"context"
	"testing"
	"time"

	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	mockRepo "github.com/alviankristi/catalyst-backend-task/internal/repository/mock"
	"github.com/alviankristi/catalyst-backend-task/internal/service/model"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProductServiceTestSuite struct {
	suite.Suite
	productService        ProductService
	productRepositoryMock *mockRepo.ProductRepositoryMock
}

func TestProductServiceTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}

func (t *ProductServiceTestSuite) SetupTest() {
	t.productRepositoryMock = new(mockRepo.ProductRepositoryMock)

	t.productService = NewProductService(t.productRepositoryMock)
}

func (t *ProductServiceTestSuite) TestGetProductByBrandIdSuccess() {
	e := []*entity.ProductEntity{{
		Name:       "name",
		BrandId:    1,
		Price:      10,
		BaseEntity: entity.NewBaseEntity(1, time.Now()),
	}}
	t.productRepositoryMock.On("GetByBrandId", mock.Anything, mock.Anything).Return(e, nil)
	result, err := t.productService.GetByBrandId(context.Background(), 1)

	t.Nil(err)
	t.NotEmpty(result[0].Id)
	t.NotEmpty(result[0].Name)
}

func (t *ProductServiceTestSuite) TestGetProductByBrandIdFailed() {

	t.productRepositoryMock.On("GetByBrandId", mock.Anything, mock.Anything).Return(nil, response.DatabaseError)
	result, err := t.productService.GetByBrandId(context.Background(), 1)

	t.Equal(err, response.DatabaseError)
	t.Nil(result)
}

func (t *ProductServiceTestSuite) TestGetProductSuccess() {
	e := &entity.ProductEntity{
		Name:       "name",
		BrandId:    1,
		Price:      10,
		BaseEntity: entity.NewBaseEntity(1, time.Now()),
	}
	t.productRepositoryMock.On("GetById", mock.Anything, mock.Anything).Return(e, nil)
	result, err := t.productService.GetById(context.Background(), 1)

	t.Nil(err)
	t.NotEmpty(result.Id)
	t.NotEmpty(result.Name)
}

func (t *ProductServiceTestSuite) TestGetProductFailed() {

	t.productRepositoryMock.On("GetById", mock.Anything, mock.Anything).Return(nil, response.DatabaseError)
	result, err := t.productService.GetById(context.Background(), 1)

	t.Equal(err, response.DatabaseError)
	t.Nil(result)
}

func (t *ProductServiceTestSuite) TestCreateProductSuccess() {
	e := &entity.ProductEntity{
		Name:       "name",
		BaseEntity: entity.NewBaseEntity(1, time.Now()),
	}
	t.productRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(e, nil)
	result, err := t.productService.Create(context.Background(), model.ProductModel{
		Name: "name",
	})

	t.Nil(err)
	t.NotEmpty(result.Id)
	t.NotEmpty(result.Name)
}

func (t *ProductServiceTestSuite) TestCreateProductFailed() {

	t.productRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(nil, response.DatabaseError)
	result, err := t.productService.Create(context.Background(), model.ProductModel{
		Name: "name",
	})

	t.Equal(err, response.DatabaseError)
	t.Nil(result)
}
