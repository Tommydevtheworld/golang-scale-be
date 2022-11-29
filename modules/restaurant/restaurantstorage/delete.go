package restaurantstorage

import (
	"context"
	"simple_golang/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Delete(ctx context.Context, id string) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.
		TableName()).
		Delete("id", id); err != nil {

	}
	return nil
}
