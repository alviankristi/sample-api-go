package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/alviankristi/catalyst-backend-task/internal/service"
	"github.com/alviankristi/catalyst-backend-task/internal/service/model"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
	"github.com/go-playground/validator/v10"
)

// regex fot matching the endpoint
var (
	productRe      = regexp.MustCompile(`^\/product[\/]*$`)
	brandProductRe = regexp.MustCompile(`^\/product\/brand[\/]*$`)
)

type productHandler struct {
	service *service.Service
}

//Create Product Handler
func NewProductHandler(service *service.Service) ApiHandler {
	return &productHandler{
		service: service,
	}
}

func (h *productHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	// checking url path suit the regex
	switch {
	case r.Method == http.MethodPost && productRe.MatchString(r.URL.Path):
		h.create(w, r)
		return
	case r.Method == http.MethodGet && productRe.MatchString(r.URL.Path):
		h.get(w, r)
		return
	case r.Method == http.MethodGet && brandProductRe.MatchString(r.URL.Path):
		h.getByBrandId(w, r)
		return
	default:
		log.Printf("productHandler.ServeHTTP() - not found : %v", r.URL.Path)
		response.NotFound(w)
		return
	}
}

func (h *productHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("id")
	if queryParam == "" {
		log.Printf("productHandler.get() - error : id not found")
		response.NotFound(w)
		return
	}

	id, _ := strconv.Atoi(queryParam)

	result, err := h.service.ProductService.GetById(r.Context(), id)
	if err != nil {
		response.BadRequest(w, response.ErrRender(err))
		return
	}

	response.Ok(w, result)

}

func (h *productHandler) getByBrandId(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query().Get("id")
	if queryParam == "" {
		log.Printf("productHandler.getByBrandId() - error : id not found")
		response.NotFound(w)
		return
	}

	id, _ := strconv.Atoi(queryParam)

	result, err := h.service.ProductService.GetByBrandId(r.Context(), id)
	if err != nil {
		response.BadRequest(w, response.ErrRender(err))
		return
	}

	response.Ok(w, result)

}

func (h *productHandler) create(w http.ResponseWriter, r *http.Request) {

	//convert the request body to struct
	var model *model.ProductModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("productHandler.create() - error decode request body : %v", err)
		response.RenderInvalidDecodeRequestBody(w)
		return
	}

	//validate struct value
	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		response.BadRequest(w, response.ErrRender(err))
		return
	}

	//call service
	result, err := h.service.ProductService.Create(r.Context(), *model)
	if err != nil {
		response.BadRequest(w, response.ErrRender(err))
		return
	}

	response.Created(w, result)
}
