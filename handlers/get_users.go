package handlers

import (
	"card-project/models"
	"card-project/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetUsers(params operations.GetUsersParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	user, err := h.controller.GetUsers(ctx)

	if err != nil {
		return operations.NewGetUsersDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET Users in storage",
			},
		})
	}

	return operations.NewGetUsersOK().WithPayload(user)
}
