package middleware

import (
	"errors"
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	"github.com/gin-gonic/gin"
)

func RoleRequired(ctx appctx.AppContext, allowsRole ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowsRole {
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}
		if !hasFound {
			panic(common.ErrNoPermission(errors.New("invalid role user")))
		}

		c.Next()
	}
}
