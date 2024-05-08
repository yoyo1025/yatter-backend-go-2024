package statuses

import (
	"net/http"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Status string
	// TODO: Medias Field
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	// account_info := auth.AccountOf(r) 認証情報を取得する

	panic("Must Implement Status Creation")

}
