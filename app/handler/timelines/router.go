package timelines

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	timelineUsecase usecase.Timelines
}

func NewRouter(u usecase.Timelines) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		timelineUsecase: u,
	}

	r.Get("/public", h.FetchTimelines)

	return r
}
