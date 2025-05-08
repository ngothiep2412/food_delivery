package ginuser

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	"g05-food-delivery/component/hasher"
	userbiz "g05-food-delivery/module/user/biz"
	usermodel "g05-food-delivery/module/user/model"
	userstorage "g05-food-delivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := ctx.GetMaiDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
