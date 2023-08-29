package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

type TypeService struct {
	cache   Cache
	reader  repository.ReaderDB
	writer  repository.WriterDB
	manager repository.Manager
	queries EntityDataProvider
}

func NewTypeService(c *cache.Cache, rdb, wdb *sqlx.DB, queries EntityDataProvider) *TypeService {
	return &TypeService{
		cache:   c,
		reader:  rdb,
		writer:  wdb,
		manager: repository.NewManager(),
		queries: queries,
	}
}

// CreateType will create a new type in the database.
func (t *TypeService) CreateType(ctx context.Context, in *entities.EntityType) (*entities.EntityType, error) {
	createTypeParam := repository.CreateTypeParams{
		ParentID:        in.Parent.TypeID,
		Wbtn:            in.Wbtn,
		TypeName:        in.TypeName,
		TypeDescription: in.TypeDescription,
	}

	//Manager just wraps BeginTxx and returns an interface, so we can easily test.
	m := repository.NewManager()

	txn, err := m.Transaction(ctx, t.writer, &sql.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("BeginTxx: %w", err)
	}
	defer txn.Rollback()

	created, err := t.queries.CreateType(ctx, txn, &createTypeParam)
	if err != nil {
		return nil, fmt.Errorf("CreateType: %w", err)
	}

	if len(in.Attributes) > 0 {
		for i, attr := range in.Attributes {
			// we might not have the full attribute here.  we might just have the wbatn.
			attribute, err := t.getAttribute(ctx, txn, attr.Wbatn)
			if err != nil {
				return nil, fmt.Errorf("GetAttributeByTypeName: %w", err)
			}

			typeAttribParams := repository.CreateTypeAttributeParams{
				TypeID:      created.ID,
				AttributeID: attribute.ID,
				Ordinal:     int32(i),
				IsRequired:  true,
			}

			_, err = t.queries.CreateTypeAttribute(ctx, txn, &typeAttribParams)
			if err != nil {
				return nil, fmt.Errorf("CreateTypeAttribute: %w", err)
			}
		}
	}

	fullType, err := t.getType(ctx, txn, created.ID, true)
	if err != nil {
		return nil, fmt.Errorf("getType: %w", err)
	}

	if err := txn.Commit(); err != nil {
		return nil, fmt.Errorf("commit: %w", err)
	}

	return fullType, nil
}

// most likely entity service will take a reference to type service and leverage this code.
// so code in other file will go away.
//
//nolint:dupl // ok for duplicate while I figure out how this should work.
func (t *TypeService) getType(
	ctx context.Context,
	db repository.DBTX,
	id uuid.UUID,
	attributeDefinitions bool) (*entities.EntityType, error) {
	wt, err := t.queries.GetTypeByID(ctx, db, id)
	if err != nil {
		return nil, fmt.Errorf("GetTypeByID: %w", err)
	}

	pt, err := t.queries.GetTypeByID(ctx, db, wt.ParentID)
	if err != nil {
		return nil, fmt.Errorf("GetTypeByID: %w", err)
	}

	entityType := entities.EntityType{
		ID:              wt.ID,
		Wbtn:            wt.Wbtn,
		TypeName:        wt.TypeName,
		TypeDescription: wt.TypeDescription,
		CreatedAt:       wt.CreatedAt.Time,
		UpdatedAt:       wt.UpdatedAt.Time,

		Parent: &entities.TypeReference{
			TypeID:   pt.ID,
			TypeName: pt.Wbtn,
		},
	}

	if attributeDefinitions {
		atts, err := t.queries.GetAttributesForType(ctx, db, id)
		if err != nil {
			return nil, fmt.Errorf("GetAttributesForType: %w", err)
		}

		entityType.Attributes = make([]*entities.Attribute, len(atts))

		for i, at := range atts {
			entityType.Attributes[i] = &entities.Attribute{
				ID:            at.ID,
				Wbatn:         at.Wbatn,
				AttributeName: at.AttributeName,
				DataType:      at.DataType,
				Label:         at.Label,
				CreatedAt:     at.CreatedAt,
				UpdatedAt:     at.UpdatedAt,
			}
		}
	}

	return &entityType, nil
}

func (t *TypeService) getAttribute(ctx context.Context, db repository.DBTX, wbatn string) (*entities.Attribute, error) {
	attr, found := t.cache.Get(wbatn)
	if found {
		a, ok := attr.(*entities.Attribute)
		if ok {
			return a, nil
		}
	}

	attribute, err := t.queries.GetAttributeByWBATN(ctx, db, wbatn)
	if err != nil {
		return nil, fmt.Errorf("GetAttributeByWBATN: %w", err)
	}

	t.cache.Set(wbatn, attribute, cache.DefaultExpiration)

	return &entities.Attribute{
		ID:            attribute.ID,
		Wbatn:         attribute.Wbatn,
		AttributeName: attribute.AttributeName,
		DataType:      attribute.DataType,
		Label:         attribute.Label,
		CreatedAt:     attribute.CreatedAt,
		UpdatedAt:     attribute.UpdatedAt,
	}, nil
}
