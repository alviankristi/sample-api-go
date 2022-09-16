package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	"github.com/alviankristi/catalyst-backend-task/pkg/database"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
	"github.com/stretchr/testify/suite"
)

type TransactionRepositoryTestSuite struct {
	suite.Suite
	mock                  sqlmock.Sqlmock
	db                    *sql.DB
	transactionRepository TransactionRespository
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}

func (t *TransactionRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	t.Nil(err)
	t.db = db
	t.mock = mock
	t.transactionRepository = NewTransactionRepository(db)

}

func (t *TransactionRepositoryTestSuite) TestGetByIdSuccess() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"id", "customer_id ", "fullname", "amount", "created_date", "id", "transaction_id", "transaction_id",
		"total_order", "total_price", "name", "name", "id"}).AddRow(1, 1, "customer name", 1000, time.Now(), 1, 1, 1, 1, 1, "name", "name", 1).
		AddRow(1, 1, "customer name", 1000, time.Now(), 1, 1, 1, 1, 1, "name", "name", 1)
	t.mock.ExpectQuery(regexp.QuoteMeta(getTransactionById)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.transactionRepository.GetById(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Equal(result.Id, 1)

	t.Nil(err)
}

func (t *TransactionRepositoryTestSuite) TestGetByTransactionIdEmpty() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"id", "customer_id ", "fullname", "amount", "created_date", "id", "transaction_id", "transaction_id",
		"total_order", "total_price", "name", "name", "id"})

	t.mock.ExpectQuery(regexp.QuoteMeta(getTransactionById)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.transactionRepository.GetById(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Empty(result)
	t.Nil(err)
}

func (t *TransactionRepositoryTestSuite) TestGetByTransactionIdScanError() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"customer_id ", "id", "fullname", "amount", "created_date", "id", "transaction_id", "transaction_id",
		"total_order", "total_price", "name", "name", "id"}).AddRow("asd", "dsa", "customer name", 1000, 1, 1, 1, 1, 1, 1, "name", "name", 1).
		AddRow("asd", "dsa", "customer name", 1000, 1, 1, 1, 1, 1, 1, "name", "name", 1)

	t.mock.ExpectQuery(regexp.QuoteMeta(getTransactionById)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.transactionRepository.GetById(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.NotNil(result)
	t.Nil(err)
}

func (t *TransactionRepositoryTestSuite) TestGetByTransactionIdRowNextError() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"id", "customer_id ", "fullname", "amount", "created_date", "id", "transaction_id", "transaction_id",
		"total_order", "total_price", "name", "name", "id"}).AddRow(1, 1, "customer name", 1000, time.Now(), 1, 1, 1, 1, 1, "name", "name", 1).
		AddRow(2, 1, "customer name", 1000, time.Now(), 1, 1, 1, 1, 1, "name", "name", 1).
		AddRow(2, 1, "customer name", 1000, time.Now(), 1, 1, 1, 1, 1, "name", "name", 1).
		RowError(2, errors.New("error"))

	t.mock.ExpectQuery(regexp.QuoteMeta(getTransactionById)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.transactionRepository.GetById(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Nil(result)
	t.NotNil(err)
}

func (t *TransactionRepositoryTestSuite) TestGetByTransactionIdDbError() {
	ctx := context.Background()

	t.mock.ExpectQuery(regexp.QuoteMeta(getTransactionById)).WithArgs(1).WillReturnError(sql.ErrConnDone)

	result, err := t.transactionRepository.GetById(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Nil(result)
	t.Equal(response.DatabaseError, err)
}

func (t *TransactionRepositoryTestSuite) TestCreateSuccess() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"price"}).AddRow(10)
	t.mock.ExpectBegin()
	t.mock.ExpectPrepare(productPrice).ExpectQuery().WillReturnRows(mockedRow)
	t.mock.ExpectExec(regexp.QuoteMeta(createTransaction)).WithArgs(1, database.AnyTime{}, 10).WillReturnResult(sqlmock.NewResult(1, 1))
	t.mock.ExpectExec(regexp.QuoteMeta(createTransactionDetail)).WithArgs(1, 1, 1, 10).WillReturnResult(sqlmock.NewResult(1, 1))
	t.mock.ExpectCommit()
	x := &entity.CreatedTransactionEntity{
		CustomerId: 1,
		Orders:     []entity.CreatedTransactionDetailEntity{{ProductId: 1, TotalOrder: 1}},
	}
	result, err := t.transactionRepository.Create(ctx, x)
	t.Equal(result.Id, 1)
	t.NotEmpty(result.Id)
	t.NotEmpty(result.CreatedDate)
	if err != nil {
		t.Error(err)
	}
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
