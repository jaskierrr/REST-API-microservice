package handlers

import (
	"card-project/controller"
	"card-project/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type handlers struct {
	controller controller.Controller
}

type Handlers interface {
	GetUsersID(params operations.GetUsersIDParams) middleware.Responder
	PostUsers(params operations.PostUsersParams) middleware.Responder
	DeleteUsersID(params operations.DeleteUsersIDParams) middleware.Responder
	GetUsers(params operations.GetUsersParams) middleware.Responder
	Link(api *operations.CardProjectAPI)
}

func New(controller controller.Controller) Handlers {
	return &handlers{
		controller: controller,
	}
}

func (h *handlers) Link(api *operations.CardProjectAPI) {
	api.GetUsersIDHandler = operations.GetUsersIDHandlerFunc(h.GetUsersID)
	api.PostUsersHandler = operations.PostUsersHandlerFunc(h.PostUsers)
	api.DeleteUsersIDHandler = operations.DeleteUsersIDHandlerFunc(h.DeleteUsersID)
	api.GetUsersHandler = operations.GetUsersHandlerFunc(h.GetUsers)
}
