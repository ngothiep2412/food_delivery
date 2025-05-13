package restaurantrepo

import (
	"g05-food-delivery/common"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
	"golang.org/x/net/context"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

//type LikeRestaurantRepo interface {
//	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
//}

type ListRestaurantBiz struct {
	store ListRestaurantStore
	//listStore LikeRestaurantRepo
}

func NewListRestaurantRepo(store ListRestaurantStore) *ListRestaurantBiz {
	return &ListRestaurantBiz{store: store}
}

func (biz *ListRestaurantBiz) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, "User")

	if err != nil {
		return nil, err
	}
	//
	//ids := make([]int, len(result))
	//
	//for i := range ids {
	//	ids[i] = result[i].Id
	//}

	//likeMap, err := biz.listStore.GetRestaurantLikes(ctx, ids)
	//
	//if err != nil {
	//	log.Println(err)
	//
	//	return result, nil
	//}
	//
	//for i, item := range result {
	//	result[i].LikedCount = likeMap[item.Id]
	//}

	return result, nil
}
