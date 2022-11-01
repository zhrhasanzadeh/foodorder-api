package model

type FoodOrder struct {
	OrderID int    `json:"id"`
	Price   int    `json:"price"`
	Title   string `json:"title"`
}
