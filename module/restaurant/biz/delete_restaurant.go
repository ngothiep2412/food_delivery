package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id int) error
}

type DeleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurant(store DeleteRestaurantStore) *DeleteRestaurantBiz {
	return &DeleteRestaurantBiz{store: store}
}

func (biz *DeleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid restaurant id")
	}

	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("restaurant not found")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
