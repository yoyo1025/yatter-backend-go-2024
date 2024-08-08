package statuses

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) DeleteStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	if err := h.statusUsecase.DeleteStatus(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
}
