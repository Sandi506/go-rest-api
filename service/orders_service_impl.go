package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
	"go-rest-api/repository"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := domain.Order{
		CustomerId: request.CustomerID,
	}

	order = service.OrderRepository.Save(ctx, tx, order)

	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) Update(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	order.CustomerId = request.Id
	order.TotalAmount = request.TotalAmount

	order = service.OrderRepository.Update(ctx, tx, order)

	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) Delete(ctx context.Context, orderId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindById(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrderRepository.Delete(ctx, tx, order)
}

func (service *OrderServiceImpl) FindById(ctx context.Context, orderId int) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindById(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) FindAll(ctx context.Context) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := service.OrderRepository.FindAll(ctx, tx)

	return helper.ToOrderResponses(order)
}
