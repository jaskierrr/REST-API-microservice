package handlers_test

import (
	"card-project/controller"
	"card-project/handlers"
	"card-project/logger"
	mock "card-project/mocks"
	"card-project/models"
	"card-project/restapi/operations"
	"card-project/service"
	"context"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

func Test_GetCardID(t *testing.T) {
	type fields struct {
		rabbitmq *mock.MockRabbitMQ
		cardRepo *mock.MockCardsRepo
		userRepo *mock.MockUsersRepo
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLogger()

	rabbitmqMock := mock.NewMockRabbitMQ(ctrl)
	cardRepoMock := mock.NewMockCardsRepo(ctrl)
	userRepoMock := mock.NewMockUsersRepo(ctrl)

	testFields := &fields{
		rabbitmq: rabbitmqMock,
		cardRepo: cardRepoMock,
		userRepo: userRepoMock,
	}

	service := service.New(userRepoMock, cardRepoMock, rabbitmqMock, logger)
	controller := controller.New(service, logger)
	h := handlers.New(controller, validator.New(validator.WithRequiredStructEnabled()), logger)

	card := models.Card{
		ID:         1,
		UserID:     1,
		BankID:     1,
		Number:     1,
		CreateDate: strfmt.DateTime(time.Now()),
	}

	reqArgDef := operations.GetCardsIDParams{
		HTTPRequest: &http.Request{},
		ID:          1,
	}

	reqArgErr := operations.GetCardsIDParams{
		HTTPRequest: &http.Request{},
		ID:          111,
	}

	resErr := operations.NewGetCardsIDDefault(404).WithPayload(&models.ErrorResponse{
		Error: &models.ErrorResponseAO0Error{
			Message: "Failed to GET Card in storage, card id: " + strconv.FormatInt(reqArgErr.ID, 10) + " " + errors.New("no rows in result set").Error(),
		},
	})

	ctx := context.Background()

	tests := []struct {
		name    string
		args    operations.GetCardsIDParams
		prepare func(f *fields)
		wantRes middleware.Responder
	}{
		{
			name: "valid",
			args: reqArgDef,
			prepare: func(f *fields) {
				// если указанные вызовы не станут выполняться в ожидаемом порядке, тест будет провален
				gomock.InOrder(
					f.cardRepo.EXPECT().GetCardID(ctx, 1).Return(card, nil),
				)
			},
			wantRes: operations.NewGetCardsIDOK().WithPayload(&card),
		},
		{
			name: "wrong_ID",
			args: reqArgErr,
			prepare: func(f *fields) {
				gomock.InOrder(
					f.cardRepo.EXPECT().GetCardID(ctx, 111).Return(models.Card{}, errors.New("no rows in result set")),
				)
			},
			wantRes: resErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare(testFields)
			}

			response := h.GetCardsID(tt.args)

			if !reflect.DeepEqual(response, tt.wantRes) {
				t.Errorf("GetCard() = %v, want %v", response, tt.wantRes)
			}
		})
	}

}
