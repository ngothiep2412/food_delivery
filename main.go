package main

import (
	"g05-food-delivery/component/appctx"
	"g05-food-delivery/component/uploadprovider"
	"g05-food-delivery/middleware"
	"g05-food-delivery/pubsub/pblocal"
	"g05-food-delivery/skio"
	"g05-food-delivery/subscriber"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
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
	secretKey := os.Getenv("SYSTEM_KEY")

	//
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, S3Domain, s3SecretKey)

	r := gin.Default()

	ps := pblocal.NewPubSub()

	appCtx := appctx.NewAppContext(db, s3Provider, secretKey, ps)

	// setup subscribers
	//subscriber.Setup(appCtx, context.Background())
	_ = subscriber.NewEngine(appCtx).Start()

	r.Use(middleware.Recover(appCtx))

	r.Static("/static", "./static")
	r.StaticFile("/demo", "./demo.html")

	// POST /restaurants
	v1 := r.Group("/api/v1")

	setupRoute(appCtx, v1)
	setupAdminRoute(appCtx, v1)
	//
	rlEngine := skio.NewEngine()
	appCtx.SetRealtimeEngine(rlEngine)

	_ = rlEngine.Run(appCtx, r)

	//startSocketIOServer(r, appCtx)

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

	r.Run()
}

//func startSocketIOServer(engine *gin.Engine, appCtx appctx.AppContext) {
//	server, _ := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{websocket.Default},
//	})
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		//s.SetContext("")
//		fmt.Println("socket connected:", s.ID(), "IP:", s.RemoteAddr())
//
//		s.Join("Shipper")
//
//		s.Emit("test", "test: hello")
//		return nil
//	})
//
//	//go func() {
//	//	for range time.NewTicker(time.Second).C {
//	//		server.BroadcastToRoom("/", "Shipper", "test", "Ahihi")
//	//	}
//	//}()
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		fmt.Println("meet error:", e)
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed:", reason)
//	})
//
//	// Handle user like restaurant events
//	server.OnEvent("/", "test", func(s socketio.Conn, msg interface{}) {
//		log.Println(msg)
//	})
//
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//	}
//
//	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) { // reflex để autodetect nếu struct unmusal binary
//		log.Println("server receive notice:", p.Name, p.Age)
//
//		p.Age = 33
//		s.Emit("notice", p)
//	})
//
//	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
//		db := appCtx.GetMaiDBConnection()
//		store := userstorage.NewSQLStore(db)
//
//		tokenpProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())
//
//		payload, err := tokenpProvider.Validate(token)
//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//
//		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//
//		if user.Status == 0 {
//			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
//			s.Close()
//			return
//		}
//
//		user.Mask(false)
//
//		s.Emit("your_profile", user)
//	})
//
//	go server.Serve()
//
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}
