package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) GetById(
	ctx context.Context,
	data *restaurantmodel.RestaurantId) (
	*restaurantmodel.Restaurant,
	error) {
	db := *s.db
	var result restaurantmodel.Restaurant

	if err := db.Where(&data).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, common.RecordNotFound)
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil

}
