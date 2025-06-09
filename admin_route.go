package main

import (
	"g05-food-delivery/component/appctx"
	"github.com/gin-gonic/gin"
)

func setupAdminRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	//admin := v1.Group("/admin",
	//	middleware.RequireAuth(appCtx),
	//	middleware.RequireAuth(appCtx), middleware.RoleRequired(appCtx, "admin", "mod"),
	//)
	//{
	//	admin.GET("/profile", ginuser.Profile(appCtx))
	//}
}
