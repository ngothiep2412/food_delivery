package rstlikebiz

import (
	"g05-food-delivery/common"
	restaurantlikemodel "g05-food-delivery/module/restaurantlike/model"
	"golang.org/x/net/context"
	"log"
	"time"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type InclikedCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore InclikedCountResStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore, incStore InclikedCountResStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:    store,
		incStore: incStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		panic(restaurantlikemodel.ErrCannotLikeRestaurant(err))
	}

	go func() {
		defer common.AppRecover()
		time.Sleep(5 * time.Second)
		if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
