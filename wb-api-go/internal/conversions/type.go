package conversions

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"

	entityv1 "github.com/ssargent/apis/pkg/worldbuilder/entity/v1"
)

func EntityType(in *entityv1.Type) (*entities.EntityType, error) {
	var e entities.EntityType

	TypeID, err := uuid.Parse(in.TypeId)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse(typeid): %w", err)
	}

	parentTypeID, err := uuid.Parse(in.Parent.TypeId)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse(parent.typeid): %w", err)
	}

	e.ID = TypeID
	e.Wbtn = in.Wbtn
	e.TypeName = in.TypeName
	e.TypeDescription = in.TypeDescription
	e.CreatedAt = time.Unix(in.CreatedAt, 0)
	e.UpdatedAt = time.Unix(in.UpdatedAt, 0)
	e.Parent = &entities.TypeReference{
		TypeID:   parentTypeID,
		TypeName: in.Parent.TypeName,
	}

	if len(in.Attributes) > 0 {
		e.Attributes = make([]*entities.Attribute, len(in.Attributes))

		for i, a := range in.Attributes {
			e.Attributes[i] = &entities.Attribute{
				Wbatn:         a.Wbatn,
				AttributeName: a.AttributeName,
			}
		}
	}

	return &e, nil
}

func ProtoType(in *entities.EntityType) (*entityv1.Type, error) {
	var e entityv1.Type

	e.TypeId = in.ID.String()
	e.Wbtn = in.Wbtn
	e.TypeName = in.TypeName
	e.TypeDescription = in.TypeDescription
	e.CreatedAt = in.CreatedAt.Unix()
	e.UpdatedAt = in.UpdatedAt.Unix()
	e.Parent = &entityv1.TypeParent{
		TypeId:   in.Parent.TypeID.String(),
		TypeName: in.Parent.TypeName,
	}

	if len(in.Attributes) > 0 {
		e.Attributes = make([]*entityv1.TypeAttribute, len(in.Attributes))

		for i, a := range in.Attributes {
			e.Attributes[i] = &entityv1.TypeAttribute{
				Wbatn:         a.Wbatn,
				AttributeName: a.AttributeName,
			}
		}
	}

	return &entityv1.Type{}, nil
}
