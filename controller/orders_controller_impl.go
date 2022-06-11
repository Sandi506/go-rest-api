package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"go-rest-api/service"
	"net/http"
	"strconv"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Controller Create Start")
	orderCreateRequest := web.OrderCreateRequest{}
	helper.ReadFromRequestBody(request, &orderCreateRequest)

	orderResponse := controller.OrderService.Create(request.Context(), orderCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	logrus.Info("Order Controller Create End")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Controller Update Start")
	orderUpdateRequest := web.OrderUpdateRequest{}
	helper.ReadFromRequestBody(request, &orderUpdateRequest)

	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderUpdateRequest.Id = id

	orderResponse := controller.OrderService.Update(request.Context(), orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	logrus.Info("Order Controller Update End")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Controller Delete Start")
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	controller.OrderService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	logrus.Info("Order Controller Delete End")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Controller Find By Id Start")
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	logrus.Info("Order Controller Find By Id End")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Controller Find All Start")
	orderResponses := controller.OrderService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponses,
	}

	logrus.Info("Order Controller Find All End")
	helper.WriteToResponseBody(writer, webResponse)
}
