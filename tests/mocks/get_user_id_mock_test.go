package handlers_test

import (
	"card-project/controller"
	"card-project/handlers"
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

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

func Test_GetUserID(t *testing.T) {
	type fields struct {
		rabbitmq *mock.MockRabbitMQ
		cardRepo *mock.MockCardsRepo
		userRepo *mock.MockUsersRepo
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rabbitmqMock := mock.NewMockRabbitMQ(ctrl)
	cardRepoMock := mock.NewMockCardsRepo(ctrl)
	userRepoMock := mock.NewMockUsersRepo(ctrl)

	testFields := &fields{
		rabbitmq: rabbitmqMock,
		cardRepo: cardRepoMock,
		userRepo: userRepoMock,
	}

	service := service.New(userRepoMock, cardRepoMock, rabbitmqMock)
	controller := controller.New(service)
	h := handlers.New(controller, validator.New(validator.WithRequiredStructEnabled()))

	user := models.User{
		ID:        1,
		FirstName: "Ivan",
		LastName:  "Makaroshka",
	}

	reqArgDef := operations.GetUsersIDParams{
		HTTPRequest: &http.Request{},
		ID:          1,
	}

	reqArgErr := operations.GetUsersIDParams{
		HTTPRequest: &http.Request{},
		ID:          111,
	}

	resErr := operations.NewGetUsersIDDefault(404).WithPayload(&models.ErrorResponse{
		Error: &models.ErrorResponseAO0Error{
			Message: "Failed to GET User in storage, user id: " + strconv.FormatInt(reqArgErr.ID, 10) + " " + errors.New("no rows in result set").Error(),
		},
	})

	ctx := context.Background()

	tests := []struct {
		name    string
		args    operations.GetUsersIDParams
		prepare func(f *fields)
		wantRes middleware.Responder
	}{
		{
			name: "valid",
			args: reqArgDef,
			prepare: func(f *fields) {
				// если указанные вызовы не станут выполняться в ожидаемом порядке, тест будет провален
				gomock.InOrder(
					f.userRepo.EXPECT().GetUserID(ctx, 1).Return(user, nil),
				)
			},
			wantRes: operations.NewGetUsersIDOK().WithPayload(&user),
		},
		{
			name: "wrong_ID",
			args: reqArgErr,
			prepare: func(f *fields) {
				gomock.InOrder(
					f.userRepo.EXPECT().GetUserID(ctx, 111).Return(models.User{}, errors.New("no rows in result set")),
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

			response := h.GetUsersID(tt.args)

			if !reflect.DeepEqual(response, tt.wantRes) {
				t.Errorf("GetCard() = %v, want %v", response, tt.wantRes)
			}
		})
	}

}
