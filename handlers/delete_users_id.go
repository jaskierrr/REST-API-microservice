package handlers

import (
	"card-project/models"
	"card-project/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) DeleteUsersID(params operations.DeleteUsersIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	commandTag, err := h.controller.DeleteUserID(ctx, params.ID)

	if commandTag.RowsAffected() == 0 {
		return operations.NewDeleteUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "user not found, user id: " + params.ID,
			},
		})
	}

	if err != nil {
		return operations.NewDeleteUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE User in storage, user id: " + params.ID,
			},
		})
	}

	return operations.NewDeleteUsersIDNoContent()
}
