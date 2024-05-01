package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type status struct {
	db *sqlx.DB
}

func NewStatus(db *sqlx.DB) *status {
	return &status{
		db: db,
	}
}

var _ repository.Status = (*status)(nil)

func (s *status) Create(ctx context.Context, tx *sqlx.Tx, account *object.Account, status *object.Status) (*object.Status, error) {
	res, err := s.db.Exec("insert into status (content, account_id, created_at) values (?, ?, ?)", status.Content, account.ID, status.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert status: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	status.ID = int(id)

	return status, nil
}

func (s *status) Get(ctx context.Context, statusID int) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", statusID).StructScan(entity)
	if err != nil {
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}

func (s *status) List(ctx context.Context, account *object.Account, maxID, sinceID, limit int) ([]*object.Status, error) {
	q := `select * from status where account_id = ? ORDER BY id DESC LIMIT ?`
	args := []interface{}{account.ID, limit}
	if sinceID != 0 && maxID != 0 {
		q = `select * from status where account_id = ? AND ? <= id AND id >= ?
		ORDER BY id DESC LIMIT ?
		`
		args = []interface{}{account.ID, sinceID, maxID, limit}
	}

	if sinceID != 0 {
		q = `select * from status where account_id = ? 
		AND id >= ?
		ORDER BY id DESC LIMIT ?
		`
		args = []interface{}{account.ID, sinceID, limit}
	}

	if maxID != 0 {
		q = `select * from status where account_id = ?
		AND id <= ?
		ORDER BY id DESC LIMIT ?
		`
		args = []interface{}{account.ID, maxID, limit}
	}

	rows, err := s.db.QueryxContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list status from db: %w", err)
	}
	defer rows.Close()

	var statuses []*object.Status
	for rows.Next() {
		entity := new(object.Status)
		if err := rows.StructScan(entity); err != nil {
			return nil, fmt.Errorf("failed to scan status: %w", err)
		}

		statuses = append(statuses, entity)
	}

	return statuses, nil
}
