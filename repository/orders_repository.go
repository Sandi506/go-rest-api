package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Update(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Delete(ctx context.Context, tx *sql.Tx, order domain.Order)
	FindById(ctx context.Context, tx *sql.Tx, orderId int) (domain.Order, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Order
}
