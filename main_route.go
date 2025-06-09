package main

import (
	"g05-food-delivery/component/appctx"
	"g05-food-delivery/memcache"
	"g05-food-delivery/middleware"
	"g05-food-delivery/module/restaurant/transport/ginrestaurant"
	"g05-food-delivery/module/restaurantlike/transport/ginrstlike"
	"g05-food-delivery/module/upload/transport/ginupload"
	userstorage "g05-food-delivery/module/user/storage"
	"g05-food-delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func setupRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	userStore := userstorage.NewSQLStore(appCtx.GetMaiDBConnection())
	userCachingStore := memcache.NewUserCaching(memcache.NewCaching(), userStore)
	restaurants := v1.Group("/restaurants")

	v1.POST("/upload", ginupload.UploadImage(appCtx))

	restaurants.POST("", middleware.RequireAuth(appCtx, userCachingStore), ginrestaurant.CreateRestaurant(appCtx))

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var data Restaurant

		appCtx.GetMaiDBConnection().Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.GET("", middleware.RequireAuth(appCtx, userCachingStore), ginrestaurant.ListRestaurant(appCtx))

	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		appCtx.GetMaiDBConnection().Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.DELETE("/:id", middleware.RequireAuth(appCtx, userCachingStore), ginrestaurant.DeleteRestaurant(appCtx))

	restaurants.POST("/:id/liked-users", middleware.RequireAuth(appCtx, userCachingStore), ginrstlike.UserLikeRestaurant(appCtx))
	restaurants.DELETE("/:id/liked-users", middleware.RequireAuth(appCtx, userCachingStore), ginrstlike.UserUnLikeRestaurant(appCtx))
	restaurants.GET("/:id/liked-users", ginrstlike.ListUsers(appCtx))

	// User
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/authenticate", ginuser.Login(appCtx))
}
