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

func UserUnLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMaiDBConnection())
		//decStore := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())

		biz := rstlikebiz.NewUserUnlikeRestaurantBiz(store, appCtx.GetPubSub())

		data := restaurantlikemodel.UnLike{
			UserId:       requester.GetUserId(),
			RestaurantId: int(uid.GetLocalID()),
		}

		if err := biz.UnlikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
