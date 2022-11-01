package utils

import (
	"api-orderfood/internal/model"
	"encoding/json"
)

func JsonToStruct(j string) (model.FoodOrder, error) {
	var foodOrder model.FoodOrder
	err := json.Unmarshal([]byte(j), &foodOrder)
	if err != nil {
		return model.FoodOrder{}, err
	}
	return foodOrder, nil
}

func StructToJson(f model.FoodOrder) (string, error) {
	foodOrder, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(foodOrder), nil
}
