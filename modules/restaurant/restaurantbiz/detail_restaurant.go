package restaurantbiz

import (
	"context"
	"errors"
	"simple_golang/modules/restaurant/restaurantmodel"
)

type DetailRestaurantStore interface {
	DetailData(
		ctx context.Context,
		id string,
	) (restaurantmodel.Restaurant, error)
}

type detailRestaurantBiz struct {
	store DetailRestaurantStore
}

func NewDetailRestaurantBiz(store DetailRestaurantStore) *detailRestaurantBiz {
	return &detailRestaurantBiz{store: store}
}

func (biz *detailRestaurantBiz) DetailRestaurant(
	ctx context.Context,
	id string,
) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.DetailData(ctx, id)

	if err != nil {
		return nil, err
	}

	if result.Status == 0 {
		return nil, errors.New("data deleted")
	}

	return &result, err
}
