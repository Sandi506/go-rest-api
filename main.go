package main

import (
	"github.com/go-playground/validator/v10"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/repository"
	"go-rest-api/service"
	"net/http"
)

func main() {

	//category
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Customer
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	//Order
	orderRepository := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepository, db, validate)
	orderController := controller.NewOrderController(orderService)

	orderproductRepository := repository.NewOrderProductRepository()
	orderproductService := service.NewOrderProductService(orderproductRepository, db, validate)
	orderproductController := controller.NewOrderProductController(orderproductService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := app.NewRouter(categoryController, customerController, orderController, orderproductController, productController)

	server := http.Server{
		Addr:    "https://booking-system-management.herokuapp.com/",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
