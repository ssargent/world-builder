package service

//go:generate mockgen -source=interfaces.go -destination mocks/service.go  github.com/ssargent/world-builder/wb-api-go/internal/service EntityDataProvider, Cache, WriterDB, ReaderDB

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
)

type EntityDataProvider interface {
	repository.AttributeDefinitionQuerier
	repository.EntityAssociationQuerier
	repository.EntityAttributeQuerier
	repository.EntityQuerier
	repository.EntityHistoryQuerier
	repository.ReferenceQuerier
	repository.TypeQuerier
}

type Cache interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
}

type WriterDB interface {
	repository.DBTX
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

type ReaderDB interface {
	repository.DBTX
}
