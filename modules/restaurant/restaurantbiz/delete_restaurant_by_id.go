package restaurantbiz

import (
	"context"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantByIdStore interface {
	DeleteRestaurantById(
		ctx context.Context,
		id *restaurantmodel.RestaurantId) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantByIdStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantByIdStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz deleteRestaurantBiz) DeleteRestaurantById(ctx context.Context, id *restaurantmodel.RestaurantId) error {
	if err := id.Validate(); err != nil {
		return err
	}

	err := biz.store.DeleteRestaurantById(ctx, id)
	return err
}
