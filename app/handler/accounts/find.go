package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "username was empty", http.StatusBadRequest)
		return
	}

	resp, err := h.app.Dao.Account().FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if resp == nil {
		httperror.Error(w, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
