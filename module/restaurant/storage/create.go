package restaurantstorage

import (
	"context"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
