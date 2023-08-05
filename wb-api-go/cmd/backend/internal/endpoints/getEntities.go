package endpoints

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	entityv1 "github.com/ssargent/apis/pkg/worldbuilder/entity/v1"
)

func (e *EntityServer) GetEntities(
	ctx context.Context,
	req *connect.Request[entityv1.GetEntitiesRequest]) (*connect.Response[entityv1.GetEntitiesResponse], error) {
	ents, err := e.service.FilterByCriteria(
		ctx,
		safeString(req.Msg.Criteria.TypeName),
		safeString(req.Msg.Criteria.ParentWbrn))
	if err != nil {
		return nil, fmt.Errorf("FilterByCriteria: %w", err)
	}

	endpointEnts := make([]*entityv1.Entity, len(ents))
	for i, sent := range ents {
		endpointEnts[i] = fromEntity(sent)
	}

	res := connect.NewResponse(&entityv1.GetEntitiesResponse{
		Criteria: req.Msg.Criteria,
		Entities: endpointEnts,
	})
	res.Header().Set("Entity-Version", "v1")

	return res, nil
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
