package handlers_test

import (
	"card-project/controller"
	"card-project/handlers"
	"card-project/logger"
	mock "card-project/mocks"
	"card-project/models"
	"card-project/restapi/operations"
	"card-project/service"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

func Test_PostUserID(t *testing.T) {
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

	reqArgDef := operations.PostUsersParams{
		HTTPRequest: &http.Request{},

		User: &models.NewUser{
			FirstName: "Ivan",
			LastName:  "Makaroshka",
		},
	}

	user := models.User{
		ID:        1,
		FirstName: "Ivan",
		LastName:  "Makaroshka",
	}

	// reqArgErr := operations.PostUsersParams{
	// 	HTTPRequest: &http.Request{},

	// 	User: &models.NewUser{
		// FirstName: "Ivan",
		// LastName:  "Makaroshka",
	// },
	// }

	// resErr := operations.NewGetCardsIDDefault(500).WithPayload(&models.ErrorResponse{
	// 	Error: &models.ErrorResponseAO0Error{
	// 		Message: "Failed to POST Card in storage " + errors.New("").Error(),
	// 	},
	// })

	tests := []struct {
		name    string
		args    operations.PostUsersParams
		prepare func(f *fields)
		wantRes middleware.Responder
	}{
		{
			name: "valid",
			args: reqArgDef,
			prepare: func(f *fields) {
				// если указанные вызовы не станут выполняться в ожидаемом порядке, тест будет провален
				gomock.InOrder(
					f.rabbitmq.EXPECT().ProducePostUser(gomock.Any(), gomock.Any()).Return(nil),
				)
			},
			wantRes: operations.NewPostUsersCreated().WithPayload(&user),
		},
		// {
		// 	name: "wrong_ID",
		// 	args: reqArgErr,
		// 	prepare: func(f *fields) {
		// 		gomock.InOrder(
		// 			f.rabbitmq.EXPECT().ProducePostCard(gomock.Any(), reqArgErr).Return(errors.New("no rows in result set")),
		// 		)
		// 	},
		// 	wantRes: resErr,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare(testFields)
			}

			response := h.PostUsers(tt.args)

			//! тест работает, но response := h.PostCards(tt.args) генерит свой id, а tt.wantRes имеет изначально заданный, я незнаю как из respons вытащить id и вставить его в tt.wantRes
			if !reflect.DeepEqual(response, tt.wantRes) {
				t.Errorf("PostUser() = %v, want %v", response, tt.wantRes)
			}
		})
	}

}
