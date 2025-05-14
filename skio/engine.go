package skio

import (
	"fmt"
	"g05-food-delivery/component/tokenprovider/jwt"
	userstorage "g05-food-delivery/module/user/storage"
	"g05-food-delivery/module/user/transport/skuser"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"sync"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	SecretKey() string
	GetRealtimeEngine() RealtimeEngine
}

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket // user có thể nhiều online socket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(ctx AppContext, engine *gin.Engine) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  &sync.RWMutex{},
	}
}

func (engine *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()

	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}

	engine.locker.Unlock()
}

func (engine *rtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.Lock()

	defer engine.locker.Unlock()

	return engine.storage[userId]
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	var sockets []AppSocket

	if scks, ok := engine.storage[userId]; ok {
		scks = append(sockets, scks...)
		sockets = scks
		return sockets
	}

	return sockets
}

func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)

	return nil
}

func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)

	for _, socket := range sockets {
		socket.Emit(key, data)
	}

	return nil
}

func (engine *rtEngine) Run(appCtx AppContext, r *gin.Engine) error {
	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	if err != nil {
		return err
	}

	engine.server = server

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), "IP:", s.RemoteAddr())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("disconnected:", reason)
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		db := appCtx.GetMaiDBConnection()
		store := userstorage.NewSQLStore(db)

		tokenProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", "you has been banned/deleted")
			s.Close()
			return
		}

		user.Mask(false)

		// Important: New AppSocket
		appSck := NewAppSocket(s, user)
		engine.saveAppSocket(user.Id, appSck)

		//appSck.Join(user.GetRole()) // the same
		//if user.GetRole() == "admin" {
		//	appSck.Join("admin")
		//}

		server.OnEvent("/", "UserUpdateLocation", skuser.OnUserUpdateLocation(appCtx, user))

		s.Emit("authenticated", user)
	})

	go server.Serve()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	return nil
}
