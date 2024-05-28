package api

import "EatRecmd/model"

// / GetEats call eats api for restaurants. More intricate in real life.
func GetEats(requestBody string) model.Recommand {
	res := model.Recommand{
		Id:    "a new uuid",
		Items: make([]model.EatsRestaurantResults, 0),
	}
	return res
}
