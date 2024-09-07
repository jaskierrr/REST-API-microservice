package handlers

import (
	"card-project/restapi/operations"
	"fmt"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetUsersID(params operations.GetUsersIDParams) middleware.Responder {

	fmt.Println(params.ID)
	user, err := h.controller.GetUserID(params.ID)
	if err != nil {
		log.Fatal("Failed to get User in storage: ", err)
	}

	fmt.Println(user)

	// h.controller.GetUsers() -> внутри делаешь скл
	//if err != nil {
	// поменять KPIGet на свои методы
	//	return swaggerapi.NewKPIGetDefault(baseErrCode).
	//		WithPayload(&swaggerapi.KPIGetDefaultBody{
	//			Error: &swaggerapi.KPIGetDefaultBodyKPIGetDefaultBodyAO0Error{
	//				Message: err.Error(),
	//			},
	//		})
	//
	//return swaggerapi.NewKPIGetOK().WithPayload(&swaggerapi.KPIGetOKBody{})
	return nil
}
