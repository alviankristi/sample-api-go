package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/alviankristi/catalyst-backend-task/config"
)

func Open(config *config.Config) *sql.DB {
	db, err := sql.Open("mysql", config.DbConfig.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Duration(config.DbConfig.MaxConnLifetime))
	db.SetMaxIdleConns(config.DbConfig.MaxIdleConn)
	db.SetMaxOpenConns(config.DbConfig.MaxOpenConn)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed open connection: %v", err)
	}
	return db
}
