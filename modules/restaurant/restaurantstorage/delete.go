package restaurantstorage

import (
	"context"
	"simple_golang/common"
	"simple_golang/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.
		TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
