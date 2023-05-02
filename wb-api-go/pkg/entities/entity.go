package entities

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID                uuid.UUID
	TypeID            uuid.UUID
	ParentID          uuid.UUID
	WBRN              string
	EntityName        string
	EntityDescription string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
