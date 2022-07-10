package restaurantstorage

import (
	"context"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateById(ctx context.Context, id *restaurantmodel.RestaurantId, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where(id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
