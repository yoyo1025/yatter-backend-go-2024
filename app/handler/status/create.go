package status

import (
	"encoding/json"
	"net/http"
	"time"
	"yatter-backend-go/app/usecase"
)

type AddStatusRequest struct {
	Status string `json:"status"`
}

type handler struct {
	StatusUsecase usecase.Status
}

func NewHandler(u usecase.Status) *handler {
	return &handler{
		StatusUsecase: u,
	}
}

func (h *handler) PostStatus(w http.ResponseWriter, r *http.Request) {
	var req AddStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	dto, err := h.StatusUsecase.Create(ctx, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newCreateStatusHTTPResposnse(dto)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type CreateStatusHTTPResponse struct {
	ID      int `json:"id"`
	Account struct {
		ID        int       `json:"id"`
		UserName  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"account"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func newCreateStatusHTTPResposnse(dto *usecase.CreateStatusDTO) *CreateStatusHTTPResponse {
	return &CreateStatusHTTPResponse{
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
