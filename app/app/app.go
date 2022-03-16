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
	StatusRepository  repository.Status
}

// Create dependency manager
func NewApp() (*App, error) {
	// panic if lacking something
	app := new(App)
	daoCfg := config.MySQLConfig()

	var err error
	if app.DB, err = dao.NewDB(daoCfg); err != nil {
		return nil, err
	}
	app.AccountRepository = dao.NewAccount(app.DB)
	app.StatusRepository = dao.NewStatus(app.DB)

	return app, nil
}
