package rstlikebiz

import (
	"g05-food-delivery/component/asyncjob"
	restaurantlikemodel "g05-food-delivery/module/restaurantlike/model"
	"golang.org/x/net/context"
	"log"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.UnLike) error
}

type UserDecreaseRestaurantStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decStore UserDecreaseRestaurantStore
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, decStore UserDecreaseRestaurantStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store:    store,
		decStore: decStore,
	}
}

// Side effect

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.UnLike,
) error {
	err := biz.store.Delete(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, j); err != nil {
		log.Println(err)
	}

	return nil

}
