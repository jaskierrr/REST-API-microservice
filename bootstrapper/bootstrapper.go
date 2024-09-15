package bootstrapper

import (
	"card-project/config"
	"card-project/controller"
	"card-project/database"
	"card-project/handlers"
	repositories "card-project/repositories/users"
	"card-project/restapi"
	"card-project/restapi/operations"
	"card-project/service"
	"context"
	"log"

	"github.com/go-openapi/loads"
)

const connConfigString = "postgres://%s:%s@%s:%s/%s"

type RootBootstrapper struct {
	Infrastructure struct {
		// Logger
		Server *restapi.Server
		DB     database.DB
	}
	Controller controller.Controller
	Config     *config.Config
	Handlers   handlers.Handlers
	Repository repositories.UsersRepo
	Service    service.Service
}

type RootBoot interface {
	registerAPIServer(cfg config.Config) error
	// registerRepositoriesAndServices(db database.DB)
	RunAPI() error
}

func New() RootBoot {
	return RootBootstrapper{
		Config: config.NewConfig(),
	}
}

func (r RootBootstrapper) registerAPIServer(cfg config.Config) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := operations.NewCardProjectAPI(swaggerSpec)

	r.Controller = controller.New(r.Service)

	r.Handlers = handlers.New(r.Controller)
	r.Handlers.Link(api)
	if r.Handlers == nil {
		log.Fatal("handlers initialization failed")
	}

	r.Infrastructure.Server = restapi.NewServer(api)
	r.Infrastructure.Server.Port = cfg.ServerPort
	r.Infrastructure.Server.ConfigureAPI()
	if err := r.Infrastructure.Server.Serve(); err != nil {
		log.Fatalln(err)
	}

	return nil
}

// func (r RootBootstrapper) registerRepositoriesAndServices(db database.DB) {
// 	r.Repository = repositories.NewUserRepo(db)
// 	r.Service = service.New(r.Repository)
// }

func (r RootBootstrapper) RunAPI() error {
	ctx := context.Background()

	r.Infrastructure.DB = database.NewDB().NewConn(ctx, connConfigString, *r.Config)

	// r.registerRepositoriesAndServices(r.Infrastructure.DB)

	r.Repository = repositories.NewUserRepo(r.Infrastructure.DB)
	r.Service = service.New(r.Repository)

	err := r.registerAPIServer(*r.Config)
	if err != nil {
		log.Fatal("cant start server")
	}

	return nil
}
