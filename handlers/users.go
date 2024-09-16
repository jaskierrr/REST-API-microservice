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
				Message: "Failed to GET Users in storage " + err.Error(),
			},
		})
	}

	return operations.NewGetUsersOK().WithPayload(user)
}

func (h *handlers) GetUsersID(params operations.GetUsersIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	user, err := h.controller.GetUserID(ctx, params.ID)

	if err != nil {
		return operations.NewGetUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET User in storage, user id: " + params.ID + " " + err.Error(),
			},
		})
	}

	return operations.NewGetUsersIDOK().WithPayload(&user)
}

func (h *handlers) DeleteUsersID(params operations.DeleteUsersIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	commandTag, err := h.controller.DeleteUserID(ctx, params.ID)

	if commandTag.RowsAffected() == 0 {
		return operations.NewDeleteUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "user not found, user id: " + params.ID + " " + err.Error(),
			},
		})
	}

	if err != nil {
		return operations.NewDeleteUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE User in storage, user id: " + params.ID + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteUsersIDNoContent()
}

func (h *handlers) PostUsers(params operations.PostUsersParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	user, err := h.controller.PostUser(ctx, *params.User)

	if err != nil {
		return operations.NewGetUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST User in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostUsersCreated().WithPayload(&user)
}
