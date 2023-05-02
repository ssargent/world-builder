package service

import (
	"github.com/google/uuid"
	. "github.com/ssargent/world-builder/wb-api-go/internal/drivers"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

type EntityRepository interface {
	FindByID(db DB, id uuid.UUID) (*entities.Entity, error)
	FindByWBRN(db DB, wbrn string) (*entities.Entity, error)
	FindByTypeName(db DB, typeName string) (*entities.Entity, error)
}
