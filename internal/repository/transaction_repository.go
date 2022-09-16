package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
)

var (
	getTransactionById string = `SELECT a.id, a.customer_id,e.fullname,a.amount,a.created_date,b.id,b.product_id,b.total_order,b.total_price,c.price, c.name,d.name,d.id
	FROM transactions a JOIN transaction_details b ON a.id =b.transaction_id
	JOIN products c ON c.id = b.product_id
	JOIN brands d ON d.id = c.brand_id 
	JOIN customers e ON e.id = a.customer_id
	WHERE a.id = ?`
	createTransaction       string = `INSERT INTO transactions (customer_id, created_date, amount) VALUES (?, ?, ?)`
	createTransactionDetail string = `INSERT INTO transaction_details (transaction_id, product_id, total_order, total_price) VALUES (?, ?, ?, ?)`
)

type TransactionRespository interface {
	Create(context context.Context, transaction *entity.CreatedTransactionEntity) (*entity.TransactionEntity, error)
	GetById(context context.Context, id int) (*entity.TransactionInformationEntity, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRespository {
	return &transactionRepository{
		db: db,
	}
}

func (repo *transactionRepository) GetById(context context.Context, id int) (*entity.TransactionInformationEntity, error) {
	rows, err := repo.db.QueryContext(context, getTransactionById, id)
	if err != nil {
		log.Printf("transactionRepository.GetById() error : %v", err)
		return nil, response.DatabaseError
	}

	defer rows.Close()

	details := []*entity.TransactionDetailInformationEntity{}
	result := &entity.TransactionInformationEntity{}
	if rows.Next() {
		detail := &entity.TransactionDetailInformationEntity{}

		err := rows.Scan(&result.Id, &result.CustomerId, &result.CustomerName,
			&result.Amount, &result.CreatedDate, &detail.Id,
			&detail.ProductId, &detail.TotalOrder, &detail.TotalPrice, &detail.Price, &detail.ProductName,
			&detail.BrandName, &detail.BrandId)
		if err != nil {
			log.Printf("transactionRepository.GetById() - rows.Scan() error : %v", err)
		} else {
			details = append(details, detail)
		}
		for rows.Next() {
			detail := &entity.TransactionDetailInformationEntity{}

			err := rows.Scan(&result.Id, &result.CustomerId, &result.CustomerName,
				&result.Amount, &result.CreatedDate, &detail.Id,
				&detail.ProductId, &detail.TotalOrder, &detail.TotalPrice, &detail.Price, &detail.ProductName,
				&detail.BrandName, &detail.BrandId)
			if err != nil {
				log.Printf("transactionRepository.GetById() - rows.Scan() error : %v", err)
			} else {
				details = append(details, detail)
			}
		}
	} else {
		return nil, nil
	}

	err = rows.Err()
	if err != nil {
		log.Printf("transactionRepository.GetById() - rows.Err() error : %v", err)
		return nil, response.DatabaseError
	}
	result.Orders = details
	return result, nil
}

func (repo *transactionRepository) Create(context context.Context, transaction *entity.CreatedTransactionEntity) (*entity.TransactionEntity, error) {
	// Get a Tx for making transaction requests.
	tx, err := repo.db.BeginTx(context, nil)
	if err != nil {
		log.Printf("transactionRepository.Create() - repo.db.BeginTx() error : %v", err)
		return nil, response.TransactionFailed
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	ds, amount, err := repo.calculate(context, transaction, tx)
	if err != nil {
		log.Printf("transactionRepository.Create() - repo.db.BeginTx() error : %v", err)
		return nil, response.TransactionFailed
	}

	now := time.Now()
	result, err := tx.ExecContext(context, createTransaction, transaction.CustomerId, now, amount)
	if err != nil {
		log.Printf("transactionRepository.Create() - tx.ExecContext() createTransaction error : %v", err)
		return nil, response.TransactionFailed
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("transactionRepository.Create() - result.LastInsertId() createTransaction error : %v", err)
		return nil, response.TransactionFailed
	}

	transactionId := int(id)
	details := []entity.TransactionDetailEntity{}
	for _, order := range ds {
		result, err := tx.ExecContext(context, createTransactionDetail, transactionId, order.ProductId, order.TotalOrder, order.TotalPrice)
		if err != nil {
			log.Printf("transactionRepository.Create() - tx.ExecContext() createTransaction error : %v", err)
			return nil, response.TransactionFailed
		}
		orderId, err := result.LastInsertId()
		if err != nil {
			log.Printf("transactionRepository.Create() - result.LastInsertId() createTransaction error : %v", err)
			return nil, response.TransactionFailed
		}
		detail := entity.TransactionDetailEntity{
			Id:         int(orderId),
			ProductId:  order.ProductId,
			TotalOrder: order.TotalOrder,
			Price:      order.Price,
			TotalPrice: order.TotalPrice,
		}
		details = append(details, detail)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, response.TransactionFailed
	}

	return &entity.TransactionEntity{
		BaseEntity: entity.NewBaseEntity(int64(transactionId), now),
		CustomerId: transaction.CustomerId,
		Amount:     amount,
		Orders:     details,
	}, nil

}

func (repo *transactionRepository) calculate(context context.Context, transaction *entity.CreatedTransactionEntity, tx *sql.Tx) (map[int]entity.CalculatedTransaction, int, error) {
	stmt, err := tx.PrepareContext(context, productPrice)
	if err != nil {
		log.Printf("transactionRepository.calculate() - tx.PrepareContext error : %v", err)
		return nil, 0, response.TransactionProductNotFound
	}
	defer stmt.Close()

	var ds map[int]entity.CalculatedTransaction = map[int]entity.CalculatedTransaction{}
	amount := 0
	for _, order := range transaction.Orders {
		var price *int

		err := stmt.QueryRow(order.ProductId).Scan(&price)
		if err == nil {
			ds[order.ProductId] = entity.CalculatedTransaction{
				ProductId:  order.ProductId,
				TotalOrder: order.TotalOrder,
				TotalPrice: order.TotalOrder * *price,
			}
			amount += order.TotalOrder * *price
		} else {
			return nil, 0, response.TransactionProductNotFound
		}

	}
	return ds, amount, nil
}
