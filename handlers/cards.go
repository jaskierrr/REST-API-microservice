package handlers

import (
	"card-project/models"
	"card-project/restapi/operations"
	"log/slog"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetCards(params operations.GetCardsParams) middleware.Responder {
	h.logger.Info("Trying to GET cards from storage")

	ctx := params.HTTPRequest.Context()
	card, err := h.controller.GetCards(ctx)

	if err != nil {
		h.logger.Error(
			"Failed to GET cards from storage",
			slog.String("error", err.Error()),
		)
		return operations.NewGetCardsDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET Cards from storage " + err.Error(),
			},
		})
	}

	return operations.NewGetCardsOK().WithPayload(card)
}

func (h *handlers) GetCardsID(params operations.GetCardsIDParams) middleware.Responder {
	h.logger.Info("Trying to GET card from storage, card id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewGetCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET card from storage, card id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	card, err := h.controller.GetCardID(ctx, int(params.ID))

	if err != nil {
		h.logger.Error(
			"Failed to GET Card from storage",
			slog.String("ID", convertI64tStr(params.ID)),
			slog.String("error", err.Error()),
		)
		return operations.NewGetCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET card from storage, card id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewGetCardsIDOK().WithPayload(&card)
}

func (h *handlers) DeleteCardsID(params operations.DeleteCardsIDParams) middleware.Responder {
	h.logger.Info("Trying to DELETE card from storage, card id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewDeleteCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE card from storage, card id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	err := h.controller.DeleteCardID(ctx, int(params.ID))

	if err != nil {
		h.logger.Error(
			"Failed to DELETE card from storage",
			slog.String("ID", convertI64tStr(params.ID)),
			slog.String("error", err.Error()),
		)
		return operations.NewDeleteCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE card from storage, card id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteCardsIDNoContent()
}

func (h *handlers) PostCards(params operations.PostCardsParams) middleware.Responder {
	h.logger.Info(
		"Trying to POST card in storage",
		slog.Any("card", params.Card),
	)

	err := validate.Struct(params.Card)
	if err != nil {
		h.logger.Error(
			"Failed to POST card in storage",
			slog.Any("card", params.Card),
			slog.String("error", err.Error()),
		)
		return operations.NewGetCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST card in storage " + err.Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	card, err := h.controller.PostCard(ctx, *params.Card)

	if err != nil {
		h.logger.Error(
			"Failed to POST card in storage",
			slog.Any("card", params.Card),
			slog.String("error", err.Error()),
		)
		return operations.NewGetCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST card in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostCardsCreated().WithPayload(&card)
}
