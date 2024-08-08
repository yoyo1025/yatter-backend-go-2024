package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	GetStatusByID(ctx context.Context, id int64) (*object.StatusDetail, error)
	InsertStatus(ctx context.Context, content string, accountID int64) (int64, error)
	DeleteStatus(ctx context.Context, id int64) error
}
