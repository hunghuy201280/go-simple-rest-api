package uploadstorage

import (
	"context"
	"simple-rest-api/common"
)

func (store *sqlStore) CreateImage(ctx context.Context, data *common.Image) error {
	db := store.db
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
