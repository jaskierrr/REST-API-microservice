package handlers_test

import (
	"card-project/controller"
	"card-project/handlers"
	mock "card-project/mocks"
	"card-project/models"
	"card-project/restapi/operations"
	"card-project/service"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

func Test_DeleteUserID(t *testing.T) {
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

	reqArgDef := operations.DeleteUsersIDParams{
		HTTPRequest: &http.Request{},
		ID:          1,
	}

	reqArgErr := operations.DeleteUsersIDParams{
		HTTPRequest: &http.Request{},
		ID:          111,
	}

	resErr := operations.NewDeleteUsersIDDefault(500).WithPayload(&models.ErrorResponse{
		Error: &models.ErrorResponseAO0Error{
			Message: "Failed to DELETE User in storage, user id: " + strconv.FormatInt(reqArgErr.ID, 10) + " " + errors.New("no rows in result set").Error(),
		},
	})


	tests := []struct {
		name    string
		args    operations.DeleteUsersIDParams
		prepare func(f *fields)
		wantRes middleware.Responder
	}{
		{
			name: "valid",
			args: reqArgDef,
			prepare: func(f *fields) {
				// если указанные вызовы не станут выполняться в ожидаемом порядке, тест будет провален
				gomock.InOrder(
					f.rabbitmq.EXPECT().ProduceDeleteUser(gomock.Any(), 1).Return(nil),
				)
			},
			wantRes: operations.NewDeleteUsersIDNoContent(),
		},
		{
			//? я не понял что я могу тут проверить, если ошибка у меня рождается только если рэббиту плохо, но я так понимаю у меня в хендлере уже стоит обработка ошибки на неверный id, поэтому этот тест проходит, так как я насильно ее получаю в EXPECT?
			name: "wrong_ID",
			args: reqArgErr,
			prepare: func(f *fields) {
				gomock.InOrder(
					f.rabbitmq.EXPECT().ProduceDeleteUser(gomock.Any(), 111).Return(errors.New("no rows in result set")),
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

			response := h.DeleteUsersID(tt.args)

			if !reflect.DeepEqual(response, tt.wantRes) {
				t.Errorf("DeleteUser() = %v, want %v", response, tt.wantRes)
			}
		})
	}

}
