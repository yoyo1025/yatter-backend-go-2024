package usecase

import (
	"context"
	"errors"
	"strconv"
	"yatter-backend-go/app/domain/auth"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, status string) (*CreateStatusDTO, error)
	Get(ctx context.Context, statusID string) (*GetStatusDTO, error)
	ListPublicStatuses(ctx context.Context, maxID, sinceID, limit string) (*ListPublicStatusesDTO, error)
}

type status struct {
	db          *sqlx.DB
	statusRepo  repository.Status
	accountRepo repository.Account
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status, accountRepo repository.Account) *status {
	return &status{
		db:          db,
		statusRepo:  statusRepo,
		accountRepo: accountRepo,
	}
}

type CreateStatusDTO struct {
	Account *object.Account
	Status  *object.Status
}

func (s *status) Create(ctx context.Context, status string) (*CreateStatusDTO, error) {
	acc := auth.AccountOf(ctx)
	if acc == nil {
		return nil, errors.New("authorized account is not found")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	st := object.NewStatus(status)
	st, err = s.statusRepo.Create(ctx, tx, acc, st)
	if err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Account: acc,
		Status:  st,
	}, nil
}

type GetStatusDTO struct {
	Account *object.Account
	Status  *object.Status
}

func (s *status) Get(ctx context.Context, statusID string) (*GetStatusDTO, error) {
	sid, err := strconv.Atoi(statusID)

	st, err := s.statusRepo.Get(ctx, sid)
	if err != nil {
		return nil, err
	}

	acc, err := s.accountRepo.FindByID(ctx, st.AccountID)
	if err != nil {
		return nil, err
	}

	return &GetStatusDTO{
		Account: acc,
		Status:  st,
	}, nil
}

type ListPublicStatusesDTO struct {
	Account  []*object.Account
	Statuses []*object.Status
}

func (s *status) ListPublicStatuses(ctx context.Context, maxID, sinceID, limit string) (*ListPublicStatusesDTO, error) {
	acc := auth.AccountOf(ctx)
	if acc == nil {
		return nil, errors.New("authorized account is not found")
	}

	mid, _ := strconv.Atoi(maxID)
	sid, _ := strconv.Atoi(sinceID)
	lim, _ := strconv.Atoi(limit)

	if lim == 0 {
		lim = 40
	}

	statuses, err := s.statusRepo.List(ctx, acc, mid, sid, lim)
	if err != nil {
		return nil, err
	}

	accounts := make([]*object.Account, len(statuses))
	for i := range accounts {
		accounts[i] = acc
	}

	return &ListPublicStatusesDTO{
		Account:  accounts,
		Statuses: statuses,
	}, nil
}
