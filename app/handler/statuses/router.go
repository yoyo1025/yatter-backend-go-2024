package statuses

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	statusUsecase usecase.Statuses
}

// Create Handler for `/v1/statuses/`
func NewRouter(u usecase.Statuses) http.Handler {
	r := chi.NewRouter()

	// r.Group()により、特定のグループに対してミドルウェアを適用する
	// グループに対して適用されたミドルウェアは、そのグループに属する全てのエンドポイントに対して適用される
	// r.Group(func(r chi.Router) {
	// 	// リクエストの認証を行う
	// 	r.Use(auth.Middleware(ar))
	// 	h := &handler{ar}
	// 	r.Post("/", h.Create)
	// })
	h := &handler{
		statusUsecase: u,
	}
	r.Get("/{id}", h.FetchStatuse)

	return r
}
