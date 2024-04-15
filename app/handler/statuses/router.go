package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Account) http.Handler {
	r := chi.NewRouter()

	// リクエストの認証を行う
	r.Use(auth.Middleware(ar))
	h := &handler{ar}
	r.Post("/", h.Create)

	return r
}
