package handlers

import (
	"card-project/controller"
	"card-project/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type handlers struct{
	controller controller.Controller
}

type Handlers interface {
	GetUsersID(params operations.GetUsersIDParams) middleware.Responder
	Init(api *operations.CardProjectAPI)
}

func New() Handlers {
	return &handlers{
		controller: controller.New(),
	}
}

func (h *handlers) Init(api *operations.CardProjectAPI) {
	api.GetUsersIDHandler = operations.GetUsersIDHandlerFunc(h.GetUsersID)
}
