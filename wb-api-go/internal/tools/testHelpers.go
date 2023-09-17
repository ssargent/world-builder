package tools

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

func UUID(id int8) uuid.UUID {
	empty := make([]byte, 16)
	empty[0] = byte(id)
	val, _ := uuid.FromBytes(empty)

	return val
}

func Entity(id, parent, entityType int8) *repository.WorldEntity {
	ent := repository.WorldEntity{
		ID:                UUID(id),
		ParentID:          UUID(parent),
		TypeID:            UUID(entityType),
		EntityName:        fmt.Sprintf("Entity-%d", id),
		EntityDescription: fmt.Sprintf("Entity-%d-description", id),
	}
	return &ent
}

func EntityReference(id, entityType int8) *repository.GetEntityReferenceRow {
	return &repository.GetEntityReferenceRow{
		EntityID:     UUID(id),
		EntityName:   fmt.Sprintf("entity-%d", id),
		ResourceName: fmt.Sprintf("wbrn:entity-%d", id),
		TypeName:     fmt.Sprintf("wbtn:type-%d", entityType),
	}
}

func EntityType(id, parent int8) *repository.WorldType {
	return &repository.WorldType{
		ID:              UUID(id),
		ParentID:        UUID(parent),
		Wbtn:            fmt.Sprintf("wbtn:type-%d", id),
		TypeName:        fmt.Sprintf("type-%d", id),
		TypeDescription: fmt.Sprintf("type-%d", id),
	}
}

func AttributesForType(id int8) []*repository.WorldAttributeDefinition {
	defs := make([]*repository.WorldAttributeDefinition, 0)

	defs = append(defs, &repository.WorldAttributeDefinition{
		ID:            UUID(10),
		Wbatn:         fmt.Sprintf("wban:test"),
		AttributeName: fmt.Sprintf("attribute%d-1", id),
		Label:         fmt.Sprintf("attribute%d-1", id),
		DataType:      "string",
		CreatedAt:     sql.NullTime{Time: time.Now()},
		UpdatedAt:     sql.NullTime{Time: time.Now()},
	})

	defs = append(defs, &repository.WorldAttributeDefinition{
		ID:            UUID(11),
		Wbatn:         fmt.Sprintf("wban:test2"),
		AttributeName: fmt.Sprintf("attribute%d-2", id),
		Label:         fmt.Sprintf("attribute%d-2", id),
		DataType:      "string",
		CreatedAt:     sql.NullTime{Time: time.Now()},
		UpdatedAt:     sql.NullTime{Time: time.Now()},
	})

	return defs
}

func TypeAttribute(tid, id int8, ord int32, required bool) *repository.WorldTypeAttribute {
	return &repository.WorldTypeAttribute{
		TypeID:      UUID(tid),
		AttributeID: UUID(id),
		Ordinal:     ord,
		IsRequired:  required,
	}
}

func Attribute(id int8, wbatn string) *entities.Attribute {
	return &entities.Attribute{
		ID:            UUID(id),
		Wbatn:         wbatn,
		AttributeName: wbatn,
		Label:         "Label",
		DataType:      "dt",
	}
}

func EntityAttributes(id int8) []*repository.WorldEntityAttribute {
	entityAttributes := make([]*repository.WorldEntityAttribute, 0)

	entityAttributes = append(entityAttributes, &repository.WorldEntityAttribute{
		ID:             uuid.New(),
		EntityID:       UUID(id),
		AttributeID:    UUID(10),
		AttributeValue: "echo",
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	})
	entityAttributes = append(entityAttributes, &repository.WorldEntityAttribute{
		ID:             uuid.New(),
		EntityID:       UUID(id),
		AttributeID:    UUID(11),
		AttributeValue: "ohce",
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	})

	return entityAttributes
}
