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

type OrderProductControllerImpl struct {
	OrderProductService service.OrderProductService
}

func NewOrderProductController(orderproductService service.OrderProductService) OrderProductController {
	return &OrderProductControllerImpl{
		OrderProductService: orderproductService,
	}
}

func (controller OrderProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Product Controller Create Start")
	orderproductCreateRequest := web.OrderProductCreateRequest{}
	helper.ReadFromRequestBody(request, &orderproductCreateRequest)

	orderproductResponse := controller.OrderProductService.Create(request.Context(), orderproductCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderproductResponse,
	}

	logrus.Info("Order Product Controller Create End")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller OrderProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Product Controller Update Start")
	orderproductUpdateRequest := web.OrderProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &orderproductUpdateRequest)

	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderproductUpdateRequest.Id = id

	orderproductResponse := controller.OrderProductService.Update(request.Context(), orderproductUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderproductResponse,
	}

	logrus.Info("Order Product Controller Update End")
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller OrderProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Product Controller Delete Start")
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	controller.OrderProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	logrus.Info("Order Product Controller Update End")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller OrderProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Product Controller Find By Id Start")
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderproductResponse := controller.OrderProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderproductResponse,
	}

	logrus.Info("Order Product Controller Find By Id End")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller OrderProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Order Product Controller Find By All Start")
	orderproductResponses := controller.OrderProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderproductResponses,
	}

	logrus.Info("Order Product Controller Find All End")
	helper.WriteToResponseBody(writer, webResponse)
}
