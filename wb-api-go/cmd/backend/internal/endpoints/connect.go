package endpoints

import "github.com/ssargent/world-builder/wb-api-go/internal/service"

type EntityServer struct {
	service *service.EntityService
}

func NewEntityServer(svc *service.EntityService) *EntityServer {
	return &EntityServer{
		service: svc,
	}
}
