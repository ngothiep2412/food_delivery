package ginuser

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(ctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
