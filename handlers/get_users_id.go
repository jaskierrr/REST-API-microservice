package handlers

import (
	"card-project/models"
	"card-project/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetUsersID(params operations.GetUsersIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	user, err := h.controller.GetUserID(ctx, params.ID)

	if err != nil {
		return operations.NewGetUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET User in storage, user id: " + params.ID,
			},
		})
	}

	return operations.NewGetUsersIDOK().WithPayload(&user)
}
