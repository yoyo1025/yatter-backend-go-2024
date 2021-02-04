package timelines

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) Public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cond, err := parseCondition(r.URL.Query())
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	resp, err := h.app.Dao.Status().FindMany(ctx, cond)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
