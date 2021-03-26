package dao

import (
	"database/sql"
	"fmt"

	"yatter-backend-go/app/domain/object"

	"github.com/go-gorp/gorp/v3"
)

type DBConfig interface {
	FormatDSN() string
}

func initDb(config DBConfig) (*gorp.DbMap, error) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("sql.Open failed: %w", err)
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	// add a table, setting the table name and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(object.Account{}, "account").SetKeys(true, "ID")

	return dbmap, nil
}
