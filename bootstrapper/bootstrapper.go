package bootstrapper

import (
	"card-project/config"
	"card-project/controller"
	"card-project/database"
	"card-project/handlers"
	"card-project/logger"
	cards_repo "card-project/repositories/cards"
	"card-project/repositories/rabbitmq"
	users_repo "card-project/repositories/users"
	"card-project/restapi"
	"card-project/restapi/operations"
	"card-project/service"
	"context"
	"log"
	"log/slog"

	"github.com/go-openapi/loads"
	"github.com/go-playground/validator/v10"
)

const (
	dbConfigString = "postgres://%s:%s@%s:%s/%s"
	rabbitConfigString = "amqp://%s:%s@%s:%s/"
)

type RootBootstrapper struct {
	Infrastructure struct {
		Logger *slog.Logger
		Server *restapi.Server
		DB     database.DB
	}
	Controller     controller.Controller
	Config         *config.Config
	Handlers       handlers.Handlers
	UserRepository users_repo.UsersRepo
	CardRepository cards_repo.CardsRepo
	RabbitMQ       rabbitmq.RabbitMQ
	Service        service.Service

	Validator *validator.Validate
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
	r.Validator = validator.New(validator.WithRequiredStructEnabled())

	r.Handlers = handlers.New(r.Controller, r.Validator)
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
	r.Infrastructure.Logger = logger.NewLogger()

	r.Infrastructure.DB = database.NewDB().NewConn(ctx, dbConfigString, *r.Config)
	r.UserRepository = users_repo.NewUserRepo(r.Infrastructure.DB)
	r.CardRepository = cards_repo.NewCardRepo(r.Infrastructure.DB)
	r.RabbitMQ = rabbitmq.NewRabbitMQ().NewConn(r.UserRepository, r.CardRepository, rabbitConfigString, *r.Config)
	go r.RabbitMQ.NewConsumer(ctx)
	r.Service = service.New(r.UserRepository, r.CardRepository, r.RabbitMQ)

	// r.registerRepositoriesAndServices(r.Infrastructure.DB)
	err := r.registerAPIServer(*r.Config)
	if err != nil {
		log.Fatal("cant start server")
	}

	return nil
}
