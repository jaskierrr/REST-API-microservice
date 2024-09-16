package handlers

import (
	"card-project/models"
	"card-project/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetCards(params operations.GetCardsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	card, err := h.controller.GetCards(ctx)

	if err != nil {
		return operations.NewGetCardsDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET Cards in storage " + err.Error(),
			},
		})
	}

	return operations.NewGetCardsOK().WithPayload(card)
}

func (h *handlers) GetCardsID(params operations.GetCardsIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	card, err := h.controller.GetCardID(ctx, params.ID)

	if err != nil {
		return operations.NewGetCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET Card in storage, card id: " + params.ID + " " + err.Error(),
			},
		})
	}

	return operations.NewGetCardsIDOK().WithPayload(&card)
}

func (h *handlers) DeleteCardsID(params operations.DeleteCardsIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	commandTag, err := h.controller.DeleteCardID(ctx, params.ID)

	if commandTag.RowsAffected() == 0 {
		return operations.NewDeleteCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "card not found, card id: " + params.ID + " " + err.Error(),
			},
		})
	}

	if err != nil {
		return operations.NewDeleteCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE Card in storage, card id: " + params.ID + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteCardsIDNoContent()
}

func (h *handlers) PostCards(params operations.PostCardsParams) middleware.Responder {
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
