package web

type OrderProductCreateRequest struct {
	OrderId   int `validate:"required,min=1,max=100" json:"order-id"`
	ProductId int `validate:"required,min=1,max=100" json:"product-id"`
	Qty       int `validate:"required,min=1,max=100" json:"qty"`
	Price     int `validate:"required,min=1,max=100" json:"price"`
	Amount    int `validate:"required,min=1,max=100" json:"amount"`
}
