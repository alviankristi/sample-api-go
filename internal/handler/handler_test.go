package handler

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (t *HandlerTestSuite) SetupTest() {}

func (t *HandlerTestSuite) TestNewHandler() {
	db, _, _ := sqlmock.New()
	handler := NewHandler(db)

	t.NotNil(handler)
}
