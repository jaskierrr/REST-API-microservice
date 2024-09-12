package handlers

import (
	"card-project/models"
	"card-project/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) PostUsers(params operations.PostUsersParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	user, err := h.controller.PostUser(ctx, *params.User)

	if err != nil {
		return operations.NewGetUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST User in storage",
			},
		})
	}

	return operations.NewPostUsersCreated().WithPayload(&user)
}
