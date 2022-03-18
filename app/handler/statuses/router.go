package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	sr repository.Status
}

func NewRouter(am func(http.Handler) http.Handler, sr repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr}
	r.Get("/{id}", h.Find)
	r.With(am).Post("/", h.Create)

	return r
}
