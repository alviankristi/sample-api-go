package mock

import (
	"context"

	"github.com/alviankristi/catalyst-backend-task/internal/service/model"
	"github.com/stretchr/testify/mock"
)

type BrandServiceMock struct {
	mock.Mock
}

func (m *BrandServiceMock) Create(ctx context.Context, brand model.BrandModel) (*model.BrandResponseModel, error) {
	args := m.Called(ctx, brand)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.BrandResponseModel), args.Error(1)
}
