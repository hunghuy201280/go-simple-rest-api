package restaurantbiz

import (
	"context"
	"errors"
	"fmt"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type GetByIdRestaurantStore interface {
	GetById(
		ctx context.Context,
		data *restaurantmodel.RestaurantId) (
		*restaurantmodel.Restaurant,
		error)
}

type getByIdRestaurantBiz struct {
	store GetByIdRestaurantStore
}

func NewGetByIdRestaurantBiz(store GetByIdRestaurantStore) *getByIdRestaurantBiz {
	return &getByIdRestaurantBiz{store: store}
}

func (biz getByIdRestaurantBiz) GetByIdRestaurant(
	ctx context.Context,
	data *restaurantmodel.RestaurantId) (
	*restaurantmodel.Restaurant,
	error) {

	if data == nil || data.Id <= 0 {
		return nil, errors.New(fmt.Sprintf("invalid Id %d", data.Id))
	}

	result, err := biz.store.GetById(ctx, data)

	return result, err

}
