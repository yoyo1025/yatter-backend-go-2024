package timelines

import (
	"encoding/json"
	"net/http"
)

func (h *handler) FetchTimelines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto, err := h.timelineUsecase.FetchTimelines(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
