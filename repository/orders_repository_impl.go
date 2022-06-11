package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "insert into orders (customer_id, total_amount) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, order.CustomerId, order.TotalAmount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.Id = int(id)
	return order
}

func (repository *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "update orders set customer_id = ?, total_amount = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, order.CustomerId, order.TotalAmount, order.Id)
	helper.PanicIfError(err)

	return order
}

func (repository *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, order domain.Order) {
	SQL := "delete from orders where id = ?"
	_, err := tx.ExecContext(ctx, SQL, order.Id)
	helper.PanicIfError(err)
}

func (repository *OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderId int) (domain.Order, error) {
	logrus.Info("Order Repository Find By Id Start")
	SQL := "select id, customer_id, total_amount from orders where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, orderId)
	helper.PanicIfError(err)
	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.CustomerId, &order.TotalAmount)
		helper.PanicIfError(err)
		logrus.Info("Order Repository Find By Id End")
		return order, nil
	} else {
		logrus.Info("Order Repository Find By Id End")
		return order, errors.New("order is not found")
	}
}

func (repository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Order {
	logrus.Info("Order Repository Find All Start")
	SQL := "select id, customer_id, total_amount from orders"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(&order.Id, &order.CustomerId, &order.TotalAmount)
		helper.PanicIfError(err)
		orders = append(orders, order)
	}
	logrus.Info("Order Repository Find All End")
	return orders
}
