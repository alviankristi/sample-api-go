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

type TransactionServiceTestSuite struct {
	suite.Suite
	transactionService        TransactionService
	transactionRepositoryMock *mockRepo.TransactionRepositoryMock
}

func TestTransactionServiceTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionServiceTestSuite))
}

func (t *TransactionServiceTestSuite) SetupTest() {
	t.transactionRepositoryMock = new(mockRepo.TransactionRepositoryMock)

	t.transactionService = NewTransactionService(t.transactionRepositoryMock)
}

func (t *TransactionServiceTestSuite) TestCreateTransactionSuccess() {
	e := &entity.TransactionEntity{
		BaseEntity: entity.NewBaseEntity(1, time.Now()),
		CustomerId: 1,
		Amount:     0,
		Orders:     []entity.TransactionDetailEntity{{Id: 1, ProductId: 1, TotalOrder: 1, Price: 10, TotalPrice: 10}},
	}
	t.transactionRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(e, nil)
	m := model.TransactionModel{
		CustomerId: 1,
		Orders: []model.TransactionDetailModel{
			{
				ProductId:  1,
				TotalOrder: 1,
			},
		},
	}
	result, err := t.transactionService.Create(context.Background(), m)

	t.Nil(err)
	t.NotEmpty(result.Id)
}

func (t *TransactionServiceTestSuite) TestCreateTransactionFailed() {
	e := model.TransactionModel{
		CustomerId: 1,
		Orders:     []model.TransactionDetailModel{{ProductId: 1, TotalOrder: 1}},
	}
	t.transactionRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(nil, response.DatabaseError)
	result, err := t.transactionService.Create(context.Background(), e)

	t.Equal(err, response.DatabaseError)
	t.Nil(result)
}
