package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}

}
func (repository *CustomerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "insert into customer(name, address, email,phone_number) value (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Address, customer.Email, customer.PhoneNumber)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	customer.Id = int(id)
	return customer
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "update customer set name = ?, address = ?, email = ?, phone_number = ? where id =?"
	_, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Address, customer.Email, customer.PhoneNumber, customer.Id)
	helper.PanicIfError(err)
	return customer
}
func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer) {
	SQL := "delete from customer where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Id)
	helper.PanicIfError(err)
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error) {
	logrus.Info("Customer Repository Find By Id Start")
	SQL := "select id, name, address, email, phone_number from customer where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Email, &customer.PhoneNumber)
		helper.PanicIfError(err)
		logrus.Info("Customer Repository Find By Id End")
		return customer, nil
	} else {
		logrus.Info("Cutomer Repository Find By Id End")
		return customer, errors.New("customer is not found")
	}
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	logrus.Info("Customer Repository Find All Start")
	SQL := "select id, name, address, email, phone_number from customer"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Email, &customer.PhoneNumber)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}
	logrus.Info("Customer Repository Find All End")
	return customers
}
