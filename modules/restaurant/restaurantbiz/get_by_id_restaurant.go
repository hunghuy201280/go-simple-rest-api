package restaurantbiz

import (
	"context"
	"simple-rest-api/common"
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

	if err := data.Validate(); err != nil {
		return nil, err
	}

	result, err := biz.store.GetById(ctx, data)

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}
	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return result, err

}
