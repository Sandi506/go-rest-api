package web

type ProductUpdateRequest struct {
	Id         int    `validate:"required"`
	Name       string `validate:"required"`
	Price      int    `validate:"required"`
	CategoryId int    `validate:"required"`
}
