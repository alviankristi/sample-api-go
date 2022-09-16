package server

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/alviankristi/catalyst-backend-task/config"
	"github.com/alviankristi/catalyst-backend-task/internal/handler"
)

type Server struct {
	*http.Server
	*config.Config
}

// create api server
func NewServer(config *config.Config, db *sql.DB) (*Server, error) {

	if config.Api.Hostname == "" {
		return nil, errors.New("config.Api.Hostname is empty")
	}
	if config.Api.Port == "" {
		return nil, errors.New("config.Api.Port is empty")
	}

	apiHandler := handler.NewHandler(db)

	srv := &http.Server{
		Addr:    config.Api.Hostname + ":" + config.Api.Port,
		Handler: apiHandler,
	}

	return &Server{srv, config}, nil
}
