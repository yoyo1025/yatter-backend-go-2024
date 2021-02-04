package dao

import (
	"fmt"
	"log"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-gorp/gorp/v3"
)

type (
	Dao interface {
		Account() repository.Account
		Status() repository.Status
		InitAll() error
	}

	dao struct {
		dbmap *gorp.DbMap
	}
)

func New(config DBConfig) (Dao, error) {
	dbmap, err := initDb(config)
	if err != nil {
		return nil, err
	}

	return &dao{dbmap: dbmap}, nil
}

func (d *dao) Account() repository.Account {
	return NewAccount(d.dbmap)
}

func (d *dao) Status() repository.Status {
	return NewStatus(d.dbmap)
}

func (d *dao) InitAll() error {
	if err := d.exec("SET FOREIGN_KEY_CHECKS=0"); err != nil {
		return fmt.Errorf("Can't disable FOREIGN_KEY_CHECKS: %w", err)
	}

	defer func() {
		err := d.exec("SET FOREIGN_KEY_CHECKS=0")
		if err != nil {
			log.Printf("Can't restore FOREIGN_KEY_CHECKS: %+v", err)
		}
	}()

	for _, table := range []string{"account", "status"} {
		if err := d.exec("TRUNCATE TABLE " + table); err != nil {
			return fmt.Errorf("Can't truncate table "+table+": %w", err)
		}
	}

	return nil
}

func (d *dao) exec(query string, args ...interface{}) error {
	_, err := d.dbmap.Exec(query, args...)
	return err
}
