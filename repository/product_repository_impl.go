package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(name,price, category_id) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "update product set name = ?, price = ?, cataegory_id = ?, where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (web.ProductResponse, error) {
	logrus.Info("Product Repository Find By Id Start")
	SQL := "select p.id, p.name, p.price, p.category_id, c.name FROM product p INNER JOIN category c on p.category_id = c.id where p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := web.ProductResponse{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		logrus.Info("Product Repository Find By Id End")
		return product, nil
	} else {
		logrus.Info("Product Repository Find By Id End")
		return product, errors.New("product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []web.ProductResponse {
	logrus.Info("Product Repository Find All Start")
	SQL := "SELECT p.id, p.name, p.price, p.category_id, c.name FROM product p INNER JOIN category c on p.category_id = c.id"
	//SQL := "select id, name,price,category_id from product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []web.ProductResponse
	for rows.Next() {
		product := web.ProductResponse{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	logrus.Info("Product Repository Find All End")
	return products
}
