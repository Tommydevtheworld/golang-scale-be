package restaurantbiz

import (
	"context"
	"simple_golang/common"
	"simple_golang/modules/restaurant/restaurantmodel"
)

type DetailRestaurantStore interface {
	DetailData(
		ctx context.Context,
		id string,
	) (*restaurantmodel.Restaurant, error)
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
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return result, err
}
