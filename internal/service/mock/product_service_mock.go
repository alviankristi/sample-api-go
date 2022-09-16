package mock

import (
	"context"

	"github.com/alviankristi/catalyst-backend-task/internal/service/model"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (m *ProductServiceMock) Create(ctx context.Context, product model.ProductModel) (*model.ProductResponseModel, error) {
	args := m.Called(ctx, product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ProductResponseModel), args.Error(1)
}
func (m *ProductServiceMock) GetById(ctx context.Context, id int) (*model.ProductResponseModel, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ProductResponseModel), args.Error(1)
}
func (m *ProductServiceMock) GetByBrandId(ctx context.Context, id int) ([]*model.ProductResponseModel, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.ProductResponseModel), args.Error(1)
}
