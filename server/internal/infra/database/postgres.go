package database

import (
	"database/sql"

	"github.com/m-bromo/rolldeck/config"
)

const PostgresDriver = "postgres"

func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(PostgresDriver, cfg.PostgresDB.DSN)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
