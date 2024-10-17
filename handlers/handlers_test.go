package handlers

import (
	mock "card-project/mocks"
	"card-project/models"
	"card-project/restapi/operations"
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
)

func TestHandlers_GetCard(t *testing.T) {
	type fields struct {
		controller *mock.MockController
		rabbitmq   *mock.MockRabbitMQ
		service    *mock.MockService
		cardRepo   *mock.MockCardsRepo
		userRepo   *mock.MockUsersRepo
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controllerMock := mock.NewMockController(ctrl)
	rabbitmqMock := mock.NewMockRabbitMQ(ctrl)
	serviceMock := mock.NewMockService(ctrl)
	cardRepoMock := mock.NewMockCardsRepo(ctrl)
	userRepoMock := mock.NewMockUsersRepo(ctrl)

	testFields := &fields{
		controller: controllerMock,
		rabbitmq:   rabbitmqMock,
		service:    serviceMock,
		cardRepo:   cardRepoMock,
		userRepo:   userRepoMock,
	}

	h := handlers{
		controller: controllerMock,
	}

	card := models.Card{
		ID:         1,
		UserID:     1,
		BankID:     1,
		Number:     1,
		CreateDate: strfmt.DateTime(time.Now()),
	}

	responseArg := operations.GetCardsIDParams{
		HTTPRequest: &http.Request{},
		ID: card.ID,
	}

	ctx := context.Background()

	tests := []struct {
		name    string
		args    operations.GetCardsIDParams
		prepare func(f *fields)
		wantErr bool
		wantRes middleware.Responder
	}{
		{
			name: "Get Card by ID",
			args: responseArg,
			prepare: func(f *fields) {
				// если указанные вызовы не станут выполняться в ожидаемом порядке, тест будет провален
				gomock.InOrder(
					f.controller.EXPECT().GetCardID(ctx, 1).Return(card, nil),
				)
			},
			wantErr: false,
			// wantRes: middleware.Responder,
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
