package restaurantlikestorage

import (
	"context"
	"g05-food-delivery/common"
	restaurantlikemodel "g05-food-delivery/module/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, data *restaurantlikemodel.UnLike) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", data.UserId, data.RestaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil

}
