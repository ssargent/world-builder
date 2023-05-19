// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type WorldAttributeDefinition struct {
	ID            uuid.UUID    `json:"id"`
	Wbatn         string       `json:"wbatn"`
	AttributeName string       `json:"attribute_name"`
	Label         string       `json:"label"`
	DataType      string       `json:"data_type"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
}

type WorldEntity struct {
	ID                uuid.UUID      `json:"id"`
	TypeID            uuid.UUID      `json:"type_id"`
	ParentID          uuid.UUID      `json:"parent_id"`
	Wbrn              string         `json:"wbrn"`
	EntityName        string         `json:"entity_name"`
	EntityDescription string         `json:"entity_description"`
	Notes             sql.NullString `json:"notes"`
	CreatedAt         sql.NullTime   `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
}

type WorldEntityAttribute struct {
	ID             uuid.UUID    `json:"id"`
	EntityID       uuid.UUID    `json:"entity_id"`
	AttributeID    uuid.UUID    `json:"attribute_id"`
	AttributeValue string       `json:"attribute_value"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}

type WorldType struct {
	ID              uuid.UUID    `json:"id"`
	ParentID        uuid.UUID    `json:"parent_id"`
	Wbtn            string       `json:"wbtn"`
	TypeName        string       `json:"type_name"`
	TypeDescription string       `json:"type_description"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

type WorldTypeAttribute struct {
	TypeID      uuid.UUID `json:"type_id"`
	AttributeID uuid.UUID `json:"attribute_id"`
	Ordinal     int32     `json:"ordinal"`
	IsRequired  bool      `json:"is_required"`
}
