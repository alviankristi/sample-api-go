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

type ProductHandlerTestSuite struct {
	suite.Suite
	handler ApiHandler
}

func TestProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}

func (t *ProductHandlerTestSuite) SetupTest() {
	response := &model.ProductResponseModel{
		Name:    "name",
		Id:      1,
		BrandId: 1,
		Price:   10,
	}
	productService := new(mockService.ProductServiceMock)
	productService.On("Create", mock.Anything, mock.Anything).Return(response, nil)
	mockService := &service.Service{
		ProductService: productService,
	}

	t.handler = NewProductHandler(mockService)
}

func (t *ProductHandlerTestSuite) TestGetByBrandHandlerQueryParamIdUnset() {
	productService := new(mockService.ProductServiceMock)
	response := []*model.ProductResponseModel{{Name: "name", Id: 1, Price: 10, BrandId: 1}}
	productService.On("GetByBrandId", mock.Anything, mock.Anything).Return(response, nil)
	mockService := &service.Service{
		ProductService: productService,
	}

	t.handler = NewProductHandler(mockService)

	req, err := http.NewRequest("GET", "/product/brand", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusNotFound)
	t.NotNil(rr.Body)
}

func (t *ProductHandlerTestSuite) TestGetByBrandHandlerSuccess() {
	productService := new(mockService.ProductServiceMock)
	response := []*model.ProductResponseModel{{Name: "name", Id: 1, Price: 10, BrandId: 1}}
	productService.On("GetByBrandId", mock.Anything, mock.Anything).Return(response, nil)
	mockService := &service.Service{
		ProductService: productService,
	}

	t.handler = NewProductHandler(mockService)

	req, err := http.NewRequest("GET", "/product/brand?id=1", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusOK)
	t.NotNil(rr.Body)
}

func (t *ProductHandlerTestSuite) TestGetByBrandHandlerError() {
	productService := new(mockService.ProductServiceMock)
	productService.On("GetByBrandId", mock.Anything, mock.Anything).Return(nil, response.DatabaseError)
	mockService := &service.Service{
		ProductService: productService,
	}

	t.handler = NewProductHandler(mockService)

	req, err := http.NewRequest("GET", "/product/brand?id=1", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusBadRequest)
	t.NotNil(rr.Body)
}

func (t *ProductHandlerTestSuite) TestGetHandlerQueryParamIdUnset() {
	productService := new(mockService.ProductServiceMock)
	response := &model.ProductResponseModel{Name: "name", Id: 1, Price: 10, BrandId: 1}
	productService.On("GetById", mock.Anything, mock.Anything).Return(response)
	mockService := &service.Service{
		ProductService: productService,
	}

	t.handler = NewProductHandler(mockService)

	req, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusNotFound)
	t.NotNil(rr.Body)
}

func (t *ProductHandlerTestSuite) TestGetHandlerDbError() {

	productService := new(mockService.ProductServiceMock)
	productService.On("GetById", mock.Anything, mock.Anything).Return(nil, response.DatabaseError)
	mockService := &service.Service{
		ProductService: productService,
	}

	t.handler = NewProductHandler(mockService)

	req, err := http.NewRequest("GET", "/product?id=1", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)
	t.Equal(rr.Code, http.StatusBadRequest)
	t.NotNil(rr.Body)
}
func (t *ProductHandlerTestSuite) TestGetHandlerSuccess() {

	productService := new(mockService.ProductServiceMock)
	response := &model.ProductResponseModel{Name: "name", Id: 1, Price: 10, BrandId: 1}
	productService.On("GetById", mock.Anything, mock.Anything).Return(response, nil)
	mockService := &service.Service{
		ProductService: productService,
	}

	t.handler = NewProductHandler(mockService)

	req, err := http.NewRequest("GET", "/product?id=1", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)
	t.Equal(rr.Code, http.StatusOK)
	t.NotNil(rr.Body)
}

func (t *ProductHandlerTestSuite) TestCreateHandler() {
	body := &model.ProductModel{
		Name:    "name",
		BrandId: 1,
		Price:   10,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, err := http.NewRequest("POST", "/product", &buf)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusCreated)

	t.NotNil(rr.Body)

}

func (t *ProductHandlerTestSuite) TestCreateHandlerErrorReader() {

	body := &model.ProductModel{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, err := http.NewRequest("POST", "/product", errReader.ErrReader(0))
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusBadRequest)
	t.NotNil(rr.Body)
}

func (t *ProductHandlerTestSuite) TestCreateHandlerValidationError() {

	body := &model.ProductModel{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, err := http.NewRequest("POST", "/product", &buf)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusBadRequest)
	t.NotNil(rr.Body)
}

func (t *ProductHandlerTestSuite) TestCreateHandlerServiceProductNameDuplicateError() {
	productService := new(mockService.ProductServiceMock)
	productService.On("Create", mock.Anything, mock.Anything).Return(nil, response.BrandProductNameDuplicate)
	mockService := &service.Service{
		ProductService: productService,
	}

	t.handler = NewProductHandler(mockService)

	body := &model.ProductModel{
		Name:    "name",
		BrandId: 1,
		Price:   10,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, err := http.NewRequest("POST", "/product", &buf)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusBadRequest)
	t.NotNil(rr.Body)
}

func (t *ProductHandlerTestSuite) TestProductHandlerNotFoundEndpoint() {

	req, err := http.NewRequest("Get", "/product/abcd", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(t.handler)
	handler.ServeHTTP(rr, req)

	t.Equal(rr.Code, http.StatusNotFound)
	t.NotNil(rr.Body)
}
