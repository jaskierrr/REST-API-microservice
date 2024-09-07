package main

import (
	"log"

	"github.com/go-openapi/loads"

	"card-project/handlers"
	"card-project/restapi"
	"card-project/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCardProjectAPI(swaggerSpec)
	server := restapi.NewServer(api)

	server.Port = 8080

	defer server.Shutdown()

	h := handlers.New()
	h.Init(api)

	if h == nil {
		log.Fatal("handlers initialization failed")
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
