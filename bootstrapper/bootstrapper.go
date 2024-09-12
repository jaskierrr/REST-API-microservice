package bootstrapper

import (
	"card-project/config"
	"card-project/controller"
	"card-project/database"
	"card-project/handlers"
	repositories "card-project/repositories/users"
	"card-project/restapi"
	"card-project/restapi/operations"
	"context"
	"log"

	"github.com/go-openapi/loads"
)

const connConfigString = "postgres://%s:%s@%s:%s/%s"

type RootBootstrapper struct {
	Infrastructure struct {
		// Logger
		Server *restapi.Server
		DB database.DB
	}
	Controller controller.Controller
	Config *config.Config
	Handlers handlers.Handlers
	Repository repositories.UsersRepo
}

type RootBoot interface {
	registerAPIServer(cfg config.Config) error
	// registerRepositoties(db database.DB) error
	RunAPI() error
}

func New() RootBoot{
	return RootBootstrapper{}
}


func (r RootBootstrapper) registerAPIServer(cfg config.Config) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := operations.NewCardProjectAPI(swaggerSpec)

	r.Controller = controller.New(r.Repository)


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

// func (r RootBootstrapper) registerRepositoties(db database.DB) error{
// 	r.Repository = repositories.NewUserRepo(db)

// 	return nil
// }


func (r RootBootstrapper) RunAPI() error{
	ctx := context.Background()
	r.Config = config.NewConfig()

	r.Infrastructure.DB = database.NewDB().NewConn(ctx, connConfigString, *r.Config)

	// err := r.registerRepositoties(r.Infrastructure.DB)
	r.Repository = repositories.NewUserRepo(r.Infrastructure.DB)


	err := r.registerAPIServer(*r.Config)
	if err != nil {
		log.Fatal("cant start server")
	}




	return nil
}
