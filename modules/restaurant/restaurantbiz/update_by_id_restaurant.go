package restaurantbiz

import (
	"context"
	"errors"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type UpdateByIdRestaurantStore interface {
	UpdateById(
		ctx context.Context,
		id *restaurantmodel.RestaurantId,
		data *restaurantmodel.RestaurantUpdate,
	) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]any,
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type updateRestaurantByIdBiz struct {
	store UpdateByIdRestaurantStore
}

func NewUpdateRestaurantByIdBiz(store UpdateByIdRestaurantStore) *updateRestaurantByIdBiz {
	return &updateRestaurantByIdBiz{store: store}
}

func (biz updateRestaurantByIdBiz) UpdateByIdRestaurant(
	ctx context.Context,
	id *restaurantmodel.RestaurantId,
	data *restaurantmodel.RestaurantUpdate,
) error {

	oldData, err := biz.store.FindDataByCondition(ctx, common.JS{
		"id": id.Id,
	},
	)

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
	if err := biz.store.UpdateById(ctx, id, data); err != nil {
		return err
	}
	return nil
}
