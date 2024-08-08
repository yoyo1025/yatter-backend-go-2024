package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/auth"
)

// Request body for `POST /v1/statuses`
type AddStatusRequest struct {
	Status string
}

// Handle request for `POST /v1/statuses`
func (h *handler) CreateStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account_info := auth.AccountOf(r.Context()) // 認証情報を取得する
	if account_info == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var req AddStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dto, err := h.statusUsecase.CreateStatus(ctx, req.Status, account_info.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
