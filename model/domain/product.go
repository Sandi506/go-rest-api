package domain

import "time"

type Product struct {
	Id         int
	Name       string
	Price      int
	CategoryId int
	//CategoryName string
	CreateAt time.Time
	UpdateAt time.Time
}
