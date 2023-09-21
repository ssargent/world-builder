package service

//go:generate mockgen -source=interfaces.go -destination mocks/service.go  github.com/ssargent/world-builder/wb-api-go/internal/service EntityDataProvider, Cache, WriterDB, ReaderDB, TypeService

import (
	"context"
	"time"

	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	"github.com/ssargent/world-builder/wb-api-go/pkg/entities"
)

type EntityDataProvider interface {
	repository.AttributeDefinitionQuerier
	repository.EntityAssociationQuerier
	repository.EntityAttributeQuerier
	repository.EntityQuerier
	repository.EntityHistoryQuerier
	repository.ReferenceQuerier
	repository.TypeQuerier
}

type Cache interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
}

type TypeService interface {
	CreateType(ctx context.Context, in *entities.EntityType) (*entities.EntityType, error)
	GetType(ctx context.Context, typeRef *entities.TypeReference) (*entities.EntityType, error)
}
