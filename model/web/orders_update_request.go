package web

type OrderUpdateRequest struct {
	Id          int `validate:"required"`
	CustomerId  int `validate:"required" json:"customer-id"`
	TotalAmount int `validate:"required" json:"total-amount"`
}
