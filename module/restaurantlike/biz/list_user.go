package rstlikebiz

import (
	"g05-food-delivery/common"
	restaurantlikemodel "g05-food-delivery/module/restaurantlike/model"
	"golang.org/x/net/context"
)

type ListUserLikeRestaurantStore interface {
	GetUsersLikeRestaurant(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store: store}
}

func (biz *listUserLikeRestaurantBiz) ListUsers(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
