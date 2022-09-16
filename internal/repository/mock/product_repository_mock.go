package mock

import (
	"context"

	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) Create(context context.Context, model entity.CreatedProductEntity) (*entity.ProductEntity, error) {
	args := m.Called(context, model)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.ProductEntity), args.Error(1)
}

func (m *ProductRepositoryMock) GetById(context context.Context, id int) (*entity.ProductEntity, error) {
	args := m.Called(context, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.ProductEntity), args.Error(1)
}

func (m *ProductRepositoryMock) GetByBrandId(context context.Context, id int) ([]*entity.ProductEntity, error) {
	args := m.Called(context, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.ProductEntity), args.Error(1)
}
