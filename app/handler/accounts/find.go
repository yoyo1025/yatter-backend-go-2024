package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "username was empty", http.StatusBadRequest)
		return
	}

	resp, err := h.app.AccountRepository.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
