package restaurantbiz

import (
	"context"
	"g05-food-delivery/common"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type ListRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurant(repo ListRestaurantRepo) *ListRestaurantBiz {
	return &ListRestaurantBiz{repo: repo}
}

func (biz *ListRestaurantBiz) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.repo.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
