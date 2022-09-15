package service

import (
	"context"
	"testing"
	"time"

	"github.com/alviankristi/catalyst-backend-task/internal/repository"
	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	"github.com/alviankristi/catalyst-backend-task/internal/service/model"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BrandServiceTestSuite struct {
	suite.Suite
	brandService        BrandService
	brandRepositoryMock *repository.BrandRepositoryMock
}

func TestBrandServiceTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(BrandServiceTestSuite))
}

func (t *BrandServiceTestSuite) SetupTest() {
	t.brandRepositoryMock = new(repository.BrandRepositoryMock)

	t.brandService = NewBrandService(t.brandRepositoryMock)
}

func (t *BrandServiceTestSuite) TestCreateBrandSuccess() {
	e := &entity.BrandEntity{
		Name:       "name",
		BaseEntity: entity.NewBaseEntity(1, time.Now()),
	}
	t.brandRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(e, nil)
	result, err := t.brandService.Create(context.Background(), model.BrandModel{
		Name: "name",
	})

	t.Nil(err)
	t.NotEmpty(result.Id)
	t.NotEmpty(result.Name)
}

func (t *BrandServiceTestSuite) TestCreateBrandFailed() {

	t.brandRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(nil, response.DatabaseError)
	result, err := t.brandService.Create(context.Background(), model.BrandModel{
		Name: "name",
	})

	t.Equal(err, response.DatabaseError)
	t.Nil(result)
}
