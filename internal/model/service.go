package model

type Service interface {
	AddOrder(order FoodOrder) error
	GetUser(id int) (order FoodOrder, err error)
}
