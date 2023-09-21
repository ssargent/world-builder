package repository

//go:generate mockgen -source=interface.go  -aux_files=github.com/ssargent/world-builder/wb-api-go/internal/repository=db.go -destination mocks/repository.go  github.com/ssargent/world-builder/wb-api-go/internal/repository  WriterDB, ReaderDB, Transaction, Manager

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Manager interface {
	Transaction(ctx context.Context, db WriterDB, opts *sql.TxOptions) (Transaction, error)
}
type WriterDB interface {
	DBTX
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

type ReaderDB interface {
	DBTX
}

type Transaction interface {
	DBTX
	Commit() error
	Rollback() error
}

type AttributeDefinitionQuerier interface {
	CreateAttributeDefinition(
		ctx context.Context,
		db DBTX,
		arg *CreateAttributeDefinitionParams) (*WorldAttributeDefinition, error)
	GetAttributesForType(ctx context.Context, db DBTX, typeID uuid.UUID) ([]*WorldAttributeDefinition, error)
}

type EntityQuerier interface {
	CreateEntity(ctx context.Context, db DBTX, arg *CreateEntityParams) (*WorldEntity, error)
	GetEntitiesByParent(ctx context.Context, db DBTX, parentID uuid.UUID) ([]*WorldEntity, error)
	GetEntitiesByWBRN(ctx context.Context, db DBTX, wbrn string) ([]*WorldEntity, error)
	GetEntity(ctx context.Context, db DBTX, id uuid.UUID) (*WorldEntity, error)
	GetEntityByWBRN(ctx context.Context, db DBTX, wbrn string) (*WorldEntity, error)
	GetEntitiesByCriteria(ctx context.Context, db DBTX, arg *GetEntitiesByCriteriaParams) ([]*WorldEntity, error)
}

type EntityAssociationQuerier interface {
	CreateEntityAssociation(
		ctx context.Context,
		db DBTX,
		arg *CreateEntityAssociationParams) (*WorldEntityAssociation, error)
	GetEntityAssociationsForEntity(ctx context.Context, db DBTX, entityOne uuid.UUID) ([]*WorldEntityAssociation, error)
}

type EntityHistoryQuerier interface {
	CreateEntityHistory(ctx context.Context, db DBTX, arg *CreateEntityHistoryParams) (*WorldEntityHistory, error)
	GetEntityHistory(ctx context.Context, db DBTX, entityID uuid.UUID) ([]*WorldEntityHistory, error)
}

type TypeQuerier interface {
	CreateType(ctx context.Context, db DBTX, arg *CreateTypeParams) (*WorldType, error)
	CreateTypeAttribute(ctx context.Context, db DBTX, arg *CreateTypeAttributeParams) (*WorldTypeAttribute, error)
	GetTypeByID(ctx context.Context, db DBTX, id uuid.UUID) (*WorldType, error)
	GetTypeByWBTN(ctx context.Context, db DBTX, wbtn string) (*WorldType, error)
	GetFullTypeAttributes(ctx context.Context, db DBTX, id uuid.UUID) ([]*GetFullTypeAttributesRow, error)
}

type EntityAttributeQuerier interface {
	GetEntityAttributes(ctx context.Context, db DBTX, entityID uuid.UUID) ([]*WorldEntityAttribute, error)
	GetAttributeByWBATN(ctx context.Context, db DBTX, wbatn string) (*WorldAttributeDefinition, error)
	CreateEntityAttribute(ctx context.Context, db DBTX, arg *CreateEntityAttributeParams) (*WorldEntityAttribute, error)
}

type ReferenceQuerier interface {
	GetEntityChildReferences(ctx context.Context, db DBTX, parentID uuid.UUID) ([]*GetEntityChildReferencesRow, error)
	GetEntityReference(ctx context.Context, db DBTX, id uuid.UUID) (*GetEntityReferenceRow, error)
	GetEntityReferenceByWBRN(ctx context.Context, db DBTX, wbrn string) (*GetEntityReferenceByWBRNRow, error)
}
