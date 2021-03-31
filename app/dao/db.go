package dao

import (
	"fmt"

	"yatter-backend-go/app/config"

	"github.com/jmoiron/sqlx"
)

func initDb(config config.MySQL) (*sqlx.DB, error) {
	db, err := sqlx.Open(config.DriverName(), config.DataSourceName())
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open failed: %w", err)
	}

	return db, nil
}
