package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/config"
	"github.com/ssargent/world-builder/wb-api-go/internal/service"
)

type Handler struct {
	cfg *config.Config
	rdb *sqlx.DB
	wdb *sqlx.DB

	cache  *cache.Cache
	entity *service.EntityService
}

func NewHandler(cfg *config.Config, ent *service.EntityService) *Handler {

	return &Handler{

		cfg: cfg,

		entity: ent,
	}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()

	route(r, http.MethodGet, "/entities/{id}", h.getEntity)

	return r
}

type HTTPHandler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func route(r chi.Router, method, pattern string, h HTTPHandler, mw ...func(http.Handler) http.Handler) {
	if len(mw) > 0 {
		r.With(mw...)
	}
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if err := h(ctx, w, r); err != nil {
			fmt.Println(err)
		}
	}

	r.MethodFunc(method, pattern, fn)
}
