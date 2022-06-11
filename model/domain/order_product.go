package domain

import "time"

type OrderProduct struct {
	Id        int
	OrderId   int
	ProductId int
	Qty       int
	Price     int
	Amount    int
	CreateAt  time.Time
	UpdateAt  time.Time
}
