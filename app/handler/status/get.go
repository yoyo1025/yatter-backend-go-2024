package status

import (
	"encoding/json"
	"net/http"
	"time"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

func (h *handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	statusID := chi.URLParam(r, "status_id")
	if statusID == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	dto, err := h.StatusUsecase.Get(ctx, statusID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newGetStatusHTTPResponse(dto)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type GetStatusHTTPResponse struct {
	ID      int `json:"id"`
	Account struct {
		ID        int       `json:"id"`
		UserName  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"account"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func newGetStatusHTTPResponse(dto *usecase.GetStatusDTO) *GetStatusHTTPResponse {
	return &GetStatusHTTPResponse{
		ID: dto.Status.ID,
		Account: struct {
			ID        int       `json:"id"`
			UserName  string    `json:"username"`
			CreatedAt time.Time `json:"created_at"`
		}{
			ID:        int(dto.Account.ID),
			UserName:  dto.Account.Username,
			CreatedAt: dto.Account.CreateAt,
		},
		Content:   dto.Status.Content,
		CreatedAt: dto.Status.CreatedAt,
	}
}

type GetTimelineStatusesHTTPResponse struct {
	Statuses []*GetStatusHTTPResponse `json:"statuses"`
}

func (h *handler) GetTimelineStatuses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sinceID := r.URL.Query().Get("since_id")
	maxID := r.URL.Query().Get("max_id")
	limit := r.URL.Query().Get("limit")

	dto, err := h.StatusUsecase.ListPublicStatuses(ctx, maxID, sinceID, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newGetTimelineStatusesHTTPResponse(dto)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func newGetTimelineStatusesHTTPResponse(dto *usecase.ListPublicStatusesDTO) *GetTimelineStatusesHTTPResponse {
	statuses := make([]*GetStatusHTTPResponse, len(dto.Statuses))
	for i := range dto.Statuses {
		statuses[i] = &GetStatusHTTPResponse{
			ID: dto.Statuses[i].ID,
			Account: struct {
				ID        int       `json:"id"`
				UserName  string    `json:"username"`
				CreatedAt time.Time `json:"created_at"`
			}{
				ID:        int(dto.Account[i].ID),
				UserName:  dto.Account[i].Username,
				CreatedAt: dto.Account[i].CreateAt,
			},
			Content:   dto.Statuses[i].Content,
			CreatedAt: dto.Statuses[i].CreatedAt,
		}
	}

	return &GetTimelineStatusesHTTPResponse{
		Statuses: statuses,
	}
}
