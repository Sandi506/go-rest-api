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

type OrderProductServiceImpl struct {
	OrderProductRepository repository.OrderProductRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewOrderProductService(orderproductRepository repository.OrderProductRepository, DB *sql.DB, validate *validator.Validate) OrderProductService {
	return &OrderProductServiceImpl{
		OrderProductRepository: orderproductRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *OrderProductServiceImpl) Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderproduct := domain.OrderProduct{
		OrderId:   request.OrderId,
		ProductId: request.ProductId,
	}

	orderproduct = service.OrderProductRepository.Save(ctx, tx, orderproduct)

	return helper.ToOrderProductResponse(orderproduct)
}

func (service *OrderProductServiceImpl) Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderproduct, err := service.OrderProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	orderproduct.OrderId = request.OrderId
	orderproduct.ProductId = request.ProductId

	orderproduct = service.OrderProductRepository.Update(ctx, tx, orderproduct)

	return helper.ToOrderProductResponse(orderproduct)
}

func (service *OrderProductServiceImpl) Delete(ctx context.Context, orderId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderproduct, err := service.OrderProductRepository.FindById(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrderProductRepository.Delete(ctx, tx, orderproduct)
}

func (service *OrderProductServiceImpl) FindById(ctx context.Context, orderId int) web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderproduct, err := service.OrderProductRepository.FindById(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderProductResponse(orderproduct)
}

func (service *OrderProductServiceImpl) FindAll(ctx context.Context) []web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderproducts := service.OrderProductRepository.FindAll(ctx, tx)

	return helper.ToOrderProductResponses(orderproducts)
}
