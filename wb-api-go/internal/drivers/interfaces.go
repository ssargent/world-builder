package drivers

import (
	"context"

	"github.com/jmoiron/sqlx"
)

/*
	Drivers are interfaces for external resources,  database, cache etc..
*/

type DB interface {
	GetContext(ctx context.Context, dest interface{}, query string, arg ...interface{}) error
	QueryxContext(ctx context.Context, query string, arg ...interface{}) (*sqlx.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
