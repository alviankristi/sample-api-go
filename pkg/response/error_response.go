package response

import "errors"

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
	ValidationError    = &AppErrorResponse{Message: "validation error", ApplicationCode: 1}
	DatabaseError      = &AppErrorResponse{Message: "database error", ApplicationCode: 2}
	BrandNameDuplicate = &AppErrorResponse{Message: "duplicate brand name", ApplicationCode: 3}
)
