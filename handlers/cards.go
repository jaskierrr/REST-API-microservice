package handlers

import (
	"card-project/models"
	"card-project/restapi/operations"
	"log/slog"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetCards(params operations.GetCardsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	card, err := h.controller.GetCards(ctx)

	if err != nil {
		return operations.NewGetCardsDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET Cards from storage " + err.Error(),
			},
		})
	}

	return operations.NewGetCardsOK().WithPayload(card)
}

func (h *handlers) GetCardsID(params operations.GetCardsIDParams) middleware.Responder {
	h.logger.Info("Trying to GET card from storage, user id: " + convertI64tStr(params.ID))

	ctx := params.HTTPRequest.Context()
	card, err := h.controller.GetCardID(ctx, int(params.ID))

	if err != nil {
		return operations.NewGetCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET Card from storage, card id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewGetCardsIDOK().WithPayload(&card)
}

func (h *handlers) DeleteCardsID(params operations.DeleteCardsIDParams) middleware.Responder {
	h.logger.Info("Trying to DELETE card from storage, user id: " + convertI64tStr(params.ID))

	ctx := params.HTTPRequest.Context()
	err := h.controller.DeleteCardID(ctx, int(params.ID))

	if err != nil {
		h.logger.Info("Failed to DELETE Card from storage, card id: " + convertI64tStr(params.ID) + " " + err.Error())
		return operations.NewDeleteCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE Card from storage, card id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteCardsIDNoContent()
}

func (h *handlers) PostCards(params operations.PostCardsParams) middleware.Responder {
	h.logger.Info(
		"Trying to POST user in storage",
		slog.Any("user", params.Card),
	)

	err := validate.Struct(params.Card)
	if err != nil {
		return operations.NewGetCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST Card in storage " + err.Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	card, err := h.controller.PostCard(ctx, *params.Card)

	if err != nil {
		return operations.NewGetCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST Card in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostCardsCreated().WithPayload(&card)
}
