package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-gorp/gorp/v3"
)

type (
	account struct {
		sql gorp.SqlExecutor
	}
)

func NewAccount(sql gorp.SqlExecutor) repository.Account {
	return &account{sql: sql}
}

func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.sql.SelectOne(entity, "select * from account where username = ?", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, fmt.Errorf("%w", err)
		}
	}
	return entity, nil
}
