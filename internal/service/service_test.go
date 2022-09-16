package service

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alviankristi/catalyst-backend-task/internal/repository"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (t *ServiceTestSuite) SetupTest() {}

func (t *ServiceTestSuite) TestNewHandler() {
	db, _, _ := sqlmock.New()
	repo := repository.NewRepository(db)
	service := NewService(repo)
	t.NotNil(service)
}
