package restaurantbiz

import (
	"context"
	"g05-food-delivery/common"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
	"go.opencensus.io/trace"
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

	ctx1, span := trace.StartSpan(ctx, "restaurant.ListRestaurant")

	span.AddAttributes(
		trace.Int64Attribute("page", int64(paging.Page)),
		trace.Int64Attribute("limit", int64(paging.Limit)),
	)

	result, err := biz.repo.ListRestaurant(ctx1, filter, paging)

	span.End()
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
