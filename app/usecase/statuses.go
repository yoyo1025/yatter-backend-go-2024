package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Statuses interface {
	FetchStatus(ctx context.Context, id int64) (*GetStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type GetStatusDTO struct {
	Status *object.StatusDetail
}

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) FetchStatus(ctx context.Context, id int64) (*GetStatusDTO, error) {
	statusInfo, err := s.statusRepo.GetStatusByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetStatusDTO{
		Status: statusInfo,
	}, nil
}
