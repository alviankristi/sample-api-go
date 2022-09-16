package mock

import (
	"context"

	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) Create(context context.Context, transaction *entity.CreatedTransactionEntity) (*entity.TransactionEntity, error) {
	args := m.Called(context, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.TransactionEntity), args.Error(1)
}

func (m *TransactionRepositoryMock) GetById(context context.Context, id int) (*entity.TransactionInformationEntity, error) {
	args := m.Called(context, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.TransactionInformationEntity), args.Error(1)
}
