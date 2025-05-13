package ginrestaurant

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	restaurantbiz "g05-food-delivery/module/restaurant/biz"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
	restaurantrepo "g05-food-delivery/module/restaurant/repository"
	restaurantstorage "g05-food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBindQuery(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSQLStore(db)
		//likeStore := restaurantlikestorage.NewSQLStore(db)

		repo := restaurantrepo.NewListRestaurantRepo(store)
		biz := restaurantbiz.NewListRestaurant(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
