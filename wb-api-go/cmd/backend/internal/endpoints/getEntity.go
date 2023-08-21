package endpoints

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	entityv1 "github.com/ssargent/apis/pkg/worldbuilder/entity/v1"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

func (e *EntityServer) GetEntity(
	ctx context.Context,
	req *connect.Request[entityv1.GetEntityRequest]) (*connect.Response[entityv1.GetEntityResponse], error) {
	id, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	ent, err := e.service.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("FindByID: %w", err)
	}

	res := connect.NewResponse(&entityv1.GetEntityResponse{
		Entity: fromEntity(ent),
	})
	res.Header().Set("Entity-Version", "v1")
	return res, nil
}

func fromEntity(e *entities.Entity) *entityv1.Entity {
	ev1parent := entityv1.EntityParent{
		EntityId:     e.Parent.EntityID.String(),
		EntityName:   e.Parent.EntityName,
		ResourceName: e.Parent.ResourceName,
		TypeName:     e.Parent.TypeName,
	}

	ev1type := entityv1.EntityType{
		TypeId:   e.Type.TypeID.String(),
		TypeName: e.Type.TypeName,
	}

	ev1attribs := make([]*entityv1.EntityAttribute, len(e.Attributes))
	for i, a := range e.Attributes {
		ev1attribs[i] = &entityv1.EntityAttribute{
			Name:  a.Name,
			Type:  a.Type,
			Value: a.Value,
		}
	}

	ev1 := entityv1.Entity{
		Id:           e.ID.String(),
		Name:         e.Name,
		Description:  e.Description,
		ResourceName: e.ResourceName,
		CreatedAt:    e.CreatedAt.String(),
		UpdatedAt:    e.UpdatedAt.String(),

		Parent:     &ev1parent,
		Type:       &ev1type,
		Attributes: ev1attribs,
	}

	return &ev1
}
