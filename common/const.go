package common

import "log"

const (
	DBTypeRestaurant = 1
	DBTypeUser       = 2
)

const (
	CurrentUser = "user"
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
