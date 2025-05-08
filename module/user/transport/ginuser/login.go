package ginuser

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	"g05-food-delivery/component/hasher"
	"g05-food-delivery/component/tokenprovider/jwt"
	userbiz "g05-food-delivery/module/user/biz"
	usermodel "g05-food-delivery/module/user/model"
	userstorage "g05-food-delivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := ctx.GetMaiDBConnection()
		tokenProvider := jwt.NewTokenJwtProvider(ctx.SecretKey())

		store := userstorage.NewSQLStore(db)

		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
