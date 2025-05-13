package ginrstlike

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	restaurantstorage "g05-food-delivery/module/restaurant/storage"
	rstlikebiz "g05-food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "g05-food-delivery/module/restaurantlike/model"
	restaurantlikestorage "g05-food-delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		db := appCtx.GetMaiDBConnection()

		store := restaurantlikestorage.NewSQLStore(db)
		incStore := restaurantstorage.NewSQLStore(db)

		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
