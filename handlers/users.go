package handlers

import (
	"card-project/models"
	"card-project/restapi/operations"
	"log/slog"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetUsers(params operations.GetUsersParams) middleware.Responder {
	h.logger.Info("Trying to GET users from storage")

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
	h.logger.Info("Trying to GET user from storage, user id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewGetUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET User from storage, user id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	user, err := h.controller.GetUserID(ctx, int(params.ID))

	if err != nil {
		return operations.NewGetUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET User from storage, user id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewGetUsersIDOK().WithPayload(&user)
}

func (h *handlers) DeleteUsersID(params operations.DeleteUsersIDParams) middleware.Responder {
	h.logger.Info("Trying to DELETE user from storage, user id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewDeleteUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE User from storage, user id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	err := h.controller.DeleteUserID(ctx, int(params.ID))

	if err != nil {
		return operations.NewDeleteUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE User from storage, user id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteUsersIDNoContent()
}

func (h *handlers) PostUsers(params operations.PostUsersParams) middleware.Responder {
	h.logger.Info(
		"Trying to POST user in storage",
		slog.Any("user", params.User),
	)

	err := validate.Struct(params.User)
	if err != nil {
		return operations.NewGetUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST User in storage " + err.Error(),
			},
		})
	}

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
