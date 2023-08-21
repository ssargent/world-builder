package endpoints

import "github.com/ssargent/world-builder/wb-api-go/internal/service"

type EntityServer struct {
	service *service.EntityService
	types   *service.TypeService
}

func NewEntityServer(svc *service.EntityService, types *service.TypeService) *EntityServer {
	return &EntityServer{
		service: svc,
		types:   types,
	}
}
