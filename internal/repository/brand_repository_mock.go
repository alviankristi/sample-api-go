package repository

import (
	"context"

	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	"github.com/stretchr/testify/mock"
)

type BrandRepositoryMock struct {
	mock.Mock
}

func (m *BrandRepositoryMock) Create(ctx context.Context, name string) (*entity.BrandEntity, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.BrandEntity), args.Error(1)
}
