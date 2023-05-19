package service

import (
	"github.com/google/uuid"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

type EntityService struct {
}

func NewEntityService() *EntityService {
	return &EntityService{}
}

func (e *EntityService) FindByID(id uuid.UUID) (*entities.Entity, error) {
	return nil, nil
}

func (e *EntityService) UpdateByID(id uuid.UUID) (*entities.Entity, error) {
	return nil, nil
}
