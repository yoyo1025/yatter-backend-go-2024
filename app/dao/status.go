package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

// インターフェースの実装チェック
var _ repository.Status = (*status)(nil)

func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

func (s *status) GetStatusByID(ctx context.Context, id int64) (*object.StatusDetail, error) {
	entity := new(object.StatusDetail)
	query := `
    SELECT 
        s.id,
        a.id AS "account.id",
        a.username AS "account.username",
        a.create_at AS "account.create_at",
        a.note AS "account.note",
        a.avatar AS "account.avatar",
        a.header AS "account.header",
        s.content
    FROM status AS s
    LEFT JOIN account AS a ON s.account_id = a.id
    WHERE s.id = ?`
	err := s.db.QueryRowxContext(ctx, query, id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}
