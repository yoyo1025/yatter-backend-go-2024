package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	dto, err := h.accountUsecase.Get(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if json.NewEncoder(w).Encode(dto.Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
