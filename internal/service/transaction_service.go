package service

import (
	"context"

	"github.com/alviankristi/catalyst-backend-task/internal/repository"
	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	model "github.com/alviankristi/catalyst-backend-task/internal/service/model"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
)

type TransactionService interface {
	Create(ctx context.Context, transaction model.TransactionModel) (*model.TransactionResponseModel, error)
	GetById(ctx context.Context, id int) (*model.TransactionInformationModel, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRespository
}

func NewTransactionService(repository repository.TransactionRespository) TransactionService {
	return &transactionService{
		transactionRepository: repository,
	}
}

func (service *transactionService) GetById(ctx context.Context, id int) (*model.TransactionInformationModel, error) {
	result, err := service.transactionRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	a := &model.TransactionInformationModel{
		Id:           result.Id,
		CustomerId:   result.CustomerId,
		CustomerName: result.CustomerName,
		Amount:       result.Amount,
		Orders:       []*model.TransactionDetailInformationModel{},
	}
	for _, o := range result.Orders {
		x := &model.TransactionDetailInformationModel{
			Id:          o.Id,
			ProductId:   o.ProductId,
			ProductName: o.ProductName,
			TotalOrder:  o.TotalOrder,
			Price:       o.Price,
			TotalPrice:  o.TotalPrice,
			BrandId:     o.BrandId,
			BrandName:   o.BrandName,
		}
		a.Orders = append(a.Orders, x)
	}
	return a, nil
}

//create product
func (service *transactionService) Create(ctx context.Context, transaction model.TransactionModel) (*model.TransactionResponseModel, error) {
	if transaction.Orders == nil || len(transaction.Orders) == 0 {
		return nil, response.TransactionOrderEmpty
	}
	//call repo
	p := &entity.CreatedTransactionEntity{
		CustomerId: transaction.CustomerId,
		Orders:     []entity.CreatedTransactionDetailEntity{},
	}

	for _, detail := range transaction.Orders {
		m := entity.CreatedTransactionDetailEntity{
			ProductId:  detail.ProductId,
			TotalOrder: detail.TotalOrder,
		}
		p.Orders = append(p.Orders, m)
	}

	entity, err := service.transactionRepository.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	m := &model.TransactionResponseModel{
		Id:         entity.Id,
		CustomerId: entity.CustomerId,
		Amount:     entity.Amount,
		Orders:     []model.TransactionDetailResponseModel{},
	}
	for _, v := range entity.Orders {
		x := model.TransactionDetailResponseModel{
			ProductId:  v.ProductId,
			TotalOrder: v.TotalOrder,
			Id:         v.Id,
			TotalPrice: v.TotalPrice,
		}
		m.Orders = append(m.Orders, x)
	}

	return m, nil
}
