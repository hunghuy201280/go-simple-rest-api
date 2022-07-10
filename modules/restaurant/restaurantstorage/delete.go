package restaurantstorage

import (
	"context"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DeleteRestaurantById(
	ctx context.Context,
	id *restaurantmodel.RestaurantId) error {
	db := s.db

	deletedStatus := 0
	if err := db.Table(id.TableName()).Where(&id).Updates(restaurantmodel.RestaurantUpdate{
		Status: &deletedStatus,
	}).Error; err != nil {
		return err
	}
	return nil
}
