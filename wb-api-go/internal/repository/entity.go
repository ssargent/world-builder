package repository

import (
	"github.com/google/uuid"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

type EntityRepository struct {
}

func NewEntityRepository() *EntityRepository {
	return &EntityRepository{}
}

func (e *EntityRepository) FindByID(db DB, id uuid.UUID) (*entities.Entity, error) {
	return nil, nil
}

func (e *EntityRepository) FindByWBRN(db DB, wbrn string) (*entities.Entity, error) {
	return nil, nil
}

func (e *EntityRepository) FindByTypeName(db DB, typeName string) (*entities.Entity, error) {
	return nil, nil
}
