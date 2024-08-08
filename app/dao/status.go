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

func (s *status) InsertStatus(ctx context.Context, content string, accountID int64) (int64, error) {
	query := `
        INSERT INTO status (account_id, content, create_at) 
        VALUES (?, ?, NOW())`

	// トランザクションを開始
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to start transaction: %w", err)
	}

	// ステータスの挿入
	result, err := tx.ExecContext(ctx, query, accountID, content)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, fmt.Errorf("failed to rolleback: %w", err)
		}
		return 0, fmt.Errorf("failed to insert status: %w", err)
	}

	// 挿入されたIDを取得
	statusID, err := result.LastInsertId()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, fmt.Errorf("failed to rolleback: %w", err)
		}
		return 0, fmt.Errorf("failed to retrieve last insert id: %w", err)
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return statusID, nil
}

func (s *status) DeleteStatus(ctx context.Context, id int64) error {
	query := `
		DELETE FROM status WHERE id = ?
	`
	// トランザクションを開始
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// ステータスの削除
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("delete status failed: %w, rollback failed: %v", err, rbErr)
		}
		return fmt.Errorf("failed to delete status: %w", err)
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
