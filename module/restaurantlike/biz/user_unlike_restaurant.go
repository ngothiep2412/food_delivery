package rstlikebiz

import (
	"g05-food-delivery/common"
	restaurantlikemodel "g05-food-delivery/module/restaurantlike/model"
	"g05-food-delivery/pubsub"
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
	store UserUnlikeRestaurantStore
	//decStore UserDecreaseRestaurantStore
	ps pubsub.Pubsub
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore,
	//decStore UserDecreaseRestaurantStore,
	ps pubsub.Pubsub,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store: store,
		//decStore: decStore,
		ps: ps,
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

	if err := biz.ps.Publish(ctx, common.TopicUserUnlikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//if err := asyncjob.NewGroup(true, j); err != nil {
	//	log.Println(err)
	//}

	return nil

}
