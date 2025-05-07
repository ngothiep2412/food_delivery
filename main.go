package main

import (
	"g05-food-delivery/component/appctx"
	"g05-food-delivery/component/uploadprovider"
	"g05-food-delivery/middleware"
	"g05-food-delivery/module/restaurant/transport/ginrestaurant"
	"g05-food-delivery/module/upload/transport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"` // tag
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"address" gorm:"column:address"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	//test := Restaurant{
	//	Id:   1,
	//	Name: "test",
	//	Addr: "127.0.0.1",
	//}

	//json.Marshal(test)
	//
	//log.Println(test) //{1 test 127.0.0.1}

	dsn := os.Getenv("MYSQL_CONN_STRING")

	s3Region := os.Getenv("S3Region")
	s3BucketName := os.Getenv("S3BucketName")
	s3ApiKey := os.Getenv("S3ApiKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	S3Domain := os.Getenv("S3Domain")

	//
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, S3Domain, s3SecretKey)

	r := gin.Default()

	appCtx := appctx.NewAppContext(db, s3Provider)

	r.Use(middleware.Recover(appCtx))

	r.Static("/static", "./static")

	// POST /restaurants
	v1 := r.Group("/api/v1")

	restaurants := v1.Group("/restaurants")

	v1.POST("/upload", ginupload.UploadImage(appCtx))

	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

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

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	r.Run()

	//newRestaurant := Restaurant{Name: "Tani", Addr: "9 Pham Van Hai"}
	//
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println("New id", newRestaurant.Id)
	//
	//var myRestaurant Restaurant
	//
	//if err := db.Where("id = ?", 3).First(&myRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//newName := ""
	//updateData := RestaurantUpdate{Name: &newName}
	//
	//log.Println(myRestaurant)
	//
	//myRestaurant.Name = "200LAB"
	//
	//if err := db.Where("id = ?", 3).Updates(&updateData).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(myRestaurant)
	//
	//if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	//	log.Println(err)
	//}
}
