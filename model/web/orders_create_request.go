package web

type OrderCreateRequest struct {
	CustomerID  int `validate:"required" json:"customer-id"`
	TotalAmount int `validate:"required" json:"total-amount"`
}
