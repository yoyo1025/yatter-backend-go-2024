package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	GetStatusByID(ctx context.Context, id int64) (*object.StatusDetail, error)
}
