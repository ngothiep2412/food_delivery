package restaurantmodel

import (
	"errors"
	"g05-food-delivery/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SqlModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name"`
	Addr            string             `json:"addr" gorm:"column:addr"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover"`
	UserId          int                `json:"-" gorm:"column:owner_id"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GetUID(common.DBTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type RestaurantCreate struct {
	common.SqlModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"addr" gorm:"column:addr"`
	OwnerId         int            `json:"-" gorm:"column:owner_id"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"` // lưu json -> truy vấn lệnh where đc
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GetUID(common.DBTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name"`
	Addr  *string        `json:"address" gorm:"column:address"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo"`
	Cover *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
