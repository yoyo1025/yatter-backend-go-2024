package usecase

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Account interface {
	Create(ctx context.Context, username, password string) (*CreateAccountDTO, error)
	Fetch(ctx context.Context, username string) (*GetAccountDTO, error)
}

type account struct {
	db          *sqlx.DB
	accountRepo repository.Account
}

type CreateAccountDTO struct {
	Account *object.Account
}

type GetAccountDTO struct {
	Account *object.Account
}

var _ Account = (*account)(nil)

func NewAcocunt(db *sqlx.DB, accountRepo repository.Account) *account {
	return &account{
		db:          db,
		accountRepo: accountRepo,
	}
}

func (a *account) Create(ctx context.Context, username, password string) (*CreateAccountDTO, error) {
	acc, err := object.NewAccount(username, password)
	if err != nil {
		return nil, err
	}

	tx, err := a.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		// パニックが発生した場合、トランザクションをロールバック
		if r := recover(); r != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				// ロールバックエラーが発生した場合でも、パニックを再発行する
				panic(fmt.Sprintf("panic: %v; rollback failed: %v", r, rollbackErr))
			}
			panic(r)
		}

		if commitErr := tx.Commit(); commitErr != nil {
			panic(fmt.Sprintf("commit failed: %v", commitErr))
		}
	}()

	// accountRepoを使用していることから、usecaseはRepoに依存していることがわかる
	if err := a.accountRepo.Create(ctx, tx, acc); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return nil, fmt.Errorf("create account failed: %v, rollback failed: %v", err, rollbackErr)
		}
		return nil, err
	}

	return &CreateAccountDTO{
		Account: acc,
	}, nil
}

func (a *account) Fetch(ctx context.Context, username string) (*GetAccountDTO, error) {
	accountInfo, err := a.accountRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &GetAccountDTO{
		Account: accountInfo,
	}, nil
}
