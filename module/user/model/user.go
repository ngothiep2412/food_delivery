package usermodel

import (
	"errors"
	"g05-food-delivery/common"
)

const EntityName = "User"

type User struct {
	common.SqlModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"-" gorm:"column:password"`
	Salt            string        `json:"-" gorm:"column:salt"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"role" gorm:"column:role"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar" gorm:"type:json"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (*User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GetUID(common.DBTypeUser)
}

type UserCreate struct {
	common.SqlModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"-" gorm:"column:password"`
	Salt            string        `json:"-" gorm:"column:salt"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"role" gorm:"column:role"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar" gorm:"type:json"`
}

func (*UserCreate) TableName() string {
	user := User{}
	return user.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GetUID(common.DBTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (*UserLogin) TableName() string {
	user := User{}
	return user.TableName()
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)
	
	ErrEmailExists = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
