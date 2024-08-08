package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	GetTimelines(ctx context.Context) ([]object.Timeline, error)
}
