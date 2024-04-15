package statuses

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Status string
	// TODO: Medias Field
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := new(object.Status)
	status.Status = req.Status
	// account_info := auth.AccountOf(r) 認証情報を取得する

	panic("Must Implement Account Registration")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
