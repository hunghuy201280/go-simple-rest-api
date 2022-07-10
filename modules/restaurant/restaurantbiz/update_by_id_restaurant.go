package restaurantbiz

import (
	"context"
	"errors"
	"fmt"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type UpdateByIdRestaurantStore interface {
	UpdateById(ctx context.Context, id *restaurantmodel.RestaurantId, data *restaurantmodel.RestaurantUpdate) error
}

type updateByIdBiz struct {
	store UpdateByIdRestaurantStore
}

func NewUpdateByIdBiz(store UpdateByIdRestaurantStore) *updateByIdBiz {
	return &updateByIdBiz{store: store}
}

func (biz updateByIdBiz) UpdateByIdRestaurant(ctx context.Context, id *restaurantmodel.RestaurantId, data *restaurantmodel.RestaurantUpdate) error {
	if id == nil || id.Id <= 0 {
		return errors.New(fmt.Sprintf("invalid id %d", id))
	}
	if err := biz.store.UpdateById(ctx, id, data); err != nil {
		return err
	}
	return nil
}
