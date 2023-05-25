package internal

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/config"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/handlers"
)

type API struct {
	cfg    *config.Config
	Reader *sqlx.DB
	Writer *sqlx.DB
}

func NewApi(cfg *config.Config, rdb *sqlx.DB, wdb *sqlx.DB) *API {
	return &API{
		cfg:    cfg,
		Reader: rdb,
		Writer: wdb,
	}
}

func (a *API) ListenAndServe() {
	r := chi.NewRouter()

	h := handlers.NewHandler(a.cfg, a.Reader, a.Writer)

	r.Mount("/", h.Routes())

	http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.Port), r)
}
