package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Statuses interface {
	FetchStatus(ctx context.Context, id int64) (*GetStatusDTO, error)
	CreateStatus(ctx context.Context, content string, accountID int64) (*CreateStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type GetStatusDTO struct {
	Status *object.StatusDetail
}

type CreateStatusDTO struct {
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

func (s *status) CreateStatus(ctx context.Context, content string, accountID int64) (*CreateStatusDTO, error) {
	// ステータスの挿入
	statusID, err := s.statusRepo.InsertStatus(ctx, content, accountID)
	if err != nil {
		return nil, err
	}

	// 挿入されたステータスの詳細を取得
	statusInfo, err := s.statusRepo.GetStatusByID(ctx, statusID)
	if err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Status: statusInfo,
	}, nil
}
