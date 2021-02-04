package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-gorp/gorp/v3"
)

type (
	status struct {
		sql gorp.SqlExecutor
	}
)

func NewStatus(sql gorp.SqlExecutor) repository.Status {
	return &status{sql: sql}
}

func (r *status) Find(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)
	exists, err := r.sql.Get(entity, id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	} else if exists == nil {
		return nil, nil
	}
	return entity, nil
}

func (r *status) Create(ctx context.Context, entity *object.Status) error {
	if err := r.sql.Insert(entity); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (r *status) Update(ctx context.Context, entity *object.Status) error {
	if _, err := r.sql.Update(entity); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (r *status) FindMany(ctx context.Context, cond *object.FindStatusCondition) ([]*object.Status, error) {
	query := sq.Select("*").From("status")
	if cond.Limit > 0 && cond.Limit <= 80 {
		query = query.Limit(uint64(cond.Limit))
	} else {
		query = query.Limit(40)
	}

	if cond.SinceID > 0 {
		query = query.Where(sq.Gt{"id": cond.SinceID})
	}

	if cond.MaxID > 0 {
		query = query.Where(sq.Lt{"id": cond.MaxID})
	}

	query = query.OrderBy("id DESC")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var result []*object.Status
	if _, err := r.sql.Select(&result, sql, args...); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return result, nil
}
