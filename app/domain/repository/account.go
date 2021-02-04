package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	Find(ctx context.Context, id object.AccountID) (*object.Account, error)
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	Create(ctx context.Context, account *object.Account) error
	Update(ctx context.Context, account *object.Account) error
	//Delete(ctx context.Context, id object.AccountID) error
}
