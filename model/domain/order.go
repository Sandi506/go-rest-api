package domain

import "time"

type Order struct {
	Id          int
	OrderDate   time.Time
	CustomerId  int
	TotalAmount int
	CreateAt    time.Time
	UpdateAt    time.Time
}
