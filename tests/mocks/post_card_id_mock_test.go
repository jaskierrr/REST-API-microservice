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
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

func Test_PostCardID(t *testing.T) {
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

	service := service.New(userRepoMock, cardRepoMock, rabbitmqMock)
	controller := controller.New(service, logger)
	h := handlers.New(controller, validator.New(validator.WithRequiredStructEnabled()), logger)

	reqArgDef := operations.PostCardsParams{
		HTTPRequest: &http.Request{},

		Card: &models.NewCard{
			UserID:     1,
			BankID:     1,
			Number:     1,
			CreateDate: strfmt.DateTime(time.Now()),
		},
	}

	card := models.Card{
		ID:         1,
		UserID:     1,
		BankID:     1,
		Number:     1,
		CreateDate: strfmt.DateTime(time.Now()),
	}

	// reqArgErr := operations.PostCardsParams{
	// 	HTTPRequest: &http.Request{},

	// 	Card: &models.NewCard{
	// 	UserID:     1,
	// 	BankID:     1,
	// 	Number:     1,
	// 	CreateDate: strfmt.DateTime(time.Now()),
	// 	},
	// }

	// resErr := operations.NewGetCardsIDDefault(500).WithPayload(&models.ErrorResponse{
	// 	Error: &models.ErrorResponseAO0Error{
	// 		Message: "Failed to POST Card in storage " + errors.New("").Error(),
	// 	},
	// })


	tests := []struct {
		name    string
		args    operations.PostCardsParams
		prepare func(f *fields)
		wantRes middleware.Responder
	}{
		{
			name: "valid",
			args: reqArgDef,
			prepare: func(f *fields) {
				// если указанные вызовы не станут выполняться в ожидаемом порядке, тест будет провален
				gomock.InOrder(
					f.rabbitmq.EXPECT().ProducePostCard(gomock.Any(), gomock.Any()).Return(nil),
				)
			},
			wantRes: operations.NewPostCardsCreated().WithPayload(&card),
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

			response := h.PostCards(tt.args)

		//! тест работает, но response := h.PostCards(tt.args) генерит свой id, а tt.wantRes имеет изначально заданный, я незнаю как из respons вытащить id и вставить его в tt.wantRes
			if !reflect.DeepEqual(response, tt.wantRes) {
				t.Errorf("PostCard() = %v, want %v", response, tt.wantRes)
			}
		})
	}

}
