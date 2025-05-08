package main

import (
	"g05-food-delivery/component/appctx"
	"g05-food-delivery/middleware"
	"g05-food-delivery/module/restaurant/transport/ginrestaurant"
	"g05-food-delivery/module/upload/transport/ginupload"
	"g05-food-delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func setupRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	db := appCtx.GetMaiDBConnection()
	restaurants := v1.Group("/restaurants")

	v1.POST("/upload", ginupload.UploadImage(appCtx))

	restaurants.POST("", middleware.RequireAuth(appCtx), ginrestaurant.CreateRestaurant(appCtx))

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var data Restaurant

		db.Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

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

		db.Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.DELETE("/:id", middleware.RequireAuth(appCtx), ginrestaurant.DeleteRestaurant(appCtx))

	// User
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/authenticate", ginuser.Login(appCtx))
}
