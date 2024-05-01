package status

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

func NewRouter(su usecase.Status, ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := NewHandler(su)

	m := auth.Middleware(ar)
	r.Use(m)
	r.Post("/v1/statuses/", h.PostStatus)
	r.Get("/v1/statuses/{status_id}", h.GetStatus)
	r.Get("/v1/timelines/public", h.GetTimelineStatuses)

	return r
}
