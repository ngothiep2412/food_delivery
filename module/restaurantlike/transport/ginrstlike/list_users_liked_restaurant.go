package ginrstlike

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	rstlikebiz "g05-food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "g05-food-delivery/module/restaurantlike/model"
	restaurantlikestorage "g05-food-delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMaiDBConnection())

		biz := rstlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
