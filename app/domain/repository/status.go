package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	Find(ctx context.Context, id object.StatusID) (*object.Status, error)
	FindMany(ctx context.Context, condition *object.FindStatusCondition) ([]*object.Status, error)
	Create(ctx context.Context, account *object.Status) error
	Delete(ctx context.Context, id object.StatusID) error
}
