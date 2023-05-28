package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

type EntityService struct {
	c       Cache
	rdb     repository.DBTX
	wdb     repository.DBTX
	queries EntityDataProvider
}

func NewEntityService(c *cache.Cache, rdb, wdb *sqlx.DB, queries EntityDataProvider) *EntityService {
	return &EntityService{
		c:       c,
		rdb:     rdb,
		wdb:     wdb,
		queries: queries,
	}
}

func (e *EntityService) FindByID(ctx context.Context, id uuid.UUID) (*entities.Entity, error) {
	cached, found := e.c.Get(fmt.Sprintf("entity:%s", id))
	if found {
		return cached.(*entities.Entity), nil
	}

	ent, err := e.get(ctx, e.rdb, id, true, true, true)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	e.c.Set(fmt.Sprintf("entity:%s", id), ent, cache.DefaultExpiration)
	return ent, nil
}

func (e *EntityService) FindByWBRN(ctx context.Context, wbrn string) (*entities.Entity, error) {
	ref, err := e.queries.GetEntityReferenceByWBRN(ctx, e.rdb, wbrn)
	if err != nil {
		return nil, fmt.Errorf("GetEntityReferenceByWBRN: %w", err)
	}

	cached, found := e.c.Get(fmt.Sprintf("entity:%s", ref.EntityID))
	if found {
		return cached.(*entities.Entity), nil
	}

	ent, err := e.get(ctx, e.rdb, ref.EntityID, true, true, true)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	e.c.Set(fmt.Sprintf("entity:%s", ref.EntityID), ent, cache.DefaultExpiration)
	return ent, nil

}

func (e *EntityService) get(ctx context.Context, db repository.DBTX, id uuid.UUID, associations, attributes, children bool) (*entities.Entity, error) {
	entity, err := e.queries.GetEntity(ctx, db, id)
	if err != nil {
		return nil, fmt.Errorf("GetEntity: %w", err)
	}

	fullEntity, err := mapEntity(entity)
	if err != nil {
		return nil, fmt.Errorf("mapEntity: %w", err)
	}

	parentEntity, err := e.queries.GetEntityReference(ctx, db, entity.ParentID)
	if err != nil {
		return nil, fmt.Errorf("GetEntityReference: %w", err)
	}

	fullEntity.Parent = &entities.EntityReference{
		EntityID:     parentEntity.EntityID,
		EntityName:   parentEntity.EntityName,
		ResourceName: parentEntity.ResourceName,
		TypeName:     parentEntity.TypeName,
	}

	typeInfo, err := e.getType(ctx, db, entity.TypeID, false)
	if err != nil {
		return nil, fmt.Errorf("getType: %w", err)
	}

	fullEntity.Type = &entities.TypeReference{
		TypeID:   typeInfo.ID,
		TypeName: typeInfo.Wbtn,
	}

	if attributes {
		if err := e.populateAttributes(ctx, db, fullEntity); err != nil {
			return nil, fmt.Errorf("populateAttributes: %w", err)
		}
	}

	if children {
		entityChildRefs, err := e.queries.GetEntityChildReferences(ctx, db, entity.ID)
		if err != nil {
			return nil, fmt.Errorf("GetEntityChildReferences: %w", err)
		}

		fullEntity.Children = make([]*entities.EntityReference, len(entityChildRefs))

		for i, entityRef := range entityChildRefs {
			fullEntity.Children[i] = &entities.EntityReference{
				EntityID:     entityRef.EntityID,
				EntityName:   entityRef.EntityName,
				ResourceName: entityRef.ResourceName,
				TypeName:     entityRef.TypeName,
			}
		}
	}

	return fullEntity, nil
}

func (e *EntityService) populateAttributes(ctx context.Context, db repository.DBTX, fullEntity *entities.Entity) error {
	entityAttributes, err := e.queries.GetEntityAttributes(ctx, db, fullEntity.ID)
	if err != nil {
		return fmt.Errorf("GetEntityAttributes: %w", err)
	}

	attribs, err := e.getTypeAttributes(ctx, db, fullEntity.Type.TypeID)
	if err != nil {
		return fmt.Errorf("getTypeAttributes: %w", err)
	}

	fmt.Printf("entityAttributes %+v\n", entityAttributes)
	fullEntity.Attributes = make([]*entities.EntityAttribute, len(entityAttributes))

	for i, entAtt := range entityAttributes {
		attributeDef, ok := attribs[entAtt.AttributeID]
		if !ok {
			return fmt.Errorf("attribute definition not found: %s", entAtt.AttributeID)
		}

		fullEntity.Attributes[i] = &entities.EntityAttribute{
			Name:  attributeDef.AttributeName,
			Type:  attributeDef.DataType,
			Value: entAtt.AttributeValue,
		}
	}
	return nil
}

func (e *EntityService) getType(ctx context.Context, db repository.DBTX, id uuid.UUID, attributeDefinitions bool) (*entities.EntityType, error) {
	wt, err := e.queries.GetTypeByID(ctx, db, id)
	if err != nil {
		return nil, fmt.Errorf("GetTypeByID: %w", err)
	}

	pt, err := e.queries.GetTypeByID(ctx, db, wt.ParentID)
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
		atts, err := e.queries.GetAttributesForType(ctx, db, id)
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

func mapEntity(ent *repository.WorldEntity) (*entities.Entity, error) {
	fullEntity := entities.Entity{
		ID:           ent.ID,
		ResourceName: ent.Wbrn,
		Name:         ent.EntityName,
		Description:  ent.EntityDescription,
		Notes:        ent.Notes.String,
		CreatedAt:    ent.CreatedAt.Time,
		UpdatedAt:    ent.UpdatedAt.Time,
	}

	return &fullEntity, nil
}

func (e *EntityService) getTypeAttributes(ctx context.Context, db repository.DBTX, typeID uuid.UUID) (map[uuid.UUID]entities.Attribute, error) {
	defs, err := e.queries.GetAttributesForType(ctx, db, typeID)
	if err != nil {
		return nil, fmt.Errorf("GetAttributesForType: %w", err)
	}

	attribs := make(map[uuid.UUID]entities.Attribute)
	for _, d := range defs {
		attribs[d.ID] = entities.Attribute{
			ID:            d.ID,
			Wbatn:         d.Wbatn,
			AttributeName: d.AttributeName,
			Label:         d.Label,
			DataType:      d.DataType,
			CreatedAt:     d.CreatedAt,
			UpdatedAt:     d.UpdatedAt,
		}
	}

	return attribs, nil
}

func (e *EntityService) UpdateByID(id uuid.UUID) (*entities.Entity, error) {
	return nil, nil
}
