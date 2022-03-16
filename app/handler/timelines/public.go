package timelines

import (
	"encoding/json"
	"net/http"
)

func (h *handler) Public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cond, err := parseCondition(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.app.Dao.Status().FindMany(ctx, cond)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
