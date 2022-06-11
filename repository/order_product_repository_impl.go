package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type OrderProductRepositoryImpl struct {
}

func NewOrderProductRepository() OrderProductRepository {
	return &OrderProductRepositoryImpl{}

}
func (repository *OrderProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct) domain.OrderProduct {
	SQL := "insert into order_product(order_id, product_id, qty,price, amount) value (?,?,?,?,?)" // sesuai format DB
	result, err := tx.ExecContext(ctx, SQL, orderproduct.OrderId, orderproduct.ProductId, orderproduct.Qty, orderproduct.Price, orderproduct.Amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orderproduct.Id = int(id)
	return orderproduct
}

func (repository *OrderProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct) domain.OrderProduct {
	SQL := "update order_product set qty = ?, price = ?, amount = ? where id =?"
	_, err := tx.ExecContext(ctx, SQL, orderproduct.OrderId, orderproduct.ProductId, orderproduct.Qty, orderproduct.Price, orderproduct.Amount, orderproduct.Id)
	helper.PanicIfError(err)
	return orderproduct
}
func (repository *OrderProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct) {
	SQL := "delete from order_product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orderproduct.Id)
	helper.PanicIfError(err)
}

func (repository *OrderProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, OrderProductId int) (domain.OrderProduct, error) {
	logrus.Info("Order Product Repository Find By Id Start")
	SQL := "select id, order_id, product_id, qty, price,amount from order_product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, OrderProductId)
	helper.PanicIfError(err)
	defer rows.Close()

	orderproduct := domain.OrderProduct{}
	if rows.Next() {
		err := rows.Scan(&orderproduct.Id, &orderproduct.OrderId, &orderproduct.ProductId, &orderproduct.Qty, &orderproduct.Price, &orderproduct.Amount)
		helper.PanicIfError(err)
		logrus.Info("Order Product Repository Find By Id End")
		return orderproduct, nil
	} else {
		logrus.Info("Order Product Repository Find By Id End")
		return orderproduct, errors.New("customer is not found")
	}
}

func (repository *OrderProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.OrderProduct {
	logrus.Info("Order Product Repository Find All Start")
	SQL := "select id, order_id, product_id, qty, price,amount from order_product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orderproducts []domain.OrderProduct
	for rows.Next() {
		orderproduct := domain.OrderProduct{}
		err := rows.Scan(&orderproduct.Id, &orderproduct.OrderId, &orderproduct.ProductId, &orderproduct.Qty, &orderproduct.Price, &orderproduct.Amount)
		helper.PanicIfError(err)
		orderproducts = append(orderproducts, orderproduct)
	}
	logrus.Info("Order Product Repository Find All End")
	return orderproducts
}
