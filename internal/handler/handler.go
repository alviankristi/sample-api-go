package handler

import (
	"database/sql"
	"net/http"

	"github.com/alviankristi/catalyst-backend-task/internal/repository"
	"github.com/alviankristi/catalyst-backend-task/internal/service"
)

type ApiHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

//Create API Handler
func NewHandler(db *sql.DB) *http.ServeMux {
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	mux := http.NewServeMux()
	mux.Handle("/brand", NewBrandHandler(service))
	return mux
}
