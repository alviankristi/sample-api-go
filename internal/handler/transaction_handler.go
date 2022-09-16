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
	createTransactionRe = regexp.MustCompile(`^\/order[\/]*$`)
)

type transactionHandler struct {
	service *service.Service
}

//Create Transaction Handler
func NewTransactionHandler(service *service.Service) ApiHandler {
	return &transactionHandler{
		service: service,
	}
}

func (h *transactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	// checking url path suit the regex
	switch {
	case r.Method == http.MethodPost && createTransactionRe.MatchString(r.URL.Path):
		h.create(w, r)
		return
	case r.Method == http.MethodGet && createTransactionRe.MatchString(r.URL.Path):
		h.get(w, r)
		return
	default:
		log.Printf("transactionHandler.ServeHTTP() - not found : %v", r.URL.Path)
		response.NotFound(w)
		return
	}
}

func (h *transactionHandler) create(w http.ResponseWriter, r *http.Request) {

	//convert the request body to struct
	var model *model.TransactionModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("transactionHandler.create() - error decode request body : %v", err)
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
	result, err := h.service.TransactionService.Create(r.Context(), *model)
	if err != nil {
		response.BadRequest(w, response.ErrRender(err))
		return
	}

	response.Created(w, result)
}

func (h *transactionHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("transactionID")
	if queryParam == "" {
		log.Printf("productHandler.get() - error : id not found")
		response.NotFound(w)
		return
	}

	id, _ := strconv.Atoi(queryParam)

	result, err := h.service.TransactionService.GetById(r.Context(), id)
	if err != nil {
		response.BadRequest(w, response.ErrRender(err))
		return
	}

	response.Ok(w, result)

}
