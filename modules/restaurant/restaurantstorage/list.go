package restaurantstorage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	conditions map[string]any,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db

	for _, key := range moreKeys {
		db = db.Preload(key)
	}
	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id= ?", v.CityId)
		}
	}

	//if err := db.Count(&paging.Total).Error; err != nil {
	//	return nil, err
	//}
	if paging == nil {
		paging = &common.Paging{}
	}
	paging.Fulfill()

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	paging.Total = int64(len(result))
	return result, nil

}
