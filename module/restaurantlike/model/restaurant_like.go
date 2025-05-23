package restaurantlikemodel

import (
	"g05-food-delivery/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int                `json:"user_id" gorm:"column:user_id"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

type UnLike struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int `json:"user_id" gorm:"column:user_id"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}
func (l *Like) GetUserId() int {
	return l.UserId
}

func (ul *UnLike) GetRestaurantId() int {
	return ul.RestaurantId
}
func (ul *UnLike) GetUserId() int {
	return ul.UserId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(err,
		"cannot like this restaurant",
		"ErrCannotLikeRestaurant",
	)
}

func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(err,
		"cannot unlike this restaurant",
		"ErrCannotUnlikeRestaurant",
	)
}
