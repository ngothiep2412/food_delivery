package common

import "log"

const (
	DBTypeRestaurant = 1
	DBTypeUser       = 2
)

const (
	CurrentUser = "user"
)

const (
	TopicUserLikeRestaurant   = "TopicUserLikeRestaurant"
	TopicUserUnlikeRestaurant = "TopicUserUnlikeRestaurant"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
