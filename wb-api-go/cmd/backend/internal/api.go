package internal

import (
	"fmt"
	"net/http"
	"strings"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/apis/pkg/worldbuilder/entity/v1/entityv1connect"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/config"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/endpoints"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/handlers"
	"github.com/ssargent/world-builder/wb-api-go/internal/repository"
	"github.com/ssargent/world-builder/wb-api-go/internal/service"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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
	cache  *cache.Cache
	Reader *sqlx.DB
	Writer *sqlx.DB
	Entity *service.EntityService
}

func NewAPI(cfg *config.Config, rdb *sqlx.DB, wdb *sqlx.DB, cache *cache.Cache) *API {
	q := repository.Queries{}
	entity := service.NewEntityService(cache, rdb, wdb, &q)
	return &API{
		cfg:    cfg,
		cache:  cache,
		Reader: rdb,
		Writer: wdb,
		Entity: entity,
	}
}

func (a *API) ListenAndServe() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	mux := http.NewServeMux()

	rpcServer := endpoints.NewEntityServer(a.Entity)
	path, handler := entityv1connect.NewEntityServiceHandler(rpcServer)
	mux.Handle(path, handler)
	reflector := grpcreflect.NewStaticReflector(
		strings.ReplaceAll(path, "/", ""),
		// protoc-gen-connect-go generates package-level constants
		// for these fully-qualified protobuf service names, so you'd more likely
		// reference userv1.UserServiceName and groupv1.GroupServiceName.
	)
	r1path, r1handler := grpcreflect.NewHandlerV1(reflector)
	r2path, r2handler := grpcreflect.NewHandlerV1Alpha(reflector)

	h := handlers.NewHandler(a.cfg, a.cache, a.Entity)

	r.Mount("/v1", h.Routes())
	r.Mount(path, h2c.NewHandler(mux, &http2.Server{}))

	r.Handle(r1path, r1handler)
	r.Handle(r2path, r2handler)
	r.Handle("/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo", r2handler)

	walkFunc := func(method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler) error {
		route = strings.ReplaceAll(route, "/*/", "/")
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	h2s := &http2.Server{}
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", a.cfg.Port),
		Handler: h2c.NewHandler(r, h2s),
	}

	return srv.ListenAndServe()
}
