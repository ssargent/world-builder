package endpoints

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	entityv1 "github.com/ssargent/apis/pkg/worldbuilder/entity/v1"
	"github.com/ssargent/world-builder/wb-api-go/internal/conversions"
)

func (e *EntityServer) CreateType(ctx context.Context,
	req *connect.Request[entityv1.CreateTypeRequest]) (*connect.Response[entityv1.CreateTypeResponse], error) {
	entityType, err := conversions.EntityType(req.Msg.Type)
	if err != nil {
		return nil, fmt.Errorf("conversions.EntityType: %w", err)
	}

	created, err := e.types.CreateType(ctx, entityType)
	if err != nil {
		return nil, fmt.Errorf("types.CreateType: %w", err)
	}

	fmt.Printf("Created: %+v\n", created)

	protoType, err := conversions.ProtoType(created)
	if err != nil {
		return nil, fmt.Errorf("conversions.ProtoType: %w", err)
	}

	res := connect.NewResponse(&entityv1.CreateTypeResponse{
		Type: protoType,
	})

	return res, nil
}
