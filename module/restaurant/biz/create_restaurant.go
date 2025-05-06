package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface { //interface ->  khai báo ở nơi dùng nó
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurant(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name is required")
	}

	if err := biz.store.CreateRestaurant(ctx, data); err != nil {
		return err
	}

	return nil
}
