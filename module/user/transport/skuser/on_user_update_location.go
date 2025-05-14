package skuser

import (
	"g05-food-delivery/common"
	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"
	"log"
)

type SmallAppContext interface {
	GetMaiDBConnection() *gorm.DB
}

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx SmallAppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {

		// location belong to user ???
		log.Println("User update location user id is", requester.GetUserId(), "at location", location)
	}
}
