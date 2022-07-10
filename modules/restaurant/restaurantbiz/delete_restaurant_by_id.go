package restaurantbiz

import (
	"context"
	"errors"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantByIdStore interface {
	DeleteRestaurantById(
		ctx context.Context,
		id *restaurantmodel.RestaurantId) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]any,
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantByIdStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantByIdStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz deleteRestaurantBiz) DeleteRestaurantById(ctx context.Context, id *restaurantmodel.RestaurantId) error {
	oldData, err := biz.store.FindDataByCondition(ctx, common.JS{
		"id": id.Id,
	})

	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data not found or deleted")
	}
	//oldData:=queryData[0]

	if err := id.Validate(); err != nil {
		return err
	}

	err = biz.store.DeleteRestaurantById(ctx, id)
	return err
}
