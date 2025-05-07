package ginrestaurant

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	restaurantbiz "g05-food-delivery/module/restaurant/biz"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
	restaurantstorage "g05-food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		//go func() {
		//	defer common.AppRecover()
		//
		//	var arr []int
		//
		//	log.Println(arr[0])
		//}()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil { //bind toàn bộ request bind vào
			//c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			//
			//return

			panic(err) // chỉ xaì panic cho tầng ngoài cùng
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurant(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			//c.JSON(http.StatusBadRequest, err)
			//
			//return

			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
