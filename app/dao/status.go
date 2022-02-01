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
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// Find: IDでステータスを取得
func (r *status) Find(ctx context.Context, id object.StatusID) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// FindMany: IDでステータスを取得
func (r *status) FindMany(ctx context.Context, condition *object.FindStatusCondition) ([]*object.Status, error) {
	query := `select * from status`
	if condition.SinceID > 0 && condition.MaxID > 0 {
		query = fmt.Sprintf(`%s where id between %d AND %d`, query, condition.SinceID, condition.MaxID)
	} else if condition.SinceID > 0 {
		query = fmt.Sprintf(`%s where %d <= id`, query, condition.SinceID)
	} else if condition.MaxID > 0 {
		query = fmt.Sprintf(`%s where id <= %d`, query, condition.MaxID)
	}
	if condition.Limit > 0 {
		query = fmt.Sprintf(`%s limit %d`, query, condition.Limit)
	}

	entities := make([]*object.Status, 0)
	if err := r.db.SelectContext(ctx, &entities, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	return entities, nil
}

// Create : ステータスを作成
func (r *status) Create(ctx context.Context, status *object.Status) error {
	if _, err := r.db.ExecContext(ctx, `
	insert into status (account_id, content)
	values (?, ?)
	`, status.AccountID, status.Content); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Delete : ステータスを削除
func (r *status) Delete(ctx context.Context, id object.StatusID) error {
	if _, err := r.db.ExecContext(ctx, `delete from status where id = ?`, id); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
