package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
)

type TypeService struct {
	cache   Cache
	reader  *sqlx.DB
	writer  *sqlx.DB
	queries EntityDataProvider
}

func NewTypeService(c *cache.Cache, rdb, wdb *sqlx.DB, queries EntityDataProvider) *TypeService {
	return &TypeService{
		cache:   c,
		reader:  rdb,
		writer:  wdb,
		queries: queries,
	}
}
