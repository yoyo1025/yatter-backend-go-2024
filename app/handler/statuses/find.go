package statuses

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"
	"yatter-backend-go/app/handler/request"
)

func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := request.IDOf(r)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	resp, err := h.app.Dao.Status().Find(ctx, id)
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
