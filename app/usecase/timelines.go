package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Timelines interface {
	FetchTimelines(ctx context.Context) (*[]GetTimelinesDTO, error)
}

type timeline struct {
	db           *sqlx.DB
	timelineRepo repository.Timeline
}

func NewTimeline(db *sqlx.DB, timelineRepo repository.Timeline) *timeline {
	return &timeline{
		db:           db,
		timelineRepo: timelineRepo,
	}
}

type GetTimelinesDTO struct {
	Timeline *object.Timeline
}

func (t *timeline) FetchTimelines(ctx context.Context) (*[]GetTimelinesDTO, error) {
	// `timelineRepo.GetTimelines` を呼び出してデータを取得
	timelines, err := t.timelineRepo.GetTimelines(ctx)
	if err != nil {
		return nil, err
	}

	// `GetTimelinesDTO` 型に変換して返す
	var result []GetTimelinesDTO
	for _, timeline := range timelines {
		dto := GetTimelinesDTO{
			Timeline: &timeline,
		}
		result = append(result, dto)
	}

	return &result, nil
}
