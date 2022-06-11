package service

import (
	"context"
	"go-rest-api/model/web"
)

type OrderProductService interface {
	Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse
	Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse
	Delete(ctx context.Context, orderId int)
	FindById(ctx context.Context, orderId int) web.OrderProductResponse
	FindAll(ctx context.Context) []web.OrderProductResponse
}
