package restaurantmodel

import (
	"errors"
	"fmt"
	"simple-rest-api/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (r Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name   *string        `json:"name" gorm:"column:name;"`
	Addr   *string        `json:"address" gorm:"column:addr;"`
	Status *int           `json:"status" gorm:"column:status;"`
	Logo   *common.Image  `json:"logo" gorm:"column:logo"`
	Cover  *common.Images `json:"cover" gorm:"column:cover"`
}

func (r RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id    int            `json:"id" gorm:"column:id;"`
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo"`
	Cover *common.Images `json:"cover" gorm:"column:cover"`
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

func (id RestaurantId) TableName() string {
	return Restaurant{}.TableName()
}

func (id *RestaurantId) Validate() error {
	if id == nil || id.Id <= 0 {
		return errors.New(fmt.Sprintf("invalid Id %d", id.Id))
	}
	return nil
}
