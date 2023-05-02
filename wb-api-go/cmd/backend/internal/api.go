package internal

import (
	"github.com/jmoiron/sqlx"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/config"
)

type API struct {
	cfg *config.Config
	DB  *sqlx.DB
}

func NewApi(cfg *config.Config, db *sqlx.DB) *API {
	return &API{
		cfg: cfg,
		DB:  db,
	}
}

func (a *API) ListenAndServe() {

}
