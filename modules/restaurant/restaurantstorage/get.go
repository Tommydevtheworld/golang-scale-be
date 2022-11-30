package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
	"simple_golang/common"
	"simple_golang/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DetailData(
	ctx context.Context,
	id string,
) (*restaurantmodel.Restaurant, error) {
	var result = restaurantmodel.Restaurant{}

	db := s.db
	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id =c ?", id).
		//Where("status IN (1)").
		Take(&result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &result, nil
}

func (s *sqlStore) FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var restaurant = restaurantmodel.Restaurant{}

	db := s.db
	if err := db.Model(restaurantmodel.Restaurant{}).Where(conditions).
		Where("status in (1)").
		First(&restaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, err
	}
	return &restaurant, nil
}
