package restaurantstorage

import (
	"context"
	"errors"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) GetById(
	ctx context.Context,
	data *restaurantmodel.RestaurantId) (
	*restaurantmodel.Restaurant,
	error) {
	db := *s.db
	var result []restaurantmodel.Restaurant

	if err := db.Where(&data).Find(&result).Error; err != nil {
		return nil, err
	}
	if result == nil || len(result) == 0 {
		return nil, errors.New("not found")
	}

	return &result[0], nil

}
