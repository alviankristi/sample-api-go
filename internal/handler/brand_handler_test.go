package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alviankristi/catalyst-backend-task/internal/service"
	mockService "github.com/alviankristi/catalyst-backend-task/internal/service/mock"
	"github.com/alviankristi/catalyst-backend-task/internal/service/model"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
	errReader "github.com/alviankristi/catalyst-backend-task/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BrandHandlerTestSuite struct {
	suite.Suite
	handler ApiHandler
}

func TestBrandHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BrandHandlerTestSuite))
}

func (t *BrandHandlerTestSuite) SetupTest() {
	response := &model.BrandResponseModel{
		Name: "name",
		Id:   1,
	}
	brandService := new(mockService.BrandServiceMock)
	brandService.On("Create", mock.Anything, mock.Anything).Return(response, nil)
	mockService := &service.Service{
		BrandService: brandService,
	}

	t.handler = NewBrandHandler(mockService)
}

func (t *BrandHandlerTestSuite) TestCreateHandler() {
	body := &model.BrandModel{
		Name: "name",
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, err := http.NewRequest("POST", "/brand", &buf)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusCreated)

	t.NotNil(rr.Body)

}

func (t *BrandHandlerTestSuite) TestCreateHandlerErrorReader() {

	body := &model.BrandModel{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, err := http.NewRequest("POST", "/brand", errReader.ErrReader(0))
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusBadRequest)
	t.NotNil(rr.Body)
}

func (t *BrandHandlerTestSuite) TestCreateHandlerValidationError() {

	body := &model.BrandModel{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, err := http.NewRequest("POST", "/brand", &buf)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusBadRequest)
	t.NotNil(rr.Body)
}

func (t *BrandHandlerTestSuite) TestCreateHandlerServiceBrandNameDuplicateError() {
	brandService := new(mockService.BrandServiceMock)
	brandService.On("Create", mock.Anything, mock.Anything).Return(nil, response.BrandNameDuplicate)
	mockService := &service.Service{
		BrandService: brandService,
	}

	t.handler = NewBrandHandler(mockService)

	body := &model.BrandModel{
		Name: "name",
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, err := http.NewRequest("POST", "/brand", &buf)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusBadRequest)
	t.NotNil(rr.Body)
}

func (t *BrandHandlerTestSuite) TestBrandHandlerNotFoundEndpoint() {

	req, err := http.NewRequest("Get", "/brand/abcd", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusNotFound)
	t.NotNil(rr.Body)
}
