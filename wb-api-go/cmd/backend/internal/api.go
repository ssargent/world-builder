package internal

import (
	"fmt"
	"net/http"

	"github.com/patrickmn/go-cache"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/config"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/endpoints"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/handlers"
	"github.com/ssargent/world-builder/wb-api-go/gen/api/entity/v1/entityv1connect"
	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	"github.com/ssargent/world-builder/wb-api-go/internal/service"
)

/*
func (s *GreetServer) Greet(

	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],

	) (*connect.Response[greetv1.GreetResponse], error) {
	    log.Println("Request headers: ", req.Header())
	    res := connect.NewResponse(&greetv1.GreetResponse{
	        Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	    })
	    res.Header().Set("Greet-Version", "v1")
	    return res, nil
	}
*/

type API struct {
	cfg    *config.Config
	Reader *sqlx.DB
	Writer *sqlx.DB
	Entity *service.EntityService
}

func NewApi(cfg *config.Config, rdb *sqlx.DB, wdb *sqlx.DB, cache *cache.Cache) *API {
	q := repository.Queries{}
	entity := service.NewEntityService(cache, rdb, wdb, &q)
	return &API{
		cfg:    cfg,
		Reader: rdb,
		Writer: wdb,
		Entity: entity,
	}
}

func (a *API) ListenAndServe() error {
	r := chi.NewRouter()

	mux := http.NewServeMux()

	rpcServer := endpoints.NewEntityServer(a.Entity)
	path, handler := entityv1connect.NewEntityServiceHandler(rpcServer)
	mux.Handle(path, handler)

	h := handlers.NewHandler(a.cfg, a.Entity)

	r.Mount("/v1", h.Routes())
	r.Mount("/grpc", h2c.NewHandler(mux, &http2.Server{}))

	return http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.Port), r)
}
