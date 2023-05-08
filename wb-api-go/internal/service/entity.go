package service

import (
	"github.com/google/uuid"
	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

type EntityService struct {
	db         repository.DB
	repository EntityRepository
}

func NewEntityService(db repository.DB) *EntityService {
	return &EntityService{}
}

func (e *EntityService) FindByID(id uuid.UUID) (*entities.Entity, error) {
	return e.repository.FindByID(e.db, id)
}
