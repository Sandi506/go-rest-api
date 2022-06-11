package service

import (
	"context"
	"go-rest-api/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse
	Update(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse
	Delete(ctx context.Context, orderId int)
	FindById(ctx context.Context, orderId int) web.OrderResponse
	FindAll(ctx context.Context) []web.OrderResponse
}
