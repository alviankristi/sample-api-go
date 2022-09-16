package response

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// error response
type AppErrorResponse struct {
	Message         string
	ApplicationCode int
}

func (a *AppErrorResponse) ThrowError() error {
	return errors.New(a.Message)
}

func (a *AppErrorResponse) Error() string {
	return a.Message
}

var (
	ValidationError          = &AppErrorResponse{Message: "validation error", ApplicationCode: 1}
	DatabaseError            = &AppErrorResponse{Message: "database error", ApplicationCode: 2}
	BrandNameDuplicate       = &AppErrorResponse{Message: "duplicate brand name", ApplicationCode: 3}
	InvalidDecodeRequestBody = &AppErrorResponse{Message: "decode request body failed", ApplicationCode: 4}
)

type ErrResponse struct {
	Err            error    `json:"-"`
	HTTPStatusCode int      `json:"-"`
	StatusText     string   `json:"status"`
	Errors         []string `json:"errors,omitempty"`
	AppCode        int      `json:"code,omitempty"`
	ErrorText      string   `json:"error,omitempty"`
}

// error response for validation
func ErrValidationRequest(err validator.ValidationErrors) *ErrResponse {
	e := ValidationError
	var messages []string
	for _, e := range err {
		s := fmt.Sprintf("%v: %s", e.Field(), e.Tag())
		messages = append(messages, s)
	}

	return &ErrResponse{
		Err:            e.ThrowError(),
		HTTPStatusCode: 400,
		StatusText:     ValidationError.Message,
		AppCode:        e.ApplicationCode,
		ErrorText:      e.Error(),
		Errors:         messages,
	}
}

// error response for handled error
func ErrInvalidRequest(err *AppErrorResponse) *ErrResponse {
	return &ErrResponse{
		Err:            err.ThrowError(),
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		AppCode:        err.ApplicationCode,
		ErrorText:      err.Error(),
	}
}

// error response for not handled error
func InternalServerError(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "System failure",
		AppCode:        0,
		ErrorText:      err.Error(),
	}
}

//convert error to specific error response model
func ErrRender(err error) *ErrResponse {
	switch t := err.(type) {

	case *AppErrorResponse:
		return ErrInvalidRequest(t)
	case validator.ValidationErrors:
		return ErrValidationRequest(t)
	default:
		return InternalServerError(err)
	}

}
