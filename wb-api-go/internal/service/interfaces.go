package service

//go:generate mockgen -source=interfaces.go -destination mocks/service.go  github.com/ssargent/world-builder/wb-api-go/internal/service EntityDataProvider, Cache

import (
	"time"

	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
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
