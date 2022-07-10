package restaurantmodel

import (
	"errors"
	"strings"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (r Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (r RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id int `json:"id" gorm:"column:id;"`

	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (r RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (r RestaurantCreate) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	if len(r.Name) == 0 {
		return errors.New("restaurant name could not be empty")
	}
	return nil
}

type RestaurantId struct {
	Id int `json:"id" gorm:"column:id;" uri:"id"`
}

func (r RestaurantId) TableName() string {
	return Restaurant{}.TableName()
}
