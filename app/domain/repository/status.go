package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, tx *sqlx.Tx, account *object.Account, status *object.Status) (*object.Status, error)
	Get(ctx context.Context, statusID int) (*object.Status, error)
	List(ctx context.Context, account *object.Account, maxID, sinceID, limit int) ([]*object.Status, error)
}
