package restaurantstorage

import (
	"context"
	"simple_golang/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DetailData(
	ctx context.Context,
	id string,
) (restaurantmodel.Restaurant, error) {
	var result = restaurantmodel.Restaurant{}

	db := s.db
	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Take(&result).
		Error; err != nil {
		return result, err
	}
	return result, nil
}
