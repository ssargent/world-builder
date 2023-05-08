package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DB interface {
	GetContext(ctx context.Context, dest interface{}, query string, arg ...interface{}) error
	QueryxContext(ctx context.Context, query string, arg ...interface{}) (*sqlx.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
