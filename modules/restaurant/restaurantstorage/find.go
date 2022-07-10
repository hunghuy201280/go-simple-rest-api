package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]any,
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for _, key := range moreKeys {
		db = db.Preload(key)
	}
	if err := db.Where(conditions).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, common.RecordNotFound)
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil

}
