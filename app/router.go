package app

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/controller"
	"go-rest-api/exception"
)

func NewRouter(categoryController controller.CategoryController, customerController controller.CustomerController, orderController controller.OrderController, orderproductController controller.OrderProductController, productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/:customerId", customerController.FindById)
	router.POST("/api/customers", customerController.Create)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)

	router.GET("/api/orders", orderController.FindAll)
	router.GET("/api/orders/:ordersId", orderController.FindById)
	router.POST("/api/orders", orderController.Create)
	router.PUT("/api/orders/:orderId", orderController.Update)
	router.DELETE("/api/orders/:orderId", orderController.Delete)

	router.GET("/api/orderproducts", orderproductController.FindAll)
	router.GET("/api/orderproducts/:orderproductId", orderproductController.FindById)
	router.POST("/api/orderproducts", orderproductController.Create)
	router.PUT("/api/orderproducts/:orderproductId", orderproductController.Update)
	router.DELETE("/api/orderproducts/:orderproductId", orderproductController.Delete)

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
