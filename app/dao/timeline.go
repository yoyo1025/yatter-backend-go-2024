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
	timeline struct {
		db *sqlx.DB
	}
)

// インターフェースの実装チェック
var _ repository.Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (t *timeline) GetTimelines(ctx context.Context) ([]object.Timeline, error) {
	entity := []object.Timeline{} // スライスを直接初期化
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
	`
	err := t.db.SelectContext(ctx, &entity, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	return entity, nil // スライス自体を返す
}
