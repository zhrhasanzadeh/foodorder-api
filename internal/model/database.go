package model

type Database interface {
	//InsertOrder this is a function for insert new order to database
	InsertOrder(foodOrder FoodOrder) (err error)
	//GetOrder this a function for get order with id
	GetOrder(id int) (foodOrder FoodOrder, err error)
}
