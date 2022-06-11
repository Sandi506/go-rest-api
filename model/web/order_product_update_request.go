package web

type OrderProductUpdateRequest struct {
	Id        int `validate:"required"`
	OrderId   int `validate:"required"`
	ProductId int `validate:"required"`
	Qty       int `validate:"required"`
	Price     int `validate:"required"`
	Amount    int `validate:"required"`
}
