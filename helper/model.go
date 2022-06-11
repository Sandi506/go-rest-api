package helper

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

// category
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

// customer
func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		Address:     customer.Address,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
	}
}

func ToCustomerResponses(customers []domain.Customer) []web.CustomerResponse {
	var customerResponses []web.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}

// order
func ToOrderResponse(orders domain.Order) web.OrderResponse {
	return web.OrderResponse{
		Id:          orders.Id,
		CustomerId:  orders.CustomerId,
		TotalAmount: orders.TotalAmount,
	}
}

func ToOrderResponses(orders []domain.Order) []web.OrderResponse {
	var orderResponses []web.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, ToOrderResponse(order))
	}
	return orderResponses
}

// order product
func ToOrderProductResponse(orderproducts domain.OrderProduct) web.OrderProductResponse {
	return web.OrderProductResponse{
		Id:        orderproducts.Id,
		OrderId:   orderproducts.OrderId,
		ProductId: orderproducts.ProductId,
		Qty:       orderproducts.Qty,
		Price:     orderproducts.Price,
		Amount:    orderproducts.Amount,
	}
}

func ToOrderProductResponses(orderproducts []domain.OrderProduct) []web.OrderProductResponse {
	var orderproductResponses []web.OrderProductResponse
	for _, orderproduct := range orderproducts {
		orderproductResponses = append(orderproductResponses, ToOrderProductResponse(orderproduct))
	}
	return orderproductResponses
}

// product
func ToProductResponse(products domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:         products.Id,
		Name:       products.Name,
		Price:      products.Price,
		CategoryId: products.CategoryId,
	}
}

func ToProduct(products web.ProductResponse) domain.Product {
	return domain.Product{
		Id:         products.Id,
		Name:       products.Name,
		Price:      products.Price,
		CategoryId: products.CategoryId,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
