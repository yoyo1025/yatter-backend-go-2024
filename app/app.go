package app

import (
	"yatter-backend-go/app/config"
	"yatter-backend-go/app/dao"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

// Dependency manager for whole application
type App struct {
	*sqlx.DB

	AccountRepository repository.Account
}

// Create dependency manager
func NewApp() (*App, error) {
	// panic if lacking something

	db, err := dao.NewDB(config.MySQLConfig())
	if err != nil {
		return nil, err
	}

	return &App{
		DB:                db,
		AccountRepository: dao.NewAccount(db),
	}, nil
}
