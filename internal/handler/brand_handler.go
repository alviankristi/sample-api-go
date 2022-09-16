package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/alviankristi/catalyst-backend-task/internal/service"
	"github.com/alviankristi/catalyst-backend-task/internal/service/model"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
	"github.com/go-playground/validator/v10"
)

// regex fot matching the endpoint
var (
	createBrandRe = regexp.MustCompile(`^\/brand[\/]*$`)
)

type brandHandler struct {
	service *service.Service
}

//Create Brand Handler
func NewBrandHandler(service *service.Service) ApiHandler {
	return &brandHandler{
		service: service,
	}
}

func (h *brandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	// checking url path suit the regex
	switch {
	case r.Method == http.MethodPost && createBrandRe.MatchString(r.URL.Path):
		h.create(w, r)
		return
	default:
		log.Printf("brandHandler.ServeHTTP() - not found : %v", r.URL.Path)
		response.NotFound(w)
		return
	}
}

func (h *brandHandler) create(w http.ResponseWriter, r *http.Request) {

	//convert the request body to struct
	var model *model.BrandModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("brandHandler.create() - error decode request body : %v", err)
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
	result, err := h.service.BrandService.Create(r.Context(), *model)
	if err != nil {
		response.BadRequest(w, response.ErrRender(err))
		return
	}

	response.Created(w, result)
}
