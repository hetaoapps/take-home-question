package model

import (
	"time"
)

// / QueryRecord is the model of query that could be saved into mongoDB
type QueryRecord struct {
	UserId      int64
	Time        time.Time
	QueryString string
	ResultId    string
	Location    string
}

type GPTReturnBody struct {
	FoodType       string
	RestaurantName string
	FoodName       string
}

type Recommand struct {
	Id    string
	Items []EatsRestaurantResults
}

type EatsRestaurantResults struct {
	RestaurantName string
	FoodType       string
	Rating         string
}

// Following are the return body per the Github readme
// MenuItem represents a menu item in a store
type MenuItem struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Price string `json:"price"`
}

// Store represents a store with a list of menu items
type Store struct {
	Name     string     `json:"name"`
	ID       string     `json:"id"`
	MenuItem []MenuItem `json:"menuItem"`
}

// Recommendation represents a recommendation with a store
type Recommendation struct {
	RecommendationID string `json:"recommendationId"`
	Store            Store  `json:"store"`
}
