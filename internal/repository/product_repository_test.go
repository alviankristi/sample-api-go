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

type ProductRepositoryTestSuite struct {
	suite.Suite
	mock              sqlmock.Sqlmock
	db                *sql.DB
	productRepository ProductRepository
}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}

func (t *ProductRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	t.Nil(err)
	t.db = db
	t.mock = mock
	t.productRepository = NewProductRepository(db)

}

func (t *ProductRepositoryTestSuite) TestProductNameDuplicateSql() {
	s := `SELECT 1 FROM products where LOWER(name) = LOWER(?) and brand_id = ?`
	t.mock.ExpectQuery(regexp.QuoteMeta(s)).WithArgs("name", 1).WillReturnError(sql.ErrNoRows)
	t.db.Exec(regexp.QuoteMeta(productNameExist), "name", 1)

	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *ProductRepositoryTestSuite) TestCreateProductSql() {
	time := time.Now()
	sql := `INSERT INTO products (name, brand_id, price, created_date) VALUES (?, ?, ?, ?)`
	t.mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs("name", 1, 10, database.AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))
	t.db.Exec(regexp.QuoteMeta(createProduct), "name", time)

	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *ProductRepositoryTestSuite) TestGetByBrandIdSql() {
	sql := `SELECT id, created_date, brand_id ,name, price FROM products where brand_id = ?`
	mockedRow := sqlmock.NewRows([]string{"id", "name", "brand_id", "price"}).AddRow(1, "name", 1, 10)

	t.mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(1).WillReturnRows(mockedRow)
	t.db.Exec(regexp.QuoteMeta(getByBrandId), 1)

	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *ProductRepositoryTestSuite) TestGetByBrandIdSuccess() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"id", "created_date ", "brand_id", "name", "price"}).AddRow(1, time.Now(), 1, "name", 10)
	t.mock.ExpectQuery(regexp.QuoteMeta(getByBrandId)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.productRepository.GetByBrandId(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Equal(result[0].Id, 1)
	t.Equal(result[0].Name, "name")
	t.Equal(result[0].Price, 10)
	t.Equal(result[0].BrandId, 1)
	t.Nil(err)
}

func (t *ProductRepositoryTestSuite) TestGetByBrandIdEmpty() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"id", "created_date ", "brand_id", "name", "price"})

	t.mock.ExpectQuery(regexp.QuoteMeta(getByBrandId)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.productRepository.GetByBrandId(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Empty(result)
	t.Nil(err)
}

func (t *ProductRepositoryTestSuite) TestGetByBrandIdScanError() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"created_date ", "id", "brand_id", "name", "price"}).AddRow(time.Now(), 1, 1, "name", 10)

	t.mock.ExpectQuery(regexp.QuoteMeta(getByBrandId)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.productRepository.GetByBrandId(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Empty(result)
	t.Nil(err)
}

func (t *ProductRepositoryTestSuite) TestGetByBrandIdRowNextError() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"id", "created_date ", "brand_id", "name", "price"}).AddRow(time.Now(), 1, 1, "name", 10).RowError(0, errors.New("error"))

	t.mock.ExpectQuery(regexp.QuoteMeta(getByBrandId)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.productRepository.GetByBrandId(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Empty(result)
	t.NotNil(err)
}

func (t *ProductRepositoryTestSuite) TestGetByBrandIdDbError() {
	ctx := context.Background()

	t.mock.ExpectQuery(regexp.QuoteMeta(getByBrandId)).WithArgs(1).WillReturnError(sql.ErrConnDone)

	result, err := t.productRepository.GetByBrandId(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Nil(result)
	t.Equal(response.DatabaseError, err)
}

func (t *ProductRepositoryTestSuite) TestGetByIdSql() {
	sql := `SELECT id, created_date, brand_id ,name, price FROM products where id = ?`
	mockedRow := sqlmock.NewRows([]string{"id", "name", "brand_id", "price"}).AddRow(1, "name", 1, 10)

	t.mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(1).WillReturnRows(mockedRow)
	t.db.Exec(regexp.QuoteMeta(getSingleProduct), 1)

	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *ProductRepositoryTestSuite) TestGetByIdSuccess() {
	ctx := context.Background()
	mockedRow := sqlmock.NewRows([]string{"id", "created_date ", "brand_id", "name", "price"}).AddRow(1, time.Now(), 1, "name", 10)
	t.mock.ExpectQuery(regexp.QuoteMeta(getSingleProduct)).WithArgs(1).WillReturnRows(mockedRow)

	result, err := t.productRepository.GetById(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Equal(result.Id, 1)
	t.Equal(result.Name, "name")
	t.Equal(result.Price, 10)
	t.Equal(result.BrandId, 1)
	t.Nil(err)
}

func (t *ProductRepositoryTestSuite) TestGetByIdEmpty() {
	ctx := context.Background()
	t.mock.ExpectQuery(regexp.QuoteMeta(getSingleProduct)).WithArgs(1).WillReturnError(sql.ErrNoRows)

	result, err := t.productRepository.GetById(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Nil(result)
	t.Nil(err)
}

func (t *ProductRepositoryTestSuite) TestGetByIdDBError() {
	ctx := context.Background()
	t.mock.ExpectQuery(regexp.QuoteMeta(getSingleProduct)).WithArgs(1).WillReturnError(sql.ErrConnDone)

	result, err := t.productRepository.GetById(ctx, 1)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	t.Nil(result)
	t.NotNil(err)
}

func (t *ProductRepositoryTestSuite) TestCreateSuccess() {
	ctx := context.Background()
	product := entity.CreatedProductEntity{
		Name:    "name",
		BrandId: 1,
		Price:   10,
	}
	t.mock.ExpectQuery(regexp.QuoteMeta(productNameExist)).WithArgs("name", 1).WillReturnError(sql.ErrNoRows)
	t.mock.ExpectExec(regexp.QuoteMeta(createProduct)).WithArgs("name", 1, 10, database.AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))
	result, err := t.productRepository.Create(ctx, product)
	t.Equal(result.Name, "name")
	t.NotEmpty(result.Id)
	t.NotEmpty(result.CreatedDate)
	if err != nil {
		t.Error(err)
	}
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *ProductRepositoryTestSuite) TestCreateFailedExecute() {
	ctx := context.Background()
	product := entity.CreatedProductEntity{
		Name:    "name",
		BrandId: 1,
		Price:   10,
	}
	t.mock.ExpectQuery(regexp.QuoteMeta(productNameExist)).WithArgs("name", 1).WillReturnError(sql.ErrNoRows)
	t.mock.ExpectExec(regexp.QuoteMeta(createProduct)).WithArgs("name", 1, 10, database.AnyTime{}).WillReturnError(errors.New("Error"))
	result, err := t.productRepository.Create(ctx, product)

	t.Nil(result)
	if err != nil {
		t.Error(err)
	}
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *ProductRepositoryTestSuite) TestCreateFailedLastInsertId() {
	ctx := context.Background()
	product := entity.CreatedProductEntity{
		Name:    "name",
		BrandId: 1,
		Price:   10,
	}
	t.mock.ExpectQuery(regexp.QuoteMeta(productNameExist)).WithArgs("name", 1).WillReturnError(sql.ErrNoRows)
	t.mock.ExpectExec(regexp.QuoteMeta(createProduct)).WithArgs("name", 1, 10, database.AnyTime{}).WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))
	result, err := t.productRepository.Create(ctx, product)

	t.Nil(result)
	if err != nil {
		t.Error(err)
	}
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *ProductRepositoryTestSuite) TestCreateDuplicateProductNameFailed() {
	ctx := context.Background()
	product := entity.CreatedProductEntity{
		Name:    "name",
		BrandId: 1,
		Price:   10,
	}
	t.mock.ExpectQuery(regexp.QuoteMeta(productNameExist)).WithArgs("name", 1).WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	t.mock.ExpectExec(regexp.QuoteMeta(createProduct)).WithArgs("name", 1, 10, database.AnyTime{}).WillReturnError(errors.New("Error"))
	result, err := t.productRepository.Create(ctx, product)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	t.Nil(result)
	t.ErrorIs(err, response.BrandProductNameDuplicate)
}

func (t *ProductRepositoryTestSuite) TestCreateDuplicateProductNameDatabaseError() {
	ctx := context.Background()
	product := entity.CreatedProductEntity{
		Name:    "name",
		BrandId: 1,
		Price:   10,
	}
	t.mock.ExpectQuery(regexp.QuoteMeta(productNameExist)).WithArgs("name", 1).WillReturnError(sql.ErrConnDone)
	t.mock.ExpectExec(regexp.QuoteMeta(createProduct)).WithArgs("name", 1, 10, database.AnyTime{}).WillReturnError(errors.New("Error"))
	result, err := t.productRepository.Create(ctx, product)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	t.Nil(result)
	t.ErrorIs(err, response.DatabaseError)
}

func (t *ProductRepositoryTestSuite) TestCreateFailed() {
	ctx := context.Background()
	product := entity.CreatedProductEntity{
		Name:    "name",
		BrandId: 1,
		Price:   10,
	}
	t.mock.ExpectQuery(regexp.QuoteMeta(productNameExist)).WithArgs("name", 1).WillReturnError(sql.ErrNoRows)
	t.mock.ExpectExec(regexp.QuoteMeta(createProduct)).WithArgs("name", 1, 10, database.AnyTime{}).WillReturnError(errors.New("Error"))
	result, err := t.productRepository.Create(ctx, product)
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	t.Nil(result)
	t.NotNil(err)
}
