package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (t *RepositoryTestSuite) SetupTest() {}

func (t *RepositoryTestSuite) TestNewHandler() {
	db, _, _ := sqlmock.New()
	repo := NewRepository(db)

	t.NotNil(repo)
}
