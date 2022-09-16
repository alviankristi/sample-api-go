package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alviankristi/catalyst-backend-task/pkg/database"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
	"github.com/stretchr/testify/suite"
)

type BrandRepositoryTestSuite struct {
	suite.Suite
	mock            sqlmock.Sqlmock
	db              *sql.DB
	brandRepository BrandRepository
}

func TestBrandRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BrandRepositoryTestSuite))
}

func (t *BrandRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	t.Nil(err)
	t.db = db
	t.mock = mock
	t.brandRepository = NewBrandRepository(db)

}

func (t *BrandRepositoryTestSuite) TestBrandNameDuplicateSql() {
	s := `SELECT 1 FROM brands where LOWER(name) = LOWER(?)`
	t.mock.ExpectQuery(regexp.QuoteMeta(s)).WithArgs("name").WillReturnError(sql.ErrNoRows)
	t.db.Exec(regexp.QuoteMeta(brandNameExist), "name")

	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *BrandRepositoryTestSuite) TestCreateBrandSql() {
	time := time.Now()
	sql := `INSERT INTO brands (name, created_date) VALUES (?,?)`
	t.mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs("name", time).WillReturnResult(sqlmock.NewResult(1, 1))
	t.db.Exec(regexp.QuoteMeta(createBrand), "name", time)

	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *BrandRepositoryTestSuite) TestCreateSuccess() {
	ctx := context.Background()
	t.mock.ExpectQuery(regexp.QuoteMeta(brandNameExist)).WithArgs("name").WillReturnError(sql.ErrNoRows)
	t.mock.ExpectExec(regexp.QuoteMeta(createBrand)).WithArgs("name", database.AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))
	result, err := t.brandRepository.Create(ctx, "name")
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

func (t *BrandRepositoryTestSuite) TestCreateFailedExecute() {
	ctx := context.Background()
	t.mock.ExpectQuery(regexp.QuoteMeta(brandNameExist)).WithArgs("name").WillReturnError(sql.ErrNoRows)
	t.mock.ExpectExec(regexp.QuoteMeta(createBrand)).WithArgs("name", database.AnyTime{}).WillReturnError(errors.New("Error"))
	result, err := t.brandRepository.Create(ctx, "name")

	t.Nil(result)
	if err != nil {
		t.Error(err)
	}
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *BrandRepositoryTestSuite) TestCreateFailedLastInsertId() {
	ctx := context.Background()
	t.mock.ExpectQuery(regexp.QuoteMeta(brandNameExist)).WithArgs("name").WillReturnError(sql.ErrNoRows)
	t.mock.ExpectExec(regexp.QuoteMeta(createBrand)).WithArgs("name", database.AnyTime{}).WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))
	result, err := t.brandRepository.Create(ctx, "name")

	t.Nil(result)
	if err != nil {
		t.Error(err)
	}
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func (t *BrandRepositoryTestSuite) TestCreateDuplicateBrandNameFailed() {
	ctx := context.Background()
	t.mock.ExpectQuery(regexp.QuoteMeta(brandNameExist)).WithArgs("name").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	t.mock.ExpectExec(regexp.QuoteMeta(createBrand)).WithArgs("name", database.AnyTime{}).WillReturnError(errors.New("Error"))
	result, err := t.brandRepository.Create(ctx, "name")
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	t.Nil(result)
	t.ErrorIs(err, response.BrandNameDuplicate)
}

func (t *BrandRepositoryTestSuite) TestCreateFailed() {
	ctx := context.Background()
	t.mock.ExpectExec(regexp.QuoteMeta(createBrand)).WithArgs("name", database.AnyTime{}).WillReturnError(errors.New("Error"))
	result, err := t.brandRepository.Create(ctx, "name")
	if err := t.mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	t.Nil(result)
	t.NotNil(err)
}
