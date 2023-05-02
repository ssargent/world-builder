package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

type DB interface {
	GetContext(ctx context.Context, dest interface{}, query string, arg ...interface{}) error
	QueryxContext(ctx context.Context, query string, arg ...interface{}) (*sqlx.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type EntityRepository interface {
	FindByID(db DB, id uuid.UUID) (*entities.Entity, error)
	FindByWBRN(db DB, wbrn string) (*entities.Entity, error)
	FindByTypeName(db DB, typeName string) (*entities.Entity, error)
}
