package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type EntityReference struct {
	EntityID     uuid.UUID `json:"entity_id,omitempty"`
	EntityName   string    `json:"entity_name,omitempty"`
	ResourceName string    `json:"resource_name,omitempty"`
	TypeName     string    `json:"type_name,omitempty"`
}

type Entity struct {
	ID           uuid.UUID `json:"id,omitempty"`
	ResourceName string    `json:"resource_name,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Notes        string    `json:"notes,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`

	Attributes []*EntityAttribute `json:"attributes,omitempty"`
	Children   []*EntityReference `json:"children,omitempty"`
	Parent     *EntityReference   `json:"parent,omitempty"`
	Type       *TypeReference     `json:"type,omitempty"`
}

type Attribute struct {
	ID            uuid.UUID    `json:"id"`
	Wbatn         string       `json:"wbatn"`
	AttributeName string       `json:"attribute_name"`
	Label         string       `json:"label"`
	DataType      string       `json:"data_type"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
}

type EntityAttribute struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
type TypeReference struct {
	TypeID   uuid.UUID `json:"type_id,omitempty"`
	TypeName string    `json:"type_name,omitempty"`
}

type EntityType struct {
	ID              uuid.UUID `json:"id,omitempty"`
	Wbtn            string    `json:"wbtn,omitempty"`
	TypeName        string    `json:"type_name,omitempty"`
	TypeDescription string    `json:"type_description,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`

	Attributes []*Attribute   `json:"attributes,omitempty"`
	Parent     *TypeReference `json:"parent,omitempty"`
}
