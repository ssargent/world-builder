package service

import (
	"github.com/google/uuid"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"

	. "github.com/ssargent/world-builder/wb-api-go/internal/drivers"
)

type EntityService struct {
	db         DB
	repository EntityRepository
}

func NewEntityService(db DB) *EntityService {
	return &EntityService{}
}

func (e *EntityService) FindByID(id uuid.UUID) (*entities.Entity, error) {
	return e.repository.FindByID(e.db, id)
}
